package appleseed

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const (
	digits              = "0123456789"
	lowercaseConsonants = "bcdfghjkmnpqrstvwxz"
	lowercaseVowels     = "aeiouy"
	segmentLen          = 6
	baseLen             = 3*segmentLen - 1 // 18 chars minus one for a digit.
)

// passwordGroups returns three six-character strings: two are CVCCVC, and the
// other is either DCVCCV or CVCCVD.
func passwordGroups() []string {
	digit := string(pickFrom(digits))
	base := basePass()
	seg1 := base[:segmentLen]
	seg2 := base[segmentLen : segmentLen*2]
	seg3 := base[segmentLen*2:]

	layouts := [][]string{
		{seg3 + digit, seg1, seg2},
		{seg1, digit + seg3, seg2},
		{seg1, seg3 + digit, seg2},
		{seg1, seg2, digit + seg3},
		{seg1, seg2, seg3 + digit},
	}

	return layouts[randInt(len(layouts))]
}

func pickFrom(s string) byte {
	return s[randInt(len(s))]
}

func basePass() string {
	b := make([]byte, 0, baseLen)
	// Start with CVC five times; then add CV. The digit comes later.
	for range 5 {
		b = appendSyllable(b)
	}
	b = append(b, pickFrom(lowercaseConsonants), pickFrom(lowercaseVowels))
	uppercaseOne(b)

	return string(b)
}

func randInt(n int) int {
	i, err := rand.Int(rand.Reader, big.NewInt(int64(n)))
	// Since Go 1.24, this error should be impossible: crypto/rand.Read
	// never returns an error. This check is belt and suspenders.
	if err != nil {
		panic(fmt.Errorf("appleseed: crypto/rand failed: %w", err))
	}

	return int(i.Int64())
}

func appendSyllable(b []byte) []byte {
	return append(
		b,
		pickFrom(lowercaseConsonants),
		pickFrom(lowercaseVowels),
		pickFrom(lowercaseConsonants),
	)
}

func uppercaseOne(b []byte) {
	candidates := make([]int, 0, len(b))
	for i, c := range b {
		if c != 'o' {
			candidates = append(candidates, i)
		}
	}
	// This is safe since b will always have some consonants: it cannot
	// contain only 'o'.
	i := candidates[randInt(len(candidates))]
	b[i] -= 'a' - 'A'
}
