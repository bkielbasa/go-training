# Slow reader

You have the following type implemented.


```go
type slowReader struct{}

func (sr slowReader) Read(p []byte) (n int, err error) {
	time.Sleep(10 * time.Millisecond)
	txt := "reading\n"
	copy(p, []byte(txt))
	n = len(txt)
	return n, nil
}
```

The `slowReader` implements `io.Reader` interface. The problem is that the `slowReader` is very slow. Your task is to implement a new `cancelableReader` that will accept any `io.Reader` and `context.Context`.
The goal is to implement the `cancelableReader` to react to context cancellation signal, return an error when it happens and stop reading from the `slowReader`.

The setup code may look similar to the following:

```go
func main() {
    sl := slowReader{}

    ctx, cancel := context.WithTimeout(context.Background(), time.Second/2)
    defer cancel()

    r := newCancelableReader(ctx, sl)

    read, err := io.ReadAll(r)
    // print the output and handle the error
}
