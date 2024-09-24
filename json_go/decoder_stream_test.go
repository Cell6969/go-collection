package json_go

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestDecoderStream(t *testing.T) {
	reader, _ := os.Open("Customer.json")
	var decoder *json.Decoder = json.NewDecoder(reader)

	var customer *Customer = &Customer{}
	decoder.Decode(customer)
	fmt.Println(customer)
}
