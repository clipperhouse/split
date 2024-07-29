package split

type mode int

const (
	done mode = iota
	singleElement
	any
	sequence
	emptySequence
)
