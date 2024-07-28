package split

import (
	"bytes"
	"unicode/utf8"
)

// This Go implementation started with C# SpanSplitEnumerator<T>: https://github.com/dotnet/runtime/pull/104534

// Licensed to the .NET Foundation under one or more agreements.
// The .NET Foundation licenses this file to you under the MIT license.
// https://github.com/dotnet/runtime/blob/main/LICENSE.TXT

// Bytes splits s into subslices separated by sep and
// returns an iterator of the subslices between those separators.
// If sep is empty, Bytes splits after each UTF-8 sequence (rune).
//
// Use `for ByteIterator.Next()` to loop, and `ByteIterator.Value()` to get the subslices.
func Bytes(s []byte, sep []byte) *ByteIterator {
	var mode = sequence
	if len(sep) == 0 {
		mode = emptySequence
	}

	return &ByteIterator{
		input:      s,
		separators: sep,
		mode:       mode,
	}
}

func BytesOnByte(input []byte, separator byte) *ByteIterator {
	return &ByteIterator{
		input:     input,
		separator: separator,
		mode:      singleElement,
	}
}

// BytesOnAny splits s into subslices separated by any of the bytes found in sep
// and returns an iterator of the subslices between those separators.
// If sep is empty, BytesOnAny splits after each UTF-8 sequence (rune).
//
// Use `for ByteIterator.Next()` to loop, and `ByteIterator.Value()` to get the subslices.
func BytesOnAny(input []byte, separators []byte) *ByteIterator {
	var mode = any
	if len(separators) == 0 {
		mode = emptySequence
	}

	return &ByteIterator{
		input:      input,
		separators: separators,
		mode:       mode,
	}
}

type ByteIterator struct {
	input      []byte
	separator  byte
	separators []byte
	mode       mode
	start, end int
	cursor     int
}

func (en *ByteIterator) Value() []byte {
	return en.input[en.start:en.end]
}

func (en *ByteIterator) Next() bool {
	var index int
	var separatorLength = 1
	var slice = en.input[en.cursor:]

	switch en.mode {
	case none:
		return false
	case singleElement:
		index = bytes.IndexByte(slice, en.separator)
	case any:
		index = bytes.IndexAny(slice, string(en.separators))
	case sequence:
		index = bytes.Index(slice, en.separators)
		separatorLength = len(en.separators)
	case emptySequence:
		_, index = utf8.DecodeRune(slice)
		if index == 0 {
			return false
		}
		separatorLength = 0
	}

	en.start = en.cursor

	if index >= 0 {
		en.end = en.start + index
		en.cursor = en.end + separatorLength
	} else {
		en.cursor = len(en.input)
		en.end = len(en.input)
		en.mode = none
	}

	return true
}
