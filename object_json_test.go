package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

type Customer struct {
	FirstName string
	LastName  string
	Age       int8
	Hobbies   []string
	Addresses []Address
}

func TestJSONObject(t *testing.T) {

	customer := Customer{
		FirstName: "muhammad",
		LastName:  "rafli",
		Age:       17,
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}
