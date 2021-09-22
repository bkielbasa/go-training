package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	handler := httpHandler{
		serv: searchService{
			searches: []search{
				{
					name:    "BBC",
					feedURL: "http://feeds.bbci.co.uk/news/rss.xml",
					cache:   newCache(),
				},
				{
					name:    "NYT",
					feedURL: "https://rss.nytimes.com/services/xml/rss/nyt/World.xml",
					cache:   newCache(),
				},
				{
					name:    "CNN",
					feedURL: "http://rss.cnn.com/rss/edition.rss",
					cache:   newCache(),
				},
			},
		},
	}
	http.HandleFunc("/search", handler.searchHandler)
	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err)
}
