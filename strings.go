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

// String slices s into all substrings separated by sep.
//
// If s does not contain sep and sep is not empty, String returns a slice of length 1 whose only element is s.
//
// If sep is empty, String splits after each UTF-8 sequence. If both s and sep are empty, String returns an empty slice.
//
// String returns an iterator over substrings. Use [Iterator.Next] to loop, and [Iterator.Value] to get the current substring.
func String(s string, separator string) *StringIterator {
	return split(s, separator, stringFuncs)
}

// StringAny slices s into all substrings separated by any Unicode code point in chars.
//
// If s does not contain any Unicode code point in chars, and chars is not empty, StringAny returns a slice of length 1 whose only element is s.
//
// If chars is empty, StringAny splits after each UTF-8 sequence. If both s and chars are empty, StringAny returns an empty slice.
//
// StringAny returns an iterator over substrings. Use [Iterator.Next] to loop, and [Iterator.Value] to get the current substrings.
func StringAny(s string, chars string) *StringIterator {
	return splitAny(s, chars, stringFuncs)
}
