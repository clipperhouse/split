package split

import (
	"strings"
	"unicode/utf8"
)

// This Go implementation started with C# SpanSplitEnumerator<T>: https://github.com/dotnet/runtime/pull/104534

// Licensed to the .NET Foundation under one or more agreements.
// The .NET Foundation licenses this file to you under the MIT license.
// https://github.com/dotnet/runtime/blob/main/LICENSE.TXT

func StringOnByte(input string, separator byte) *StringIterator {
	return &StringIterator{
		input:     input,
		separator: separator,
		mode:      singleElement,
	}
}

func String(input string, separator string) *StringIterator {
	var mode = sequence
	if len(input) == 0 || len(separator) == 0 {
		mode = emptySequence
	}
	return &StringIterator{
		input:      input,
		separators: separator,
		mode:       mode,
	}
}

func StringOnAnyChar(input string, chars string) *StringIterator {
	var mode = any
	if len(input) == 0 || len(chars) == 0 {
		mode = emptySequence
	}
	return &StringIterator{
		input:      input,
		separators: chars,
		mode:       mode,
	}
}

type StringIterator struct {
	input      string
	separator  byte
	separators string
	mode       mode
	start, end int
	cursor     int
}

func (it *StringIterator) Value() string {
	return it.input[it.start:it.end]
}

func (it *StringIterator) Next() bool {
	var index int
	var separatorLength = 1
	var slice = it.input[it.cursor:]

	switch it.mode {
	case none:
		return false

	case singleElement:
		index = strings.IndexByte(slice, it.separator)
	case any:
		index = strings.IndexAny(slice, it.separators)
	case sequence:
		index = strings.Index(slice, it.separators)
		separatorLength = len(it.separators)
	case emptySequence:
		_, index = utf8.DecodeRuneInString(slice)
		if index == 0 {
			return false
		}
		separatorLength = 0
	}

	it.start = it.cursor

	if index >= 0 {
		it.end = it.start + index
		it.cursor = it.end + separatorLength
	} else {
		it.cursor = len(it.input)
		it.end = len(it.input)
		it.mode = none
	}

	return true
}

func (it *StringIterator) ToArray() []string {
	var result []string

	for it.Next() {
		result = append(result, it.Value())
	}

	return result
}
