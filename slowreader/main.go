package main

import (
	"context"
	"fmt"
	"io"
	"time"
)

type slowReader struct{}

func (sr slowReader) Read(p []byte) (n int, err error) {
	time.Sleep(10 * time.Millisecond)
	txt := "reading\n"
	copy(p, []byte(txt))
	n = len(txt)
	return n, nil
}

func main() {
	sr := slowReader{}

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	defer cancel()
	r := newReader(ctx, sr)

	msg, err := io.ReadAll(r)
	fmt.Printf("%serr: %s", msg, err)
}

type reader struct {
	ctx context.Context
	r   io.Reader
}

func (r *reader) Read(p []byte) (n int, err error) {
	select {
	case <-r.ctx.Done():
		return 0, r.ctx.Err()
	default:
		return r.r.Read(p)
	}
}

func newReader(ctx context.Context, r io.Reader) io.Reader {
	return &reader{ctx: ctx, r: r}
}
