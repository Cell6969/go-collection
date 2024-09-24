package json_go

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapDecode(t *testing.T) {
	var jsonString string = `{"id":"P0001","name":"Handphone","image_url":"http://image.com"}`
	var jsonBytes []byte = []byte(jsonString)

	var result map[string]interface{}

	json.Unmarshal(jsonBytes, &result)
	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["image_url"])
}
