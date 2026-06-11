// Package config loads resolver settings from environment variables,
// per HLD §4.4 (RPC_URL, CONTRACT_ADDRESS, RESOLVER_KEYSTORE, REST_PORT,
// UDP_PORT, BT_LISTEN_PORT).
package config

import (
	"fmt"
	"os"
	"strconv"
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
	DataDir         string // scratch dir for torrent data
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
	}
	var err error
	if cfg.RESTPort, err = getEnvInt("REST_PORT", 8080); err != nil {
		return nil, err
	}
	if cfg.UDPPort, err = getEnvInt("UDP_PORT", 5353); err != nil {
		return nil, err
	}
	if cfg.BTListenPort, err = getEnvInt("BT_LISTEN_PORT", 42069); err != nil {
		return nil, err
	}
	if cfg.CacheSize, err = getEnvInt("CACHE_SIZE", 4096); err != nil {
		return nil, err
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
