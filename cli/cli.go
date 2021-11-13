package main

import (
	"flag"
)

func main() {
	var b bool
	var s string

	b2 := flag.Bool("boolean", false, "fsjlfjslakfjdslk")
	flag.BoolVar(&b, "bool2", false, "usage")
	flag.StringVar(&s, "name", "", "my usage")
	flag.Parse()

	if b2 == nil {
	}
}
