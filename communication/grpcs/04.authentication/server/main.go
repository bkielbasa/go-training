package main

import (
	pb "chat/shared"
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
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
			log.Print(srv.Context().Err())
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

	certFile := "./out/localhost.crt"
	keyFile := "./out/localhost.key"
	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		log.Fatalf("failed to credentials: %v", err)
	}

	s := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterChatServiceServer(s, &server{
		conns: map[string]chan *pb.MessageResponse{},
	})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
