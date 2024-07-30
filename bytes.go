package split

import (
	"bytes"
	"unicode/utf8"
)

type ByteIterator = Iterator[[]byte]

var byteFuncs = funcs[[]byte]{
	index:      bytes.Index,
	indexAny:   bytes.IndexAny,
	decodeRune: utf8.DecodeRune,
}

// Bytes slices s into all subslices separated by sep.
//
// If s does not contain sep and sep is not empty, Bytes returns a slice of length 1 whose only element is s.
//
// If sep is empty, Bytes splits after each UTF-8 sequence. If both s and sep are empty, Bytes returns an empty slice.
//
// Bytes returns an iterator over subslices. Use [Iterator.Next] to loop, and [Iterator.Value] to get the current subslice.
func Bytes(s []byte, sep []byte) *ByteIterator {
	return split(s, sep, byteFuncs)
}

// BytesAny slices s into all subslices separated by any bytes in separators.
//
// If s does not contain separators and separators is not empty, BytesAny returns a slice of length 1 whose only element is s.
//
// If separators is empty, BytesAny splits after each UTF-8 sequence. If both s and separators are empty, BytesAny returns an empty slice.
//
// BytesAny returns an iterator over subslices. Use [Iterator.Next] to loop, and [Iterator.Value] to get the current subslice.
func BytesAny(s []byte, separators []byte) *ByteIterator {
	return splitAny(s, separators, byteFuncs)
}
