package main

import (
	"bufio"
	pb "chat/shared"
	"chat/traces"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc"
)

var addr = flag.String("addr", "localhost:50051", "the address to connect to")
var user = flag.String("user", "Johny Bravo", "name used to connect to the chat")

func main() {
	flag.Parse()
	if *user == "" {
		log.Fatal("please provide the user")
	}

	tp, err := traces.Provider(*user, "http://localhost:14268/api/traces")
	if err != nil {
		log.Fatal(err)
	}

	// Register our TracerProvider as the global so any imported
	// instrumentation in the future will default to using it.
	otel.SetTracerProvider(tp)

	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithInsecure(),

		// both two are important here
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
	)
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
			ctx, span := otel.Tracer("client").Start(ctx, "SendMessage")
			_, err := c.SendMessage(ctx, &pb.SendMessageRequest{
				User:    *user,
				Message: scanner.Text(),
			})
			span.End()

			if err != nil {
				log.Printf("send msg err: %s", err)
				return
			}
		}
	}()

	<-ctx.Done()
}
