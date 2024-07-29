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

type SplitTest struct {
	s   string
	sep string
	n   int
	a   []string
}

var abcd = "abcd"
var faces = "☺☻☹"
var commas = "1,2,3,4"
var dots = "1....2....3....4"

var splittests = []SplitTest{
	{"", "a", -1, []string{}},
	{"", "", -1, []string{}},
	// {abcd, "", 2, []string{"a", "bcd"}},
	// {abcd, "", 4, []string{"a", "b", "c", "d"}},
	{abcd, "", -1, []string{"a", "b", "c", "d"}},
	{faces, "", -1, []string{"☺", "☻", "☹"}},
	// {faces, "", 3, []string{"☺", "☻", "☹"}},
	// {faces, "", 17, []string{"☺", "☻", "☹"}},
	{"☺�☹", "", -1, []string{"☺", "�", "☹"}},
	// {abcd, "a", 0, nil},
	{abcd, "a", -1, []string{"", "bcd"}},
	{abcd, "z", -1, []string{"abcd"}},
	{commas, ",", -1, []string{"1", "2", "3", "4"}},
	{dots, "...", -1, []string{"1", ".2", ".3", ".4"}},
	{faces, "☹", -1, []string{"☺☻", ""}},
	{faces, "~", -1, []string{faces}},
	// {"1 2 3 4", " ", 3, []string{"1", "2", "3 4"}},
	// {"1 2", " ", 3, []string{"1", "2"}},
	// {"", "T", math.MaxInt / 4, []string{""}},
	{"\xff-\xff", "", -1, []string{"\xff", "-", "\xff"}},
	{"\xff-\xff", "-", -1, []string{"\xff", "\xff"}},
}

func eq(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestSplit(t *testing.T) {
	for i, tt := range splittests {
		a := split.String(tt.s, tt.sep).ToArray()
		if !eq(a, tt.a) {
			t.Errorf("%d: Split(%q, %q, %d) = %v; want %v", i, tt.s, tt.sep, tt.n, a, tt.a)
			continue
		}
		if tt.n == 0 {
			continue
		}
		s := strings.Join(a, tt.sep)
		if s != tt.s {
			t.Errorf("Join(Split(%q, %q, %d), %q) = %q", tt.s, tt.sep, tt.n, tt.sep, s)
		}
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
