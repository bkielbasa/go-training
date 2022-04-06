package main

import (
	"context"
	"log"
	pb "ping/shared"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewPingServiceClient(conn)

	index := 0
	for {
		trip_time := time.Now()
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.Ping(ctx, &pb.PingRequest{Data: "fsdfsdfd"})

		if err != nil {
			log.Fatalf("could not connect to: %v", err)
		}

		log.Printf("%d characters roundtrip to (%s): seq=%d time=%s", len(r.Data), "localhost:8000", index, time.Since(trip_time))
		time.Sleep(1 * time.Second)
		index++
	}
}
