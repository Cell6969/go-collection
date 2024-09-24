package json_go

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)

	var customer *Customer = &Customer{
		FirstName: "aldo",
		LastName:  "jek",
		Addressess: []Address{
			{
				Street:     "bekasi",
				Country:    "indonesia",
				PostalCode: "1111",
			},
		},
	}

	err := encoder.Encode(customer)
	if err != nil {
	}

	fmt.Println(customer)
}
