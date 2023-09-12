package allocations

// A local variables declared in a loop will escape to heap if it is referenced by a value out of the loop
func allocs() {
	var x *int
	for {
		var n = 1 // moved to heap: n
		x = &n
		break
	}
	_ = x
}
