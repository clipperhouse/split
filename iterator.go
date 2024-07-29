package split

type ByteSeq interface {
	~string | ~[]byte
}

type funcs[T ByteSeq] struct {
	Index      func(s T, sep T) int
	IndexByte  func(s T, c byte) int
	IndexAny   func(s T, chars string) int
	DecodeRune func(p T) (r rune, size int)
}

func split[T ByteSeq](s T, sep T, funcs funcs[T]) *iterator[T] {
	var mode = sequence
	if len(s) == 0 {
		mode = done
	} else if len(sep) == 0 {
		mode = emptySequence
	}

	return &iterator[T]{
		funcs:      funcs,
		input:      s,
		separators: sep,
		mode:       mode,
	}
}

func splitAny[T ByteSeq](s T, separators T, funcs funcs[T]) *iterator[T] {
	var mode = any
	if len(s) == 0 {
		mode = done
	} else if len(separators) == 0 {
		mode = emptySequence
	}

	return &iterator[T]{
		funcs:      funcs,
		input:      s,
		separators: separators,
		mode:       mode,
	}

}

type iterator[T ByteSeq] struct {
	funcs[T]
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
	case done:
		return false
	case singleElement:
		index = it.IndexByte(slice, it.separator)
	case any:
		index = it.IndexAny(slice, string(it.separators))
	case sequence:
		index = it.Index(slice, it.separators)
		separatorLength = len(it.separators)
	case emptySequence:
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

func (it *iterator[T]) ToArray() []T {
	var result []T

	for it.Next() {
		result = append(result, it.Value())
	}

	return result
}
