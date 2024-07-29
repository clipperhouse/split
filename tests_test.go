package split_test

var abcd = "abc𝚫"
var emoji = "👍🐶"
var efghi = "éfghi"

var testSeparators = []string{" ", ",", ", ", "🐶", "𝚫", "🌏👍", ""}

var testStrings []string

func init() {
	testStrings = append(testStrings, "") // empty

	for _, sep := range testSeparators {
		center := abcd + sep + emoji + sep + efghi
		testStrings = append(testStrings, sep)            // leading
		testStrings = append(testStrings, sep+center)     // leading
		testStrings = append(testStrings, center+sep)     // trailing
		testStrings = append(testStrings, sep+center+sep) // both
	}
}
