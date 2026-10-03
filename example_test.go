package appleseed_test

import (
	"fmt"

	"github.com/telemachus/appleseed"
)

// Since output is random, we don't want Output comments.

func ExamplePassword() {
	pw := appleseed.Password()
	fmt.Println(pw) // E.g., funrus-Hommez-kajzo7
}

func ExamplePasswordWithoutHyphens() {
	pw := appleseed.PasswordWithoutHyphens()
	fmt.Println(pw) // E.g., funrusHommezkajzo7
}
