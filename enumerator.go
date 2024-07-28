package split

import "bytes"

// This Go implementation started with C# SpanSplitEnumerator<T>: https://github.com/dotnet/runtime/pull/104534

// Licensed to the .NET Foundation under one or more agreements.
// The .NET Foundation licenses this file to you under the MIT license.
// https://github.com/dotnet/runtime/blob/main/LICENSE.TXT

type mode int

const (
	none mode = iota
	singleElement
	any
	sequence
	emptySequence
)

func OnByte(input []byte, separator byte) *Enumerator {
	return &Enumerator{
		input:     input,
		separator: separator,
		mode:      singleElement,
	}
}

func OnAnyByte(input []byte, separators []byte) *Enumerator {
	return &Enumerator{
		input:      input,
		separators: separators,
		mode:       any,
	}
}

func OnByteSequence(input []byte, separators []byte) *Enumerator {
	return &Enumerator{
		input:      input,
		separators: separators,
		mode:       sequence,
	}
}

type Enumerator struct {
	input      []byte
	separator  byte
	separators []byte
	mode       mode
	start, end int
	cursor     int
}

func (en *Enumerator) Value() []byte {
	return en.input[en.start:en.end]
}

func (en *Enumerator) Next() bool {
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
