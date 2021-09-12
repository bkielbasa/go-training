package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()

	go func1(ctx)

	<-ctx.Done()
}

func func1(ctx context.Context) {

	go func2(ctx)

	<-ctx.Done()
	fmt.Println("func1 finished")
}

func func2(ctx context.Context) {

	go func3(ctx)

	<-ctx.Done()
	fmt.Println("func2 finished")
}

func func3(ctx context.Context) {
	<-ctx.Done()
	fmt.Println("func3 finished")
}
