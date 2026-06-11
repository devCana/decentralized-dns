// Command ddns is the domain-owner CLI: register, set, transfer, renew,
// declare-type, publish-resource (HLD §4.3). Implemented in issue #15.
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(os.Stderr, "ddns: not yet implemented (see issue #15)")
	os.Exit(2)
}
