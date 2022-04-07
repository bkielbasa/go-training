## Generate source code

```sh
protoc --go_out=./shared/ --go_opt=paths=source_relative --go-grpc_out=./shared/ --go-grpc_opt=paths=source_relative ./ping.proto
```
