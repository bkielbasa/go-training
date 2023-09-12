package allocations

type I interface {
	M(*int)
}

type T struct{}

func (T) M(*int) {}

var t T
var i I = t

// The value parts referenced by an argument will escape to heap if the argument is passed to interface method calls
func foo() {
	var x int // does not escape
	t.M(&x)
	var y int // moved to heap: y
	i.M(&y)
}
