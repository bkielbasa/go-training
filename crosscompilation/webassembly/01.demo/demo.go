// main.go
package main

import (
	"errors"
	"syscall/js"
)

func add(this js.Value, p []js.Value) interface{} {
	if len(p) != 2 {
		return errors.New("invalid number of arguments")
	}
	return js.ValueOf(p[0].Int() + p[1].Int())
}

func main() {
	c := make(chan struct{}, 0)

	js.Global().Set("add", js.FuncOf(add))

	<-c
}
