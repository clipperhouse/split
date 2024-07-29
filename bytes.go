package split

import (
	"bytes"
	"unicode/utf8"
)

type ByteIterator = Iterator[[]byte]

var byteFuncs = funcs[[]byte]{
	Index:      bytes.Index,
	IndexByte:  bytes.IndexByte,
	IndexAny:   bytes.IndexAny,
	DecodeRune: utf8.DecodeRune,
}

// Bytes slices s into all substrings separated by sep and returns a slice of the substrings between those separators.
//
// If s does not contain sep and sep is not empty, Bytes returns a slice of length 1 whose only element is s.
//
// If sep is empty, Bytes splits after each UTF-8 sequence. If both s and sep are empty, Bytes returns an empty slice.
//
// Bytes returns an iterator over subslices. Use `for iterator.Next()` to loop, and `iterator.Value()` to get the current subslice.
func Bytes(s []byte, sep []byte) *ByteIterator {
	return split(s, sep, byteFuncs)
}

// BytesAny splits s into subslices separated by any of the bytes found in sep
// and returns an iterator of the subslices between those separators.
// If sep is empty, BytesAny splits after each UTF-8 sequence (rune).
//
// Use `for ByteIterator.Next()` to loop, and `ByteIterator.Value()` to get the subslices.
func BytesAny(s []byte, separators []byte) *ByteIterator {
	return splitAny(s, separators, byteFuncs)
}
