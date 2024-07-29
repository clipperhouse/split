package split

import (
	"bytes"
	"unicode/utf8"
)

// This Go implementation started with C# SpanSplitEnumerator<T>: https://github.com/dotnet/runtime/pull/104534

// Licensed to the .NET Foundation under one or more agreements.
// The .NET Foundation licenses this file to you under the MIT license.
// https://github.com/dotnet/runtime/blob/main/LICENSE.TXT

type ByteIterator = iterator[[]byte]

var byteFuncs = funcs[[]byte]{
	Index:      bytes.Index,
	IndexByte:  bytes.IndexByte,
	IndexAny:   bytes.IndexAny,
	DecodeRune: utf8.DecodeRune,
}

// Bytes splits s into subslices separated by sep and
// returns an iterator of the subslices between those separators.
// If sep is empty, Bytes splits after each UTF-8 sequence (rune).
//
// Use `for iterator.Next()` to loop, and `iterator.Value()` to get the subslices.
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
