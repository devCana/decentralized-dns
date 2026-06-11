// Command ddns-fetch resolves a ResourceRef and writes the verified file
// to disk (HLD §4.3). Implemented in issue #16.
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(os.Stderr, "ddns-fetch: not yet implemented (see issue #16)")
	os.Exit(2)
}
