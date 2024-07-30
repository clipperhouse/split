package split_test

import (
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/clipperhouse/split"
)

var bench = "Hello, 世界. Nice dog! 👍🐶 Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."

func BenchmarkSplitString(b *testing.B) {
	sep := " "

	b.SetBytes(int64(len(bench)))

	for i := 0; i < b.N; i++ {
		split := split.String(bench, sep)

		for split.Next() {
			io.WriteString(io.Discard, split.Value())
		}
	}
}

func BenchmarkSplitStringToArray(b *testing.B) {
	sep := " "

	b.SetBytes(int64(len(bench)))

	for i := 0; i < b.N; i++ {
		split := split.String(bench, sep).ToArray()

		for _, v := range split {
			io.WriteString(io.Discard, v)
		}
	}
}

func BenchmarkStringsSplit(b *testing.B) {
	sep := " "

	b.SetBytes(int64(len(bench)))

	for i := 0; i < b.N; i++ {
		split := strings.Split(bench, sep)
		for _, v := range split {
			io.WriteString(io.Discard, v)
		}
	}
}

func TestStrings(t *testing.T) {
	for _, sep := range testSeparators {
		for _, s := range testStrings {
			got := split.String(s, sep).ToArray() // this package
			expected := strings.Split(s, sep)     // standard library
			if !reflect.DeepEqual(got, expected) {
				t.Fatalf("\nFor input %q, separator %q\nexpected: %q,\ngot:      %q", s, sep, expected, got)
			}
		}
	}
}

func TestStringAny(t *testing.T) {
	var sep = ", "

	input := "Hello, how a,re, you "
	split := split.StringAny(input, sep)

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

func TestResetString(t *testing.T) {
	text := "Hello, how are you."
	s1 := split.String(text, " ")
	a1 := s1.ToArray()
	a2 := s1.ToArray()

	if reflect.DeepEqual(a1, a2) {
		t.Fatal("They should not be equal, as s1 has not been reset")
	}

	s1.Reset()
	a3 := s1.ToArray()

	if !reflect.DeepEqual(a1, a3) {
		t.Fatal("They should be equal, as s1 has been reset")
	}
}
