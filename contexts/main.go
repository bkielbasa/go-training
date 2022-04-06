package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*150)
	defer cancel()

	go run(ctx, time.Millisecond*100)
	go run(ctx, time.Millisecond*200)

	<-ctx.Done()

	time.Sleep(100 * time.Millisecond)
}

func run(ctx context.Context, duration time.Duration) {
	select {
	case <-time.After(duration):
		fmt.Println("done")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
