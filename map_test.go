package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapDecode(t *testing.T) {

	jsonString := `{"id":"T001", "name":"rafli"}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}

	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
}

func TestMapEncode(t *testing.T) {

	product := map[string]interface{}{
		"id":   "P001",
		"name": "muhammad",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
