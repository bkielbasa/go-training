package main

func main() {
	_ = stackIt()
}

//go:noinline
func stackIt() int {
	y := 2
	return y * 2
}
