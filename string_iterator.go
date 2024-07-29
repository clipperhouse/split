package split

// This Go implementation started with C# SpanSplitEnumerator<T>: https://github.com/dotnet/runtime/pull/104534

// Licensed to the .NET Foundation under one or more agreements.
// The .NET Foundation licenses this file to you under the MIT license.
// https://github.com/dotnet/runtime/blob/main/LICENSE.TXT

type StringIterator = iterator[string]

func StringOnByte(s string, separator byte) *StringIterator {
	return &StringIterator{
		funcs:     stringFuncs,
		input:     s,
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
		funcs:      stringFuncs,
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
		funcs:      stringFuncs,
		input:      input,
		separators: chars,
		mode:       mode,
	}
}
