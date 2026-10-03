// Package appleseed provides functions that generate Apple-style more typeable
// passwords.
package appleseed

import (
	"strings"
)

const hyphen = "-"

// Password generates and returns an Apple-style more typeable password.
// E.g., funrus-Hommez-kajzo7.
func Password() string {
	return strings.Join(passwordGroups(), hyphen)
}

// PasswordWithoutHyphens generates and returns an Apple-style more typeable
// password without hyphens. E.g., funrusHommezkajzo7.
func PasswordWithoutHyphens() string {
	return strings.Join(passwordGroups(), "")
}
