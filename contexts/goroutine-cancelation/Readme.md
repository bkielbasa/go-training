# Goroutine cancelation

## Task

Write a program that runs two goroutines. Those goroutines should execute for 100 and 200 ms. You can use `time.Sleep()` to simulate the delay.

Use `context` package to set a timeout to `150` so only the fastest goroutine will finish.
