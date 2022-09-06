package highest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func FuzzTestHTTPHandler(f *testing.F) {
	// Create a new server hosting our calculate func
	srv := httptest.NewServer(http.HandlerFunc(CalculateHighest))
	defer srv.Close()

	// Add Corpus Seeds
	testCases := []ValuesRequest{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}},
		{[]int{-50, -9, -8, -7, -6, -5, -4, -3, -2, -1}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}},
		{[]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120, 130, 140, 150, 160, 170, 180, 190, 200}},
	}

	for _, testCase := range testCases {
		// Skip error, very bad practice
		data, _ := json.Marshal(testCase)
		// Add JSON data as Corpus
		f.Add(data)
	}
	// Start fuzzing
	f.Fuzz(func(t *testing.T, data []byte) {
		// Create a new request
		// using DefaultClient, should not be done in prod
		resp, err := http.DefaultClient.Post(srv.URL, "application/json", bytes.NewBuffer(data))
		if err != nil {
			t.Errorf("Error: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
		}
		// Umarshal data
		var response int
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	})
}
