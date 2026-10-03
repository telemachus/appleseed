package appleseed_test

import (
	"regexp"
	"slices"
	"strings"
	"testing"
	"unicode"

	"github.com/telemachus/appleseed"
)

const (
	conClass       = "[bcdfghjkmnpqrstvwxzBCDFGHJKMNPQRSTVWXZ]"
	vowClass       = "[aeiouyAEIUY]"
	digClass       = "[0-9]"
	validUppercase = "BCDFGHJKMNPQRSTVWXZAEIUY"
	allChars       = "aeiouyAEIUYbcdfghjkmnpqrstvwxzBCDFGHJKMNPQRSTVWXZ0123456789"

	word       = conClass + vowClass + conClass + conClass + vowClass + conClass
	digitLast  = conClass + vowClass + conClass + conClass + vowClass + digClass
	digitFirst = digClass + conClass + vowClass + conClass + conClass + vowClass
)

var (
	pwRegexes               = shapes("-")
	pwRegexesWithoutHyphens = shapes("")
)

func TestPasswordReturnsValidPasses(t *testing.T) {
	t.Parallel()
	for range 1000 {
		pw := appleseed.Password()
		if !isValid(pw, pwRegexes) {
			t.Fatalf("Password() returned invalid pass: %q", pw)
		}
	}
}

func TestPasswordWithoutHyphensReturnsValidPasses(t *testing.T) {
	t.Parallel()
	for range 1000 {
		pw := appleseed.PasswordWithoutHyphens()
		if !isValid(pw, pwRegexesWithoutHyphens) {
			t.Fatalf("PasswordWithoutHyphens() returned invalid pass: %q", pw)
		}
	}
}

func TestPasswordUsesEveryChoice(t *testing.T) {
	t.Parallel()
	seen, digitAt := tally(1000)
	for _, r := range allChars {
		if !seen[r] {
			t.Errorf("Password() never generated %q", r)
		}
	}
	for _, i := range []int{5, 7, 12, 14, 19} {
		if !digitAt[i] {
			t.Errorf("Password() never put a digit at index %d", i)
		}
	}
}

// shapes returns anchored patterns for the five legal password shapes, with
// pieces joined by sep.
func shapes(sep string) []*regexp.Regexp {
	layouts := [][]string{
		{digitLast, word, word},
		{word, digitFirst, word},
		{word, digitLast, word},
		{word, word, digitFirst},
		{word, word, digitLast},
	}
	regexes := make([]*regexp.Regexp, 0, len(layouts))
	for _, layout := range layouts {
		regexes = append(
			regexes,
			regexp.MustCompile("^"+strings.Join(layout, sep)+"$"),
		)
	}

	return regexes
}

func isValid(pw string, regexes []*regexp.Regexp) bool {
	return matchesAny(pw, regexes) && countCharsIn(pw, validUppercase) == 1
}

func matchesAny(pw string, regexes []*regexp.Regexp) bool {
	return slices.ContainsFunc(regexes, func(re *regexp.Regexp) bool {
		return re.MatchString(pw)
	})
}

func countCharsIn(s, chars string) int {
	count := 0
	for i := range len(s) {
		if strings.IndexByte(chars, s[i]) >= 0 {
			count++
		}
	}

	return count
}

// tally generates n passwords and records which characters appear and at
// which indexes a digit appears.
func tally(n int) (map[rune]bool, map[int]bool) {
	seen := map[rune]bool{}
	digitAt := map[int]bool{}
	for range n {
		for i, r := range appleseed.Password() {
			if unicode.IsDigit(r) {
				digitAt[i] = true
			}
			seen[r] = true
		}
	}

	return seen, digitAt
}
