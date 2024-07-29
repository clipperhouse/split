// This Go implementation started with C# SpanSplitEnumerator<T>: https://github.com/dotnet/runtime/pull/104534

// Licensed to the .NET Foundation under one or more agreements.
// The .NET Foundation licenses this file to you under the MIT license.
// https://github.com/dotnet/runtime/blob/main/LICENSE.TXT

package split

type seq interface {
	~string | ~[]byte
}

type mode byte

const (
	done mode = iota
	any
	sequence
	emptySeparator
)

type funcs[T seq] struct {
	Index      func(s T, sep T) int
	IndexByte  func(s T, c byte) int
	IndexAny   func(s T, chars string) int
	DecodeRune func(p T) (r rune, size int)
}

func split[T seq](s T, sep T, funcs funcs[T]) *Iterator[T] {
	var mode = sequence
	if len(sep) == 0 {
		mode = emptySeparator
	}

	return &Iterator[T]{
		funcs:      funcs,
		input:      s,
		separators: sep,
		mode:       mode,
	}
}

func splitAny[T seq](s T, separators T, funcs funcs[T]) *Iterator[T] {
	var mode = any
	if len(separators) == 0 {
		mode = emptySeparator
	}

	return &Iterator[T]{
		funcs:      funcs,
		input:      s,
		separators: separators,
		mode:       mode,
	}

}

// Iterator is an iterator over subslices of `[]byte` or `string`. See the `Next` and `Value` methods.
type Iterator[T seq] struct {
	funcs[T]
	input      T
	separators T
	mode       mode
	start, end int
	cursor     int
}

// Value retrieves the value of the current subslice.
func (it *Iterator[T]) Value() T {
	return it.input[it.start:it.end]
}

// Next tests whether there are any remaining subslices.
//
// Use a `for iterator.Next()` loop, and retrieve the current subslice with `iterator.Value()`.
func (it *Iterator[T]) Next() bool {
	var index int
	var separatorLength = 1
	var slice = it.input[it.cursor:]

	switch it.mode {
	case done:
		return false
	case any:
		index = it.IndexAny(slice, string(it.separators))
	case sequence:
		index = it.Index(slice, it.separators)
		separatorLength = len(it.separators)
	case emptySeparator:
		_, index = it.DecodeRune(slice)
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
		it.mode = done
	}

	return true
}

// ToArray collects all the subslices into an array.
//
// This is a convenience method, and the result should identical to strings|bytes.Split from the standard library.
// You should just use strings|bytes.Split if your goal is an array of results.
func (it *Iterator[T]) ToArray() []T {
	var result []T = make([]T, 0, len(it.input)/4)

	for it.Next() {
		result = append(result, it.Value())
	}

	return result
}
