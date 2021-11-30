# Task

Add SSL/TLS to the communication between client and server.

The server can be configured by changing this code:

```go
creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
if err != nil {
    // handle the error - no ignore it!
}
s := grpc.NewServer(grpc.Creds(creds))
```

The client should be updated as follows:

```go
creds, err := credentials.NewClientTLSFromFile(certFile, "")
if err != nil {
    // handle the error - no ignore it!
}
conn, _ := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
```

To generate the certificate, you can use openssl or [certstrap](https://github.com/square/certstrap).

```sh
certstrap init --common-name "developer20.com"
certstrap request-cert -domain localhost
certstrap sign localhost --CA developer20.com
```

