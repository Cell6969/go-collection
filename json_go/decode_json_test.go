package json_go

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodedJSON(t *testing.T) {
	var jsonString string = `{"FirstName":"John","LastName":"Doe","Age":30}`
	var jsonBytes []byte = []byte(jsonString)

	var customer *Customer = &Customer{}
	var err error = json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Age)
}
