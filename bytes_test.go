package split_test

import (
	"bytes"
	"io"
	"reflect"
	"testing"

	"github.com/clipperhouse/split"
)

const sample = " On the other hand, we denounce with righteous indignation and dislike men who are so beguiled and demoralized by the charms of pleasure of the moment, so blinded by desire, that they cannot foresee the pain and trouble that are bound to ensue; and equal blame belongs to those who fail in their duty through weakness of will, which is the same as saying through shrinking from toil and pain. These cases are perfectly simple and easy to distinguish. In a free hour, when our power of choice is untrammelled and when nothing prevents our being able to do what we like best, every pleasure is to be welcomed and every pain avoided. But in certain circumstances and owing to the claims of duty or the obligations of business it will frequently occur that pleasures have to be repudiated and annoyances accepted. The wise man therefore always holds in these matters to this principle of selection: he rejects pleasures to secure other greater pleasures, or else he endures pains to avoid worse pains. "

var sampleBytes = []byte(sample)

func BenchmarkSplitBytes(b *testing.B) {
	var sep = []byte{' '}

	b.SetBytes(int64(len(sampleBytes)))

	for i := 0; i < b.N; i++ {
		splits := split.Bytes(sampleBytes, sep)
		for splits.Next() {
			io.Discard.Write(splits.Value())
		}
	}
}

func BenchmarkBytesSplit(b *testing.B) {
	sep := []byte{' '}

	b.SetBytes(int64(len(sampleBytes)))

	for i := 0; i < b.N; i++ {
		splits := bytes.Split(sampleBytes, sep)
		for _, v := range splits {
			io.Discard.Write(v)
		}
	}
}

func TestBytes(t *testing.T) {
	for _, sep := range testSeparators {
		sepb := []byte(sep)
		for _, s := range testStrings {
			b := []byte(s)
			got := split.Bytes(b, sepb).ToArray() // this package
			expected := bytes.Split(b, sepb)      // standard library
			if !reflect.DeepEqual(got, expected) {
				t.Fatalf("\nexpected: %v,\ngot:      %v", expected, got)
			}
		}
	}
}

func TestBytesAny(t *testing.T) {
	seps := []byte(", ")

	input := []byte("Hello, how a,re, you ")
	split := split.BytesAny(input, seps)

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
