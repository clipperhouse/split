package split

type mode int

const (
	none mode = iota
	singleElement
	any
	sequence
	emptySequence
)
