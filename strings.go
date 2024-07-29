package split

import (
	"strings"
	"unicode/utf8"
)

type StringIterator = Iterator[string]

var stringFuncs = funcs[string]{
	Index:      strings.Index,
	IndexByte:  strings.IndexByte,
	IndexAny:   strings.IndexAny,
	DecodeRune: utf8.DecodeRuneInString,
}

// String slices s into all substrings separated by sep and returns a slice of the substrings between those separators.
//
// If s does not contain sep and sep is not empty, String returns a slice of length 1 whose only element is s.
//
// If sep is empty, String splits after each UTF-8 sequence. If both s and sep are empty, String returns an empty slice.
//
// String returns an iterator over substrings. Use `for iterator.Next()` to loop, and `iterator.Value()` to get the current substring.
func String(s string, separator string) *StringIterator {
	return split(s, separator, stringFuncs)
}

func StringAny(s string, chars string) *StringIterator {
	return splitAny(s, chars, stringFuncs)
}
