package config

import "testing"

func TestFromEnvDefaults(t *testing.T) {
	cfg, err := FromEnv()
	if err != nil {
		t.Fatal(err)
	}
	if cfg.RPCURL != "http://127.0.0.1:8545" {
		t.Errorf("RPCURL default = %q", cfg.RPCURL)
	}
	if cfg.RESTPort != 8080 || cfg.UDPPort != 5353 || cfg.BTListenPort != 42069 {
		t.Errorf("port defaults = %d/%d/%d", cfg.RESTPort, cfg.UDPPort, cfg.BTListenPort)
	}
	if cfg.CacheSize != 4096 {
		t.Errorf("CacheSize default = %d", cfg.CacheSize)
	}
}

func TestFromEnvOverridesAndErrors(t *testing.T) {
	t.Setenv("REST_PORT", "9090")
	t.Setenv("CONTRACT_ADDRESS", "0xabc")
	cfg, err := FromEnv()
	if err != nil {
		t.Fatal(err)
	}
	if cfg.RESTPort != 9090 {
		t.Errorf("RESTPort = %d, want 9090", cfg.RESTPort)
	}
	if cfg.ContractAddress != "0xabc" {
		t.Errorf("ContractAddress = %q", cfg.ContractAddress)
	}

	t.Setenv("UDP_PORT", "not-a-number")
	if _, err := FromEnv(); err == nil {
		t.Error("expected error for non-integer UDP_PORT")
	}
}
