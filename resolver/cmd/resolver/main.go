// Command resolver runs the Decentralized DNS Resolver Server: REST/UDP
// query APIs backed by a TTL cache, the blockchain client, the BitTorrent
// engine, and the PKI/ZK verifier (HLD §3).
package main

import (
	"fmt"
	"os"

	"github.com/devCana/decentralized-dns/resolver/internal/config"
)

func main() {
	cfg, err := config.FromEnv()
	if err != nil {
		fmt.Fprintln(os.Stderr, "config:", err)
		os.Exit(1)
	}
	// Subsystems are wired in as they are implemented (issues #4-#14).
	fmt.Printf("ddns-resolver scaffold: REST=:%d UDP=:%d RPC=%s\n",
		cfg.RESTPort, cfg.UDPPort, cfg.RPCURL)
}
