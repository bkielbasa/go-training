# Highest

## Task

Write a fuzzy test for the HTTP handler `CalculateHighest`.

### Adding to corpus

To add values to the corpus us the `f.Add()` method. You can start with the following example:

<details><summary>Tip 1</summary>
```go
// Create a new server hosting our calculate func
	srv := httptest.NewServer(http.HandlerFunc(CalculateHighest))
	defer srv.Close()

	// Create example values for the fuzzer
	testCases := []ValuesRequest{
		ValuesRequest{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		ValuesRequest{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}},
		ValuesRequest{[]int{-50, -9, -8, -7, -6, -5, -4, -3, -2, -1}},
		ValuesRequest{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}},
		ValuesRequest{[]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120, 130, 140, 150, 160, 170, 180, 190, 200}},
	}

	// Add Corpus Seeds
	for _, testCase := range testCases {
		// Skip error, very bad practice
		data, _ := json.Marshal(testCase)
		// Add JSON data as Corpus
		f.Add(data)

	}
```
</details>
