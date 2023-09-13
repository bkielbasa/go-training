package main

// Refactor the following code to use generics:
type List struct {
	val  int
	Next *List
}

func main() {
	list := &List{val: 1, Next: &List{val: 2, Next: &List{val: 3, Next: &List{val: 4, Next: &List{val: 5}}}}}
	for list != nil {
		println(list.val)
		list = list.Next
	}
}
