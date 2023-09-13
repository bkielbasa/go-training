# Generic list

Refactor the struct 

```go
type List struct {
	val  int
	Next *List
}
```

to use generics so we can create a list with any type