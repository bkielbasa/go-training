package main

import (
	"net/http"
)

func main() {
	handler := httpHandler{
		serv: searchService{
			searches: []search{
				{
					name:    "BBC",
					feedURL: "http://feeds.bbci.co.uk/news/rss.xml",
				},
				{
					name:    "NYT",
					feedURL: "https://rss.nytimes.com/services/xml/rss/nyt/World.xml",
				},
				{
					name:    "CNN",
					feedURL: "http://rss.cnn.com/rss/edition.rss",
				},
			},
		},
	}
	http.HandleFunc("/search", handler.searchHandler)
	http.ListenAndServe(":8080", nil)
}
