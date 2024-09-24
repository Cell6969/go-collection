package json_go

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageURl string `json:"image_url"`
}

func TestJSONTag(t *testing.T) {
	var product *Product = &Product{
		Id:       "P0001",
		Name:     "Handphone",
		ImageURl: "http://image.com",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJSONTagDecode(t *testing.T) {
	var jsonString string = `{"id":"P0001","name":"Handphone","image_url":"http://image.com"}`
	var jsonBytes []byte = []byte(jsonString)

	var product *Product = &Product{}

	json.Unmarshal(jsonBytes, product)

	fmt.Println(product)
}
