# Allocations part 2

Build the code in the current directory.

```sh
go build -gcflags '-m -l' create.go
```

Can you see the allocation?

Then run benchmarks:

```sh
go test -bench . -benchmem
```

## When an allocation may occure?

 * declare variables
 * call the built-in new function.
 * call the built-in make function.
 * modify slices and maps with composite literals[^1].
 * convert integers to strings.
 * concatenate strings by using use +.
 * convert between strings to byte slices, and vice versa.
 * convert strings to rune slices.
 * box values into interfaces (converting non-interface values into interfaces).
 * append elements to a slice and the capacity of the slice is not large enough.
 * put new entries into maps and the underlying array (to store entries) of the map is not large enough to store the new entries.

[^1]: Composite literals construct values for structs, arrays, slices, and maps... They consist of the type of the literal followed by a brace-bound list of elements.