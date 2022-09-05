# How to debug GC?

Run your app with `GODEBUG=gctrace=1` env var.

```sh
GODEBUG=gctrace=1 go run .
```

The tracer will print such output

```
gc 1405 @6.068s 11%: 0.058+1.2+0.083 ms clock, 0.70+2.5/1.5/0+0.99 ms cpu, 7->11->6 MB, 10 MB goal, 12 P
```

## What does the output mean?

```
gc 1405 @6.068s 11%: 0.058+1.2+0.083 ms clock, 0.70+2.5/1.5/0+0.99 ms cpu, 7->11->6 MB, 10 MB goal, 12 P

// General
gc 1404     : The 1404 GC run since the program started
@6.068s     : Six seconds since the program started
11%         : Eleven percent of the available CPU so far has been spent in GC

// Wall-Clock
0.058ms     : STW        : Mark Start       - Write Barrier on
1.2ms       : Concurrent : Marking
0.083ms     : STW        : Mark Termination - Write Barrier off and clean up

// CPU Time
0.70ms      : STW        : Mark Start
2.5ms       : Concurrent : Mark - Assist Time (GC performed in line with allocation)
1.5ms       : Concurrent : Mark - Background GC time
0ms         : Concurrent : Mark - Idle GC time
0.99ms      : STW        : Mark Term

// Memory
7MB         : Heap memory in-use before the Marking started
11MB        : Heap memory in-use after the Marking finished
6MB         : Heap memory marked as live after the Marking finished
10MB        : Collection goal for heap memory in-use after Marking finished

// Threads
12P         : Number of logical processors or threads used to run Goroutines
```
