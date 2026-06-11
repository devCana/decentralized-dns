// Command ddns-lookup queries a resolver and verifies the response
// signatures and ZK proof (HLD §4.3). Implemented in issue #16.
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(os.Stderr, "ddns-lookup: not yet implemented (see issue #16)")
	os.Exit(2)
}
