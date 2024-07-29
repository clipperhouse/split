package split_test

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

	for _, sep := range testSeparators {
		testStrings = append(testStrings, sep)
		center := abcd + sep + emoji + sep + efghi
		testStrings = append(testStrings, center)
		testStrings = append(testStrings, sep+center)     // leading
		testStrings = append(testStrings, center+sep)     // trailing
		testStrings = append(testStrings, sep+center+sep) // both
	}
}
