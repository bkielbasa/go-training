package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func getResp(url string) (string, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", err
	}

	c := http.Client{}
	respBody := make(chan []byte)

	for i := 0; i < 3; i++ {
		go func() {
			resp, err := c.Do(req)
			if err != nil {
				fmt.Println(err)
				return
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err)
				return
			}

			cancel()
			respBody <- body
		}()
	}

	// 3 goroutines
	body := <-respBody
	return string(body), nil
}

func main() {
	fmt.Println(getResp("https://google.org"))

	time.Sleep(time.Second / 2)
}
