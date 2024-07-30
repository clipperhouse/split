package split_test

import (
	"reflect"
	"testing"

	"github.com/clipperhouse/split"
)

func TestReset(t *testing.T) {
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
