package split

import (
	"bytes"
	"strings"
	"unicode/utf8"
)

type ByteSeq interface {
	~string | ~[]byte
}

type funcs[T ByteSeq] struct {
	Index      func(s T, sep T) int
	IndexByte  func(b T, c byte) int
	IndexAny   func(s T, chars string) int
	DecodeRune func(p T) (r rune, size int)
}

var stringFuncs = funcs[string]{
	Index:      strings.Index,
	IndexByte:  strings.IndexByte,
	IndexAny:   strings.IndexAny,
	DecodeRune: utf8.DecodeRuneInString,
}

var byteFuncs = funcs[[]byte]{
	Index:      bytes.Index,
	IndexByte:  bytes.IndexByte,
	IndexAny:   bytes.IndexAny,
	DecodeRune: utf8.DecodeRune,
}

type iterator[T ByteSeq] struct {
	funcs      funcs[T]
	input      T
	separator  byte
	separators T
	mode       mode
	start, end int
	cursor     int
}

func (it *iterator[T]) Value() T {
	return it.input[it.start:it.end]
}

func (it *iterator[T]) Next() bool {
	var index int
	var separatorLength = 1
	var slice = it.input[it.cursor:]

	switch it.mode {
	case none:
		return false
	case singleElement:
		index = it.funcs.IndexByte(slice, it.separator)
	case any:
		index = it.funcs.IndexAny(slice, string(it.separators))
	case sequence:
		index = it.funcs.Index(slice, it.separators)
		separatorLength = len(it.separators)
	case emptySequence:
		_, index = it.funcs.DecodeRune(slice)
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

func (it *iterator[T]) ToArray() []T {
	var result []T

	for it.Next() {
		result = append(result, it.Value())
	}

	return result
}
