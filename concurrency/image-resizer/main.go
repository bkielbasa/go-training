package main

import (
	"log"
	"net/http"

	lru "github.com/hashicorp/golang-lru"
)

type service struct {
	cache *lru.Cache
}

type resizeRequest struct {
	URLs   []string `json:"urls"`
	Width  uint     `json:"width"`
	Height uint     `json:"height"`
}

type resizeResult struct {
	Result string `json:"result"`
	URL    string `json:"url,omitempty"`
	Cached bool   `json:"cached"`
}

const (
	proto    = "http://"
	hostport = "localhost:8080"
	success  = "success"
	failure  = "failure"
)

func main() {
	cache, err := lru.New(1024)
	if err != nil {
		log.Panicf("Faild to create cache: %v", err)
	}

	svc := &service{
		cache: cache,
	}

	mux := http.NewServeMux()
	mux.Handle("/v1/resize", svc.resizeHandler())
	mux.Handle("/v1/image/", svc.getImageHandler())
	address := hostport

	log.Print("Listening on ", hostport)
	// When running on docker mac, can't listen only on localhost
	panic(http.ListenAndServe(address, mux))
}
