package main

func main() {
	y := 2
	_ = stackIt3(&y) // pass y down the stack as a pointer
}

//go:noinline
func stackIt3(y *int) int {
	res := *y * 2
	return res
}
