package split_test

import (
	"crypto/rand"
)

var abcd = "abc𝚫"
var emoji = "👍🐶"
var efghi = "éfghi"

var testSeparators = []string{"", " ", ",", ", ", "🐶", "𝚫", "🌏👍", abcd, emoji, efghi}

var testStrings []string

func init() {
	testStrings = append(testStrings, "") // empty
	testStrings = append(testStrings, abcd)
	testStrings = append(testStrings, emoji)
	testStrings = append(testStrings, efghi)

	// Random string
	b := make([]byte, 500)
	rand.Read(b)
	testStrings = append(testStrings, string(b))

	// Random separator
	b = make([]byte, 4)
	rand.Read(b)
	testSeparators = append(testSeparators, string(b))

	for _, sep := range testSeparators {
		testStrings = append(testStrings, sep)
		center := abcd + sep + emoji + sep + efghi
		testStrings = append(testStrings, center)
		testStrings = append(testStrings, sep+center)     // leading
		testStrings = append(testStrings, center+sep)     // trailing
		testStrings = append(testStrings, sep+center+sep) // both
	}
}
