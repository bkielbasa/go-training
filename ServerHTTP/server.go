package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	_ "github.com/lib/pq"
)

var mx sync.Mutex
var Names []string

type Request struct {
	Name string
}
type Response struct {
	Hello string
}

const (
	host     = "localhost"
	port     = 5432
	user     = "kurs"
	password = "kamil12"
	dbname   = "names"
)

func main() {
	http.HandleFunc("/ping", handler)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/names", names)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}

}

func handler(w http.ResponseWriter, r *http.Request) {

	resp := []byte("pong")
	_, err := w.Write(resp)
	if err != nil {

	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	mx.Lock()

	body, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Print(err)
	}
	v := Request{}
	err = json.Unmarshal(body, &v)
	if err != nil {
		fmt.Print(err)
		return
	}
	mx.Unlock()
	name := v.Name
	v2 := Response{}
	//Names = append(Names, name)
	v2.Hello = "Hello " + name + ", how are you?"
	mx.Lock()
	fmt.Print(name)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err)
	}
	sqlStatement := `INSERT INTO names values ($1)`
	_, err = db.Exec(sqlStatement, name)
	if err != nil {
		// handle err
	}

	defer db.Close()

	body, err = json.Marshal(v2)
	fmt.Print(v2)
	if err != nil {
		fmt.Print(err)
		return

	}
	_, err = w.Write(body)
	if err != nil {
		fmt.Print(err)
		return
	}
	mx.Unlock()
}

func names(w http.ResponseWriter, r *http.Request) {
	mx.Lock()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err)
	}
	rows, err := db.Query("SELECT name from names")
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var name string

		err = rows.Scan(&name)
		if err != nil {
			// handle this error
			panic(err)
		}
		Names = append(Names, name)
		fmt.Println(name)
	}

	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	body, err := json.Marshal(Names)
	if err != nil {
		fmt.Print(err)
		return

	}
	_, err = w.Write(body)
	if err != nil {
		fmt.Print(err)
		return
	}
	mx.Unlock()
	defer db.Close()

}
