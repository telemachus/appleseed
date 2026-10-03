// A simplistic command line example that uses Appleseed to generate Apple-style
// more typeable passwords.
package main

import (
	"flag"
	"fmt"

	"github.com/telemachus/appleseed"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintln(
			flag.CommandLine.Output(),
			"Usage: appleseed [-no-hyphens]\n\nOptions:",
		)
		flag.PrintDefaults()
	}

	noHyphens := flag.Bool("no-hyphens", false, "generate password without hyphens")
	flag.Parse()

	if *noHyphens {
		fmt.Println(appleseed.PasswordWithoutHyphens())
	} else {
		fmt.Println(appleseed.Password())
	}
}
