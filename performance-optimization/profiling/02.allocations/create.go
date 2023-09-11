package allocations

type BigStruct struct {
	A, B, C int
	D, E, F string
	G, H, I bool
}

//go:noinline
func CreateCopy() BigStruct {
	return BigStruct{
		A: 123, B: 456, C: 789,
		D: "ABC", E: "DEF", F: "HIJ",
		G: true, H: true, I: true,
	}
}

//go:noinline
func CreatePointer() *BigStruct {
	return &BigStruct{
		A: 123, B: 456, C: 789,
		D: "ABC", E: "DEF", F: "HIJ",
		G: true, H: true, I: true,
	}
}
