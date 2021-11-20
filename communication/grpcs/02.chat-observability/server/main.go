package main

import (
	pb "chat/shared"
	"chat/traces"
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedChatServiceServer
	conns map[string]chan *pb.MessageResponse
}

func (s *server) JoinChat(in *pb.JoinChatRequest, srv pb.ChatService_JoinChatServer) error {
	log.Printf("User %v joined", in.User)
	conn := make(chan *pb.MessageResponse)

	s.conns[in.User] = conn
	defer func() {
		close(conn)
		delete(s.conns, in.User)
		log.Printf("disconnecting %s", in.User)
	}()

	for {
		select {
		case <-srv.Context().Done():
			return srv.Context().Err()
		case response := <-conn:
			if status, ok := status.FromError(srv.Send(response)); ok {
				switch status.Code() {
				case codes.OK:
					//noop
				case codes.Unavailable, codes.Canceled, codes.DeadlineExceeded:
					return errors.New("unavailable")
				default:
					return errors.New("something else...")
				}
			} else {
				log.Print(response)
			}
		}
	}

	return nil
}

func (s *server) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (*pb.EmptyResponse, error) {
	log.Printf("User %s said: %s", req.User, req.Message)

	for user, conn := range s.conns {
		if user == req.User {
			continue
		}

		conn <- &pb.MessageResponse{
			User:      req.User,
			Message:   req.Message,
			Timestamp: int32(time.Now().Unix()),
		}
	}

	return &pb.EmptyResponse{}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	tp, err := traces.Provider("server", "http://localhost:14268/api/traces")
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Cleanly shutdown and flush telemetry when the application exits.
	defer func(ctx context.Context) {
		// Do not make the application hang when it is shutdown.
		ctx, cancel = context.WithTimeout(ctx, time.Second*5)
		defer cancel()
		if err := tp.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}
	}(ctx)

	// Register our TracerProvider as the global so any imported
	// instrumentation in the future will default to using it.
	otel.SetTracerProvider(tp)

	s := grpc.NewServer(
		grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
		grpc.StreamInterceptor(otelgrpc.StreamServerInterceptor()),
	)
	pb.RegisterChatServiceServer(s, &server{
		conns: map[string]chan *pb.MessageResponse{},
	})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
