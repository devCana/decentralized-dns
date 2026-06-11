package chain

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

// WatchRecordEvents polls the contract for Registered / Transferred /
// RecordSet / RecordRemoved events and delivers them to handler in block
// order. Polling (rather than a WS subscription) works against any RPC
// transport, including plain-HTTP Hardhat (HLD open issue 5: TTL + push
// invalidation). Blocks until ctx is cancelled.
func (c *Client) WatchRecordEvents(ctx context.Context, interval time.Duration, handler func(RecordEvent)) error {
	last, err := c.ChainHead(ctx)
	if err != nil {
		return err
	}
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
		}

		head, err := c.eth.BlockNumber(ctx)
		if err != nil || head <= last {
			if err != nil {
				c.log.Warn("event poll: head", "err", err)
			}
			continue
		}
		from := last + 1
		opts := &bind.FilterOpts{Start: from, End: &head, Context: ctx}
		c.collectEvents(opts, handler)
		last = head
	}
}

func (c *Client) collectEvents(opts *bind.FilterOpts, handler func(RecordEvent)) {
	if it, err := c.dapp.FilterRecordSet(opts, nil, nil); err == nil {
		for it.Next() {
			handler(RecordEvent{
				Kind: EventRecordSet, Name: it.Event.Name, NameHash: it.Event.NameHash,
				RecordType: it.Event.RecordType, Selector: it.Event.Selector,
				Block: it.Event.Raw.BlockNumber,
			})
		}
		it.Close()
	} else {
		c.log.Warn("event poll: RecordSet", "err", err)
	}

	if it, err := c.dapp.FilterRecordRemoved(opts, nil, nil); err == nil {
		for it.Next() {
			handler(RecordEvent{
				Kind: EventRecordRemoved, Name: it.Event.Name, NameHash: it.Event.NameHash,
				RecordType: it.Event.RecordType, Selector: it.Event.Selector,
				Block: it.Event.Raw.BlockNumber,
			})
		}
		it.Close()
	} else {
		c.log.Warn("event poll: RecordRemoved", "err", err)
	}

	if it, err := c.dapp.FilterTransferred(opts, nil, nil, nil); err == nil {
		for it.Next() {
			handler(RecordEvent{
				Kind: EventTransferred, NameHash: it.Event.NameHash,
				Block: it.Event.Raw.BlockNumber,
			})
		}
		it.Close()
	} else {
		c.log.Warn("event poll: Transferred", "err", err)
	}

	if it, err := c.dapp.FilterRegistered(opts, nil, nil); err == nil {
		for it.Next() {
			handler(RecordEvent{
				Kind: EventRegistered, Name: it.Event.Name, NameHash: it.Event.NameHash,
				Block: it.Event.Raw.BlockNumber,
			})
		}
		it.Close()
	} else {
		c.log.Warn("event poll: Registered", "err", err)
	}
}
