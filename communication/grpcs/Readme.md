
## Requrements

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/square/certstrap@latest

sudo apt install protobuf-compiler
```

Don't forget to add go bin path to your $PATH

```sh
# put it in your .bashrc
export PATH=$PATH:/path/to/go/bin
```
