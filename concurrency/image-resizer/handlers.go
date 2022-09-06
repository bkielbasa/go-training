package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (s *service) resizeHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Expecting POST request"))
			return
		}

		request := resizeRequest{}
		err := json.NewDecoder(io.LimitReader(r.Body, 8*1024)).Decode(&request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Failed to parse request"))
			return
		}

		results, err := s.processResizes(request)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed to process request"))
			return
		}

		data, err := json.Marshal(results)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed to marshal response"))
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Header().Add("content-type", "application/json")
		w.Write(data)
	})
}

func (s *service) getImageHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("fetching ", r.URL.String())
		data, ok := s.cache.Get(r.URL.String())
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Add("content-type", "image/jpeg")
		w.Write(data.([]byte))
	})
}
