package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

var glob string

func main() {
	http.HandleFunc("/", getHello)

	err := http.ListenAndServe(":3333", nil)
	fmt.Printf("err: %s", err)
}

func getHello(w http.ResponseWriter, r *http.Request) {
	sizeS := r.URL.Query().Get("size")
	size, _ := strconv.Atoi(sizeS)
	s := newService(size)

	io.WriteString(w, s.cache)
}

type service struct {
	cache string
	glob  interface{}
}

//go:noinline
func newService(size int) *service {
	return &service{
		glob:  make([]byte, size),
		cache: "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.",
	}
}
