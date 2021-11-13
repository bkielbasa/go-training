package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"
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
	cache   *cache
}

func (s search) Search(ctx context.Context, term string) ([]searchResult, error) {
	data, found := s.cache.Get(s.feedURL)
	var feed *gofeed.Feed
	var err error

	if found {
		feed = data.(*gofeed.Feed)
	} else {
		fp := gofeed.NewParser()
		feed, err = fp.ParseURLWithContext(s.feedURL, ctx)
	}

	if err != nil {
		return nil, fmt.Errorf("cannot fetch data from %s: %w", s.name, err)
	}
	s.cache.Set(s.feedURL, feed)

	results := []searchResult{}
	for _, f := range feed.Items {
		if strings.Contains(strings.ToLower(f.Title), term) || strings.Contains(strings.ToLower(f.Description), term) {
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

type cache struct {
	mu       sync.Mutex
	cache    map[string]interface{}
	timeouts map[string]time.Time
}

func newCache() *cache {
	c := &cache{
		cache:    map[string]interface{}{},
		timeouts: map[string]time.Time{},
	}
	go c.start()
	return c
}

func (c *cache) start() {
	ticker := time.NewTicker(10 * time.Millisecond)

	for {
		<-ticker.C
		c.invalidateExpiredCaches()
	}
}

func (c *cache) invalidateExpiredCaches() {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()

	for term, timeout := range c.timeouts {
		if timeout.Before(now) {
			delete(c.cache, term)
			delete(c.timeouts, term)
		}
	}
}

func (c *cache) Get(term string) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	res, ok := c.cache[term]
	return res, ok
}

func (c *cache) Set(term string, data interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[term] = data
	c.timeouts[term] = time.Now().Add(time.Second * 30)
}

type searchService struct {
	searches []search
}

func (s searchService) Search(ctx context.Context, term string) (searchResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	mu := sync.Mutex{}
	var results []searchResult
	var fastest string

	for i := 0; i < len(s.searches); i++ {
		go func(ss search) {
			r, err := ss.Search(ctx, term)
			if err != nil {
				log.Print(err)
				return
			}

			mu.Lock()
			if results != nil {
				mu.Unlock()
				return
			}

			results = r
			fastest = ss.name
			mu.Unlock()

			cancel()
		}(s.searches[i])
	}

	<-ctx.Done()
	if results != nil {
		res := searchResponse{Results: results, Source: fastest}
		return res, nil
	}

	return searchResponse{}, ctx.Err()
}
