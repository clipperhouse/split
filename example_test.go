package split_test

import (
	"fmt"

	"github.com/clipperhouse/split"
)

func ExampleString() {
	text := "Hello, 世界. Nice dog! 👍🐶"
	sep := " "

	split := split.String(text, sep)

	for split.Next() {
		fmt.Println(split.Value())
	}
}
