# Task

Add tracing to the chat application. Use opentelemetry.

## Preparation

### Install jaeger

To run Jaeger in a docker instance run this command:

```
make jaeger
```

or

```
docker run -d --name jaeger -e COLLECTOR_ZIPKIN_HOST_PORT=:9411 -p 5775:5775/udp -p 6831:6831/udp -p 6832:6832/udp -p 5778:5778 -p 16686:16686 -p 14268:14268 -p 14250:14250 -p 9411:9411 jaegertracing/all-in-one:latest
```

It will run Jaeger in Docker in background.


<details>
<summary>TIP for server</summary>

Use this code while initializing the server

```
	grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
	grpc.StreamInterceptor(otelgrpc.StreamServerInterceptor()),
```
</details>

<details>
<summary>TIP for client</summary>

Use this code while initializing the client

```
grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
```
</details>
