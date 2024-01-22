package go_json

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {

	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName: "rafli",
		LastName:  "muhammad",
	}

	encoder.Encode(customer)
}
