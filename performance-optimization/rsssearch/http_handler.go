package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type searchRequest struct {
	Term string
}

type searchResponse struct {
	Source  string
	Results []searchResult
}

type httpHandler struct {
	serv searchService
}

func (h httpHandler) searchHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	req := searchRequest{}
	if err = json.Unmarshal(body, &req); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.serv.Search(r.Context(), req.Term)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if body, err = json.Marshal(result); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("content-type", "application/json")
	w.Write(body)
}
