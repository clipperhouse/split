package split

// This Go implementation started with C# SpanSplitEnumerator<T>: https://github.com/dotnet/runtime/pull/104534

// Licensed to the .NET Foundation under one or more agreements.
// The .NET Foundation licenses this file to you under the MIT license.
// https://github.com/dotnet/runtime/blob/main/LICENSE.TXT

type ByteIterator = iterator[[]byte]

// Bytes splits s into subslices separated by sep and
// returns an iterator of the subslices between those separators.
// If sep is empty, Bytes splits after each UTF-8 sequence (rune).
//
// Use `for iterator.Next()` to loop, and `iterator.Value()` to get the subslices.
func Bytes(s []byte, sep []byte) *ByteIterator {
	var mode = sequence
	if len(s) == 0 || len(sep) == 0 {
		mode = emptySequence
	}

	return &ByteIterator{
		funcs:      byteFuncs,
		input:      s,
		separators: sep,
		mode:       mode,
	}
}

func BytesOnByte(s []byte, separator byte) *ByteIterator {
	return &ByteIterator{
		funcs:     byteFuncs,
		input:     s,
		separator: separator,
		mode:      singleElement,
	}
}

// BytesOnAny splits s into subslices separated by any of the bytes found in sep
// and returns an iterator of the subslices between those separators.
// If sep is empty, BytesOnAny splits after each UTF-8 sequence (rune).
//
// Use `for ByteIterator.Next()` to loop, and `ByteIterator.Value()` to get the subslices.
func BytesOnAny(s []byte, separators []byte) *ByteIterator {
	var mode = any
	if len(s) == 0 || len(separators) == 0 {
		mode = emptySequence
	}

	return &ByteIterator{
		funcs:      byteFuncs,
		input:      s,
		separators: separators,
		mode:       mode,
	}
}
