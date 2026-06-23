// Command ddns is the domain-owner CLI (HLD §4.3, UC-1/2/3/7/9). It signs
// every transaction locally with the owner's key — the key never leaves the
// machine and is never sent to a resolver — and submits namespace operations
// directly to the NamespaceDApp / RecordSchemaRegistry contracts.
//
//	ddns register <name>
//	ddns set <name> <type> [--selector S] [--ttl N] k=v ...
//	ddns remove <name> <type> [--selector S]
//	ddns transfer <name> <newOwner> --pubkey 0x...
//	ddns renew <name>
//	ddns declare-type <name> --mandatory a,b [--optional c,d]
//	ddns publish-resource <name> <file> [--selector S] [--ttl N]
package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"mime"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/devCana/decentralized-dns/resolver/internal/chain"
	"github.com/devCana/decentralized-dns/resolver/internal/chain/bindings"
	"github.com/devCana/decentralized-dns/resolver/internal/pki"
	"github.com/devCana/decentralized-dns/resolver/internal/query"
	bttorrent "github.com/devCana/decentralized-dns/resolver/internal/torrent"
	"github.com/devCana/decentralized-dns/resolver/internal/zk"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}
	args := os.Args[2:]
	switch os.Args[1] {
	case "register":
		cmdRegister(args)
	case "set":
		cmdSet(args)
	case "remove":
		cmdRemove(args)
	case "transfer":
		cmdTransfer(args)
	case "renew":
		cmdRenew(args)
	case "declare-type":
		cmdDeclareType(args)
	case "publish-resource":
		cmdPublishResource(args)
	case "-h", "--help", "help":
		usage()
	default:
		fmt.Fprintf(os.Stderr, "ddns: unknown command %q\n\n", os.Args[1])
		usage()
		os.Exit(2)
	}
}

func usage() {
	fmt.Fprint(os.Stderr, `ddns — decentralized DNS owner CLI

usage:
  ddns register <name>
  ddns set <name> <type> [--selector S] [--ttl N] key=value ...
  ddns remove <name> <type> [--selector S]
  ddns transfer <name> <newOwnerAddr> --pubkey 0x<uncompressed-secp256k1>
  ddns renew <name>
  ddns declare-type <name> --mandatory a,b [--optional c,d]
  ddns publish-resource <name> <file> [--selector S] [--ttl N] [--content-type CT]

common flags (all subcommands):
  --rpc URL           blockchain RPC (env RPC_URL, default http://127.0.0.1:8545)
  --key HEX           owner private key (env DDNS_PRIVATE_KEY)
  --deployments PATH  deploy json with contract addresses
                      (env DDNS_DEPLOYMENTS, default contracts/deployments/localhost.json)
  --namespace ADDR    NamespaceDApp address (overrides deployments)
  --registry ADDR     RecordSchemaRegistry address (overrides deployments)
`)
}

// ---------------------------------------------------------------- subcommands

func cmdRegister(args []string) {
	fs := flag.NewFlagSet("register", flag.ExitOnError)
	co := addCommon(fs)
	_ = fs.Parse(reorder(args))
	name := needArg(fs, 0, "name")
	c := dial(co, false)

	pubKey := crypto.FromECDSAPub(&c.key.PublicKey) // 65-byte uncompressed
	price, err := c.dapp.PriceOf(callOpts(), name)
	fatal(err)
	fmt.Printf("registering %q for %s (fee %s wei)\n", name, c.from.Hex(), price)

	c.auth.Value = price
	send(c, "register", func(o *bind.TransactOpts) (*types.Transaction, error) {
		return c.dapp.Register(o, name, pubKey)
	})
	fmt.Printf("registered %q -> owner %s, pubKey %s\n", name, c.from.Hex(), hexPub(pubKey))
}

func cmdSet(args []string) {
	fs := flag.NewFlagSet("set", flag.ExitOnError)
	co := addCommon(fs)
	selector := fs.String("selector", "", `selector, e.g. "service=HTTP&transport=TCP&port=443"`)
	ttl := fs.Uint("ttl", 3600, "record TTL in seconds")
	_ = fs.Parse(reorder(args))
	name := needArg(fs, 0, "name")
	recordType := needArg(fs, 1, "type")
	fieldNames, fieldValues := parsePairs(fs.Args()[2:])

	sel := canonical(*selector)
	c := dial(co, false)

	rec := chain.Record{
		Type: recordType, Selector: sel,
		FieldNames: fieldNames, FieldVals: fieldValues, TTL: uint32(*ttl),
	}
	sig, err := pki.SignRecord(name, rec, c.key)
	fatal(err)
	commitment, err := zk.Commitment(pki.RecordMessage(name, rec))
	fatal(err)

	send(c, "setRecord", func(o *bind.TransactOpts) (*types.Transaction, error) {
		return c.dapp.SetRecord(o, name, recordType, sel, fieldNames, fieldValues, uint32(*ttl), sig, commitment)
	})
	fmt.Printf("set %s %s%s on %q\n", recordType, fieldsString(fieldNames, fieldValues), bracket(sel), name)
}

func cmdRemove(args []string) {
	fs := flag.NewFlagSet("remove", flag.ExitOnError)
	co := addCommon(fs)
	selector := fs.String("selector", "", "selector to remove")
	_ = fs.Parse(reorder(args))
	name := needArg(fs, 0, "name")
	recordType := needArg(fs, 1, "type")
	sel := canonical(*selector)
	c := dial(co, false)

	send(c, "removeRecord", func(o *bind.TransactOpts) (*types.Transaction, error) {
		return c.dapp.RemoveRecord(o, name, recordType, sel)
	})
	fmt.Printf("removed %s%s from %q\n", recordType, bracket(sel), name)
}

func cmdTransfer(args []string) {
	fs := flag.NewFlagSet("transfer", flag.ExitOnError)
	co := addCommon(fs)
	pubkey := fs.String("pubkey", "", "new owner's uncompressed secp256k1 public key (0x04...)")
	_ = fs.Parse(reorder(args))
	name := needArg(fs, 0, "name")
	newOwner := needArg(fs, 1, "newOwner")
	if !common.IsHexAddress(newOwner) {
		fatal(fmt.Errorf("invalid new owner address %q", newOwner))
	}
	if *pubkey == "" {
		fatal(errors.New("--pubkey is required (the new owner's public key)"))
	}
	newPub := common.FromHex(*pubkey)
	if len(newPub) == 0 || len(newPub) > 128 {
		fatal(fmt.Errorf("invalid --pubkey length %d", len(newPub)))
	}
	c := dial(co, false)
	send(c, "transfer", func(o *bind.TransactOpts) (*types.Transaction, error) {
		return c.dapp.Transfer(o, name, common.HexToAddress(newOwner), newPub)
	})
	fmt.Printf("transferred %q to %s (previous records no longer resolve)\n", name, newOwner)
}

func cmdRenew(args []string) {
	fs := flag.NewFlagSet("renew", flag.ExitOnError)
	co := addCommon(fs)
	_ = fs.Parse(reorder(args))
	name := needArg(fs, 0, "name")
	c := dial(co, false)
	price, err := c.dapp.PriceOf(callOpts(), name)
	fatal(err)
	c.auth.Value = price
	send(c, "renew", func(o *bind.TransactOpts) (*types.Transaction, error) {
		return c.dapp.Renew(o, name)
	})
	fmt.Printf("renewed %q (fee %s wei)\n", name, price)
}

func cmdDeclareType(args []string) {
	fs := flag.NewFlagSet("declare-type", flag.ExitOnError)
	co := addCommon(fs)
	mandatory := fs.String("mandatory", "", "comma-separated mandatory field names")
	optional := fs.String("optional", "", "comma-separated optional field names")
	_ = fs.Parse(reorder(args))
	name := needArg(fs, 0, "name")
	c := dial(co, true)
	mand := splitCSV(*mandatory)
	opt := splitCSV(*optional)
	send(c, "declareType", func(o *bind.TransactOpts) (*types.Transaction, error) {
		return c.reg.DeclareType(o, name, mand, opt)
	})
	fmt.Printf("declared type %q (mandatory=%v optional=%v)\n", name, mand, opt)
}

func cmdPublishResource(args []string) {
	fs := flag.NewFlagSet("publish-resource", flag.ExitOnError)
	co := addCommon(fs)
	selector := fs.String("selector", "", `selector, e.g. "service=HTTP"`)
	ttl := fs.Uint("ttl", 3600, "record TTL in seconds")
	contentType := fs.String("content-type", "", "MIME type (default: detected from the file)")
	dataDir := fs.String("data-dir", "", "torrent data dir (default: a temp dir)")
	btPort := fs.Int("bt-port", 0, "BitTorrent listen port (0 = random)")
	seconds := fs.Int("seconds", 0, "seed for N seconds then exit (0 = until interrupted)")
	anchorOnly := fs.Bool("anchor-only", false, "submit the record without seeding")
	_ = fs.Parse(reorder(args, "anchor-only"))
	name := needArg(fs, 0, "name")
	file := needArg(fs, 1, "file")
	sel := canonical(*selector)
	ct := *contentType
	if ct == "" {
		ct = detectContentType(file)
	}

	dir := *dataDir
	if dir == "" {
		var err error
		dir, err = os.MkdirTemp("", "ddns-seed-*")
		fatal(err)
	}
	engine, err := bttorrent.New(bttorrent.Config{DataDir: dir, ListenPort: *btPort, DisableDHT: false})
	fatal(err)
	defer engine.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	infoHash, sha, err := engine.SeedFile(ctx, file)
	cancel()
	fatal(err)
	fmt.Printf("seeded %s: infoHash=%s sha256=%s contentType=%s\n", file, infoHash, sha, ct)

	c := dial(co, false)
	rec := chain.Record{
		Type: "ResourceRef", Selector: sel,
		FieldNames: []string{"infoHash", "sha256", "contentType"},
		FieldVals:  []string{infoHash, sha, ct},
		TTL:        uint32(*ttl),
	}
	sig, err := pki.SignRecord(name, rec, c.key)
	fatal(err)
	commitment, err := zk.Commitment(pki.RecordMessage(name, rec))
	fatal(err)
	send(c, "setRecord(ResourceRef)", func(o *bind.TransactOpts) (*types.Transaction, error) {
		return c.dapp.SetRecord(o, name, "ResourceRef", sel, rec.FieldNames, rec.FieldVals, uint32(*ttl), sig, commitment)
	})
	fmt.Printf("anchored ResourceRef for %q%s\n", name, bracket(sel))

	if *anchorOnly {
		return
	}
	fmt.Printf("seeding on %v\n", engine.ListenAddrs())
	if *seconds > 0 {
		fmt.Printf("seeding for %ds...\n", *seconds)
		time.Sleep(time.Duration(*seconds) * time.Second)
		return
	}
	fmt.Println("seeding until interrupted (Ctrl-C)...")
	sig2 := make(chan os.Signal, 1)
	signal.Notify(sig2, os.Interrupt, syscall.SIGTERM)
	<-sig2
}

// ------------------------------------------------------------------- plumbing

type conn struct {
	eth  *ethclient.Client
	dapp *bindings.NamespaceDApp
	reg  *bindings.RecordSchemaRegistry
	key  *ecdsa.PrivateKey
	auth *bind.TransactOpts
	from common.Address
}

type commonOpts struct {
	rpc, key, deployments, namespace, registry *string
}

func addCommon(fs *flag.FlagSet) commonOpts {
	return commonOpts{
		rpc:         fs.String("rpc", envOr("RPC_URL", "http://127.0.0.1:8545"), "blockchain RPC URL"),
		key:         fs.String("key", "", "owner private key hex (env DDNS_PRIVATE_KEY)"),
		deployments: fs.String("deployments", envOr("DDNS_DEPLOYMENTS", "contracts/deployments/localhost.json"), "deploy json path"),
		namespace:   fs.String("namespace", "", "NamespaceDApp address"),
		registry:    fs.String("registry", "", "RecordSchemaRegistry address"),
	}
}

// dial connects to the chain and builds a keyed transactor. needRegistry
// requires the RecordSchemaRegistry address to be known.
func dial(co commonOpts, needRegistry bool) *conn {
	key := loadKey(*co.key)
	ns, reg := resolveAddrs(*co.deployments, *co.namespace, *co.registry)
	if ns == (common.Address{}) {
		fatal(errors.New("NamespaceDApp address unknown: set --deployments, --namespace, or CONTRACT_ADDRESS"))
	}
	if needRegistry && reg == (common.Address{}) {
		fatal(errors.New("RecordSchemaRegistry address unknown: set --deployments, --registry, or REGISTRY_ADDRESS"))
	}
	ctx := context.Background()
	eth, err := ethclient.DialContext(ctx, *co.rpc)
	fatal(err)
	chainID, err := eth.ChainID(ctx)
	fatal(err)
	auth, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	fatal(err)
	dapp, err := bindings.NewNamespaceDApp(ns, eth)
	fatal(err)
	var regc *bindings.RecordSchemaRegistry
	if reg != (common.Address{}) {
		regc, err = bindings.NewRecordSchemaRegistry(reg, eth)
		fatal(err)
	}
	return &conn{eth: eth, dapp: dapp, reg: regc, key: key, auth: auth, from: crypto.PubkeyToAddress(key.PublicKey)}
}

// send submits a transaction and waits for a successful receipt.
func send(c *conn, label string, fn func(*bind.TransactOpts) (*types.Transaction, error)) {
	tx, err := fn(c.auth)
	fatal(err)
	fmt.Printf("%s: tx %s submitted, waiting...\n", label, tx.Hash().Hex())
	rcpt, err := bind.WaitMined(context.Background(), c.eth, tx)
	fatal(err)
	if rcpt.Status != types.ReceiptStatusSuccessful {
		fatal(fmt.Errorf("%s reverted (tx %s)", label, tx.Hash().Hex()))
	}
	fmt.Printf("%s: confirmed in block %d (gas %d)\n", label, rcpt.BlockNumber.Uint64(), rcpt.GasUsed)
}

func resolveAddrs(deployments, nsFlag, regFlag string) (ns, reg common.Address) {
	if nsFlag != "" {
		ns = common.HexToAddress(nsFlag)
	}
	if regFlag != "" {
		reg = common.HexToAddress(regFlag)
	}
	if (ns == common.Address{} || reg == common.Address{}) && deployments != "" {
		if data, err := os.ReadFile(deployments); err == nil {
			var d struct {
				Contracts struct {
					NamespaceDApp        string
					RecordSchemaRegistry string
				} `json:"contracts"`
			}
			if json.Unmarshal(data, &d) == nil {
				if ns == (common.Address{}) && d.Contracts.NamespaceDApp != "" {
					ns = common.HexToAddress(d.Contracts.NamespaceDApp)
				}
				if reg == (common.Address{}) && d.Contracts.RecordSchemaRegistry != "" {
					reg = common.HexToAddress(d.Contracts.RecordSchemaRegistry)
				}
			}
		}
	}
	if ns == (common.Address{}) {
		if v := os.Getenv("CONTRACT_ADDRESS"); v != "" {
			ns = common.HexToAddress(v)
		}
	}
	if reg == (common.Address{}) {
		if v := os.Getenv("REGISTRY_ADDRESS"); v != "" {
			reg = common.HexToAddress(v)
		}
	}
	return
}

func loadKey(flagVal string) *ecdsa.PrivateKey {
	raw := flagVal
	if raw == "" {
		raw = os.Getenv("DDNS_PRIVATE_KEY")
	}
	if raw == "" {
		fatal(errors.New("no owner key: set --key or DDNS_PRIVATE_KEY"))
	}
	key, err := crypto.HexToECDSA(strings.TrimPrefix(strings.TrimSpace(raw), "0x"))
	fatal(err)
	return key
}

func callOpts() *bind.CallOpts { return &bind.CallOpts{Context: context.Background()} }

// ---------------------------------------------------------------- small utils

func parsePairs(args []string) (names, values []string) {
	for _, a := range args {
		k, v, ok := strings.Cut(a, "=")
		if !ok {
			fatal(fmt.Errorf("field %q must be key=value", a))
		}
		names = append(names, k)
		values = append(values, v)
	}
	return
}

func canonical(selector string) string {
	sel, err := query.ParseSelectorString(selector)
	fatal(err)
	return sel
}

func detectContentType(path string) string {
	if ct := mime.TypeByExtension(filepath.Ext(path)); ct != "" {
		return ct
	}
	f, err := os.Open(path)
	if err == nil {
		defer f.Close()
		buf := make([]byte, 512)
		n, _ := f.Read(buf)
		return http.DetectContentType(buf[:n])
	}
	return "application/octet-stream"
}

func splitCSV(s string) []string {
	if strings.TrimSpace(s) == "" {
		return nil
	}
	parts := strings.Split(s, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		if p = strings.TrimSpace(p); p != "" {
			out = append(out, p)
		}
	}
	return out
}

func fieldsString(names, values []string) string {
	var b strings.Builder
	for i, n := range names {
		if i > 0 {
			b.WriteByte(' ')
		}
		fmt.Fprintf(&b, "%s=%s", n, values[i])
	}
	return b.String()
}

func bracket(sel string) string {
	if sel == "" {
		return ""
	}
	return " [" + sel + "]"
}

func hexPub(b []byte) string { return "0x" + common.Bytes2Hex(b) }

// reorder moves flags ahead of positional args so flags may appear anywhere on
// the command line (Go's flag package otherwise stops at the first positional).
// Named bool flags take no value; everything else consumes the next token.
func reorder(args []string, boolFlags ...string) []string {
	bset := map[string]bool{}
	for _, b := range boolFlags {
		bset[b] = true
	}
	var flags, pos []string
	for i := 0; i < len(args); i++ {
		a := args[i]
		if a == "--" {
			pos = append(pos, args[i+1:]...)
			break
		}
		if len(a) > 1 && a[0] == '-' {
			flags = append(flags, a)
			name := strings.SplitN(strings.TrimLeft(a, "-"), "=", 2)[0]
			if !strings.Contains(a, "=") && !bset[name] && i+1 < len(args) {
				i++
				flags = append(flags, args[i])
			}
		} else {
			pos = append(pos, a)
		}
	}
	return append(flags, pos...)
}

func needArg(fs *flag.FlagSet, i int, label string) string {
	if i >= fs.NArg() {
		fmt.Fprintf(os.Stderr, "ddns: missing <%s>\n", label)
		os.Exit(2)
	}
	return fs.Arg(i)
}

func envOr(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func fatal(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "ddns:", err)
		os.Exit(1)
	}
}
