
bench:
	@go test -bench . -benchmem -memprofile p.out

pprof:
	@go tool pprof p.out

bench-escape:
	@go test -bench . -benchmem -memprofile p.out -gcflags -m=2
bench-cpu:
	@go test -bench . -benchmem -cpuprofile p.out -gcflags -m=2