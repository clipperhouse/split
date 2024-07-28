package split

import "strings"

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
	return &StringIterator{
		input:      input,
		separators: separator,
		mode:       sequence,
	}
}

func StringOnAnyChar(input string, chars string) *StringIterator {
	return &StringIterator{
		input:      input,
		separators: chars,
		mode:       any,
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

func (en *StringIterator) Value() string {
	return en.input[en.start:en.end]
}

func (en *StringIterator) Next() bool {
	var index int
	var separatorLength = 1
	var slice = en.input[en.cursor:]

	switch en.mode {
	case none:
		return false

	case singleElement:
		index = strings.IndexByte(slice, en.separator)
	case any:
		index = strings.IndexAny(slice, en.separators)
	case sequence:
		index = strings.Index(slice, en.separators)
		separatorLength = len(en.separators)
	case emptySequence:
		index = -1
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
