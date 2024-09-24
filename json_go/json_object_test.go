package json_go

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName string
	LastName  string
	Age       int
}

func TestJSONObject(t *testing.T) {
	var customer Customer = Customer{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONObjectFromMap(t *testing.T) {
	var person map[string]interface{} = map[string]interface{}{
		"first_name": "John",
		"last_name":  "Doe",
		"age":        30,
	}

	bytes, _ := json.Marshal(person)
	fmt.Println(string(bytes))
}
