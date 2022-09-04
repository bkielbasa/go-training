# Goroutine dead

In the `main.go` file you can find a program that runs 3 goroutines. Those goroutines use the `context` package to find out about context cancellation. When the context is canceled, we can see no output. Why?
