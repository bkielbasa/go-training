# Task

Implement a simple chat with client and server.

The protobuf that defines the communication between themm

```protobuf
/* language version */
syntax = "proto3";

/* message definitions */
message JoinChatRequest {
  string user = 1;
}

message SendMessageRequest {
  string message = 1;
  string user = 2;
}

message MessageResponse {
  string message = 1;
  string user = 2;
  int32 timestamp = 3;
}

message EmptyResponse {}

/* service with rpc method signatures */
service ChatService {
  rpc joinChat(JoinChatRequest) returns (stream MessageResponse) {}
  rpc sendMessage(SendMessageRequest) returns (EmptyResponse) {}
}
```

## Installation

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
```