# MemCPU profiling

## Run benchmakrs

```sh
go test -run none -bench . -benchtime 3s -benchmem -memprofile p.out
```


```sh
go tool pprof p.out
```

## Run pprof

```sh
(pprof) list algOne
(pprof) web list algOne
```

## Run tests with escape analysis

```sh
go test -run none -bench . -benchtime 3s -benchmem -memprofile p.out -gcflags -m=2
```

# Run pprof again

```sh
go tool pprof p.out
```

If there's a pointer in stack, it's an automatic allocation.


Take a look at inlining..

Passing concrete argument as interface causes allocation.

## Apply changes to the code

```sh
go tool pprof p.out
```

## Run tests with CPU profile

```sh
go test -run none -bench . -benchtime 3s -benchmem -cpuprofile p.out
```
