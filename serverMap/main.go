package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", handler)
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8080", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {
	response := []byte("pong")
	_, err := w.Write(response)
	if err != nil {

	}
}

func hello(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Print(err)
		return
	}

	v := map[string]string{}
	err = json.Unmarshal(body, &v)
	if err != nil {
		fmt.Print(err)
		return
	}
	name := v["name"]

	v2 := map[string]string{"message": "Hello " + name + ", How are you?"}

	body, err = json.Marshal(v2)
	if err != nil {
		fmt.Print(err)
		return

	}
	_, err = w.Write(body)
	if err != nil {
		fmt.Print(err)
		return
	}

}
