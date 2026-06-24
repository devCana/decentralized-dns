// Package config loads resolver settings from environment variables,
// per HLD §4.4 (RPC_URL, CONTRACT_ADDRESS, RESOLVER_KEYSTORE, REST_PORT,
// UDP_PORT, BT_LISTEN_PORT).
package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Config holds every tunable of the resolver process.
type Config struct {
	RPCURL          string // Web3 RPC endpoint of the blockchain node
	ContractAddress string // NamespaceDApp address (hex)
	RegistryAddress string // RecordSchemaRegistry address (hex)
	KeystorePath    string // path to the resolver identity key file
	RESTPort        int    // REST QueryAPI listen port
	UDPPort         int    // UDP QueryAPI listen port
	BTListenPort    int    // BitTorrent engine listen port
	CacheSize       int    // max entries in the TTL LRU cache
	RateRPS         int    // per-IP REST rate limit, requests/second
	RateBurst       int    // per-IP REST burst allowance
	DataDir         string // scratch dir for torrent data
	AllowPeerHints  bool   // honour client-supplied ?peer= hints on /resource
	EnforceType     bool   // reject /resource bytes that mismatch the declared content type
}

// FromEnv builds a Config from the process environment, applying defaults
// for everything except addresses, which stay empty until provided.
func FromEnv() (*Config, error) {
	cfg := &Config{
		RPCURL:          getEnv("RPC_URL", "http://127.0.0.1:8545"),
		ContractAddress: os.Getenv("CONTRACT_ADDRESS"),
		RegistryAddress: os.Getenv("REGISTRY_ADDRESS"),
		KeystorePath:    getEnv("RESOLVER_KEYSTORE", "resolver.key"),
		DataDir:         getEnv("DATA_DIR", "./data"),
		AllowPeerHints:  getEnvBool("ALLOW_PEER_HINTS", false),
		EnforceType:     getEnvBool("ENFORCE_CONTENT_TYPE", false),
	}
	var err error
	if cfg.RESTPort, err = getEnvPort("REST_PORT", 8080); err != nil {
		return nil, err
	}
	if cfg.UDPPort, err = getEnvPort("UDP_PORT", 5353); err != nil {
		return nil, err
	}
	if cfg.BTListenPort, err = getEnvInt("BT_LISTEN_PORT", 42069); err != nil {
		return nil, err
	}
	if cfg.BTListenPort < 0 || cfg.BTListenPort > 65535 {
		return nil, fmt.Errorf("BT_LISTEN_PORT %d out of range (0-65535)", cfg.BTListenPort)
	}
	if cfg.CacheSize, err = getEnvInt("CACHE_SIZE", 4096); err != nil {
		return nil, err
	}
	if cfg.CacheSize <= 0 {
		return nil, fmt.Errorf("CACHE_SIZE must be positive, got %d", cfg.CacheSize)
	}
	if cfg.RateRPS, err = getEnvInt("RATE_RPS", 20); err != nil {
		return nil, err
	}
	if cfg.RateBurst, err = getEnvInt("RATE_BURST", 40); err != nil {
		return nil, err
	}
	if cfg.RateRPS <= 0 || cfg.RateBurst <= 0 {
		return nil, fmt.Errorf("RATE_RPS and RATE_BURST must be positive, got %d and %d", cfg.RateRPS, cfg.RateBurst)
	}
	return cfg, nil
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func getEnvInt(key string, def int) (int, error) {
	v := os.Getenv(key)
	if v == "" {
		return def, nil
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return 0, fmt.Errorf("env %s: expected integer, got %q", key, v)
	}
	return n, nil
}

// getEnvPort parses a TCP/UDP listen port, requiring 1-65535.
func getEnvPort(key string, def int) (int, error) {
	n, err := getEnvInt(key, def)
	if err != nil {
		return 0, err
	}
	if n < 1 || n > 65535 {
		return 0, fmt.Errorf("env %s: port %d out of range (1-65535)", key, n)
	}
	return n, nil
}

func getEnvBool(key string, def bool) bool {
	v := strings.TrimSpace(strings.ToLower(os.Getenv(key)))
	switch v {
	case "":
		return def
	case "1", "true", "yes", "on":
		return true
	default:
		return false
	}
}
