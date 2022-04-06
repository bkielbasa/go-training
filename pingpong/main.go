package main

import (
	"encoding/json"
	"fmt"
)

type request struct {
	CamelCase int `json:"camelCase"`
	B         int `json:"nawa_nazwa"`
}

func main() {
	r := request{}
	r.CamelCase = 1
	b, _ := json.Marshal(r)
	fmt.Printf("%s\n", string(b))
}
