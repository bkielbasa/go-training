package main

import (
	"bufio"
	pb "chat/shared"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
)

var addr = flag.String("addr", "localhost:50051", "the address to connect to")
var user = flag.String("user", "Johny Bravo", "name used to connect to the chat")

func main() {
	flag.Parse()
	if *user == "" {
		log.Fatal("please provide the user")
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewChatServiceClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	r, err := c.JoinChat(ctx, &pb.JoinChatRequest{User: *user})
	if err != nil {
		log.Fatalf("could not join: %s", err)
	}

	go func() {
		for {
			msg, err := r.Recv()
			if err != nil {
				log.Printf("recv err: %s", err)
				cancel()
				return
			}

			fmt.Printf("%s: %s\n", msg.User, msg.Message)
		}
	}()

	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go func() {
		<-gracefulStop
		cancel()
	}()

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			_, err := c.SendMessage(context.Background(), &pb.SendMessageRequest{
				User:    *user,
				Message: scanner.Text(),
			})

			if err != nil {
				log.Printf("send msg err: %s", err)
				return
			}
		}
	}()

	<-ctx.Done()
}
