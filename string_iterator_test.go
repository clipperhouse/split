package split_test

import (
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/clipperhouse/split"
)

func BenchmarkSplitString(b *testing.B) {
	var sep = " "

	b.SetBytes(int64(len(sample)))

	for i := 0; i < b.N; i++ {
		splits := split.String(sample, sep)
		for splits.Next() {
			io.WriteString(io.Discard, splits.Value())
		}
	}
}

func BenchmarkStringsSplit(b *testing.B) {
	var sep = " "

	b.SetBytes(int64(len(sample)))

	for i := 0; i < b.N; i++ {
		splits := strings.Split(sample, sep)
		for _, v := range splits {
			io.WriteString(io.Discard, v)
		}
	}
}

func TestStringOnByte(t *testing.T) {
	var got []string

	var sep byte = ' '

	input := "Hello how are you "
	split := split.StringOnByte(input, sep)
	for split.Next() {
		got = append(got, split.Value())
	}

	expected := strings.Split(input, string(sep))

	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("\nexpected: %v,\ngot:      %v", expected, got)
	}
}

func TestString(t *testing.T) {
	var sep = ", "

	input := "Hello, how a,re, you "
	split := split.String(input, sep)

	var got []string
	for split.Next() {
		got = append(got, split.Value())
	}

	expected := strings.Split(input, sep)

	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("\nexpected: %v,\ngot:      %v", expected, got)
	}
}

func TestStringOnAnyChar(t *testing.T) {
	var sep = ", "

	input := "Hello, how a,re, you "
	split := split.StringOnAnyChar(input, sep)

	var got []string
	for split.Next() {
		got = append(got, split.Value())
	}

	expected := []string{
		"Hello",
		"",
		"how",
		"a",
		"re",
		"",
		"you",
		"",
	}

	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("\nexpected: %v,\ngot:      %v", expected, got)
	}
}
