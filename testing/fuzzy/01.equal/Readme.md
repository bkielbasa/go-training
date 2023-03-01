# Equal

## Task
There's a `Equal` function that checks if two byte arrays are equal. Check its correctness using fuzzy testing.

<details><summary>Tip 1</summary>

```go
func FuzzEqual(f *testing.F) {
  f.Fuzz(func(t *testing.T, b1, b2 []byte) {
  })
}

```

</details>

<details><summary>Tip 2</summary>

You can shorten the fuzzy test execution time using `-fuzztime=5s` attribute.

</details>
