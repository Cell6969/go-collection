package json_go

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArrayEncode(t *testing.T) {
	var customer *Customer = &Customer{
		FirstName: "Aldo",
		LastName:  "Jonathan",
		Hobbies:   []string{"Gaming", "Coding"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	var jsonString string = `{"FirstName":"Aldo","LastName":"Jonathan","Age":0,"Hobbies":["Gaming","Coding"]}`
	var jsonBytes []byte = []byte(jsonString)

	var customer *Customer = &Customer{}
	var err error = json.Unmarshal(jsonBytes, customer)

	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
}

func TestJSONArrayComplexEncode(t *testing.T) {
	var customer *Customer = &Customer{
		FirstName: "Aldo",
		LastName:  "Jonathan",
		Addressess: []Address{
			{
				Street:     "bekasi",
				Country:    "indonesia",
				PostalCode: "123",
			},
			{
				Street:     "bekasi1",
				Country:    "indonesia1",
				PostalCode: "1231",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	var jsonString string = `{"FirstName":"Aldo","LastName":"Jonathan","Age":0,"Hobbies":null,"Addressess":[{"Street":"bekasi","Country":"indonesia","PostalCode":"123"},{"Street":"bekasi1","Country":"indonesia1","PostalCode":"1231"}]}`
	var jsonBytes []byte = []byte(jsonString)

	var customer *Customer = &Customer{}
	var err error = json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Addressess[0].Street)
	fmt.Println(customer.Addressess[0].Country)
	fmt.Println(customer.Addressess[0].PostalCode)
}

func TestDecodeOnlyJSONArray(t *testing.T) {
	var jsonString string = `[{"Street": "Bekasi", "Country" : "Indonesia", "PostalCode": "123"},{"Street": "Bekasi", "Country" : "Indonesia", "PostalCode": "123"}]`
	var jsonBytes []byte = []byte(jsonString)

	var addresses *[]Address = &[]Address{}

	var err error = json.Unmarshal(jsonBytes, addresses)

	if err != nil {
	}

	fmt.Println(addresses)
}
