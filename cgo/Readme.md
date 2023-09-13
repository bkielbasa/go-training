# Cgo

To call a C function from Go code you have to use the `C` package like this.

```go
C.get_string()
```

# Task 1

Call a C function with an argument. To create a new C string from a Go string you have to use `C.CString(str)` function.

```go
str := "lorem ipsum"
cStr := C.CString(str)
```

There's already implemented C `void print_string(char* a)` function. Your task is to call it from Go code