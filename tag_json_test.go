package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func TestJSONTag(t *testing.T) {

	product := Product{
		Id:       "0101",
		Name:     "rafli",
		ImageURL: "anda.cakjnsda.com/kamsdad",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJSONTagDecode(t *testing.T) {

	jsonString := `{"id":"0101","NAME":"rafli","image_url":"anda.cakjnsda.com/kamsdad"}`
	jsonBytes := []byte(jsonString)

	product := &Product{}

	json.Unmarshal(jsonBytes, product)
	fmt.Println(product)
}
