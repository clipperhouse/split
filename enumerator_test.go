package split

import (
	"bytes"
	"reflect"
	"testing"
)

func TestOnByte(t *testing.T) {
	var got [][]byte

	var sep byte = ' '

	input := []byte("Hello how are you ")
	split := OnByte(input, sep)
	for split.Next() {
		got = append(got, split.Value())
	}

	byts := []byte{sep}
	expected := bytes.Split(input, byts)

	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("\nexpected: %v,\ngot:      %v", expected, got)
	}
}

func TestOnByteSequence(t *testing.T) {
	var sep = ", "
	byts := []byte(sep)

	input := []byte("Hello, how a,re, you ")
	split := OnByteSequence(input, byts)

	var got [][]byte
	for split.Next() {
		got = append(got, split.Value())
	}

	expected := bytes.Split(input, byts)

	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("\nexpected: %v,\ngot:      %v", expected, got)
	}
}

func TestOnAnyByte(t *testing.T) {
	var sep = ", "
	byts := []byte(sep)

	input := []byte("Hello, how a,re, you ")
	split := OnAnyByte(input, byts)

	var got [][]byte
	for split.Next() {
		got = append(got, split.Value())
	}

	expected := [][]byte{
		[]byte("Hello"),
		[]byte(""),
		[]byte("how"),
		[]byte("a"),
		[]byte("re"),
		[]byte(""),
		[]byte("you"),
		[]byte(""),
	}

	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("\nexpected: %v,\ngot:      %v", expected, got)
	}
}
