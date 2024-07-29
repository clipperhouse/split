package split

import (
	"strings"
	"unicode/utf8"
)

// This Go implementation started with C# SpanSplitEnumerator<T>: https://github.com/dotnet/runtime/pull/104534

// Licensed to the .NET Foundation under one or more agreements.
// The .NET Foundation licenses this file to you under the MIT license.
// https://github.com/dotnet/runtime/blob/main/LICENSE.TXT

type StringIterator = iterator[string]

var stringFuncs = funcs[string]{
	Index:      strings.Index,
	IndexByte:  strings.IndexByte,
	IndexAny:   strings.IndexAny,
	DecodeRune: utf8.DecodeRuneInString,
}

func String(s string, separator string) *StringIterator {
	return split(s, separator, stringFuncs)
}

func StringAny(s string, chars string) *StringIterator {
	return splitAny(s, chars, stringFuncs)
}
