# GC demo

The project contains an app that exposes a simple API on port `:3333`. It accepts an URL param `size` that should be a positive integer. It's a size of a slice of bytes that's created in the request.

Run the app with sizes like:

- 25000
- 50000
- 100000
- 1000000

and take a look how many times the GC is called and with what values.

To run the app with GC tracing enabled, use the following command.

```sh
GODEBUG=gctrace=1 go run .
```
