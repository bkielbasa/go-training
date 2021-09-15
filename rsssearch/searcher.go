package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/mmcdole/gofeed"
)

type searchResult struct {
	Title string
	URL   string
}

type search struct {
	name    string
	feedURL string
}

func (s search) Search(ctx context.Context, term string) ([]searchResult, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURLWithContext(s.feedURL, ctx)
	if err != nil {
		return nil, fmt.Errorf("cannot fetch data from %s: %w", s.name, err)
	}

	results := []searchResult{}
	for _, f := range feed.Items {
		if strings.Contains(f.Title, term) {
			results = append(results, searchResult{
				Title: f.Title,
				URL:   f.Link,
			})

			if len(results) == 10 {
				break
			}
		}
	}

	return results, nil
}

type searchService struct {
	searches []search
}

func (s searchService) Search(ctx context.Context, term string) ([]searchResult, string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	var results []searchResult
	var fastest string

	for i := 0; i < len(s.searches); i++ {
		go func(s search) {
			r, err := s.Search(ctx, term)
			if err != nil {
				log.Print(err)
				return
			}

			results = r
			fastest = s.name
			cancel()
		}(s.searches[i])
	}

	<-ctx.Done()
	if results != nil {
		return results, fastest, nil
	}

	return nil, "", ctx.Err()
}
