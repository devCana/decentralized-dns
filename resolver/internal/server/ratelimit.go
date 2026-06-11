package server

import (
	"sync"
	"time"

	"golang.org/x/time/rate"
)

// maxTrackedIPs bounds the limiter map; old entries are pruned past this.
const maxTrackedIPs = 4096

// ipLimiter applies a per-IP token bucket (HLD §4.1.2 request validation
// and rate limiting).
type ipLimiter struct {
	mu      sync.Mutex
	entries map[string]*limiterEntry
	rps     rate.Limit
	burst   int
	now     func() time.Time
}

type limiterEntry struct {
	lim      *rate.Limiter
	lastSeen time.Time
}

func newIPLimiter(rps, burst int) *ipLimiter {
	return &ipLimiter{
		entries: make(map[string]*limiterEntry),
		rps:     rate.Limit(rps),
		burst:   burst,
		now:     time.Now,
	}
}

// allow reports whether ip may make a request right now.
func (l *ipLimiter) allow(ip string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	e, ok := l.entries[ip]
	if !ok {
		if len(l.entries) >= maxTrackedIPs {
			l.prune()
		}
		e = &limiterEntry{lim: rate.NewLimiter(l.rps, l.burst)}
		l.entries[ip] = e
	}
	e.lastSeen = l.now()
	return e.lim.Allow()
}

// prune drops entries idle for over a minute (caller holds the lock).
func (l *ipLimiter) prune() {
	cutoff := l.now().Add(-time.Minute)
	for ip, e := range l.entries {
		if e.lastSeen.Before(cutoff) {
			delete(l.entries, ip)
		}
	}
}
