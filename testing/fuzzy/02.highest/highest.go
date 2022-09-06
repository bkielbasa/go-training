package highest

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ValuesRequest struct {
	Values []int `json:"values"`
}

func CalculateHighest(w http.ResponseWriter, r *http.Request) {
	// Declare a valuerequest
	var vr ValuesRequest

	// Decode and respond with error incase it fails
	err := json.NewDecoder(r.Body).Decode(&vr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var high int
	// Range all values
	for _, value := range vr.Values {
		// Check if value is higher than high
		if value > high {
			// If so, set high to value
			high = value
		}
	}

	// Return high
	if high == 50 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong"))
	}
	fmt.Fprintf(w, "%d", high)
}
