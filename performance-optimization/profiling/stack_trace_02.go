package main

func main() {
	example(false, false, false, 25)
}

//go:noinline
func example(b1, b2, b3 bool, i uint8) error {
	panic("Want stack trace")
}
