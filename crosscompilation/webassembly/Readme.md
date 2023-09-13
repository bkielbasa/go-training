# Simple console log

To build the app use the following command line

```sh
GOOS=js GOARCH=wasm go build -o main.wasm demo.go
```

After that run a simple HTTP server

```sh
python3 -m http.server
```

and go to http://localhost:8000.

Can you see the log in the browser?