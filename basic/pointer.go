package main

import "fmt"

type Address struct {
	City     string
	Province string
	Country  string
}

func main() {
	//address1 := Address{
	//	City:     "bekasi",
	//	Country:  "indonesia",
	//	Province: "jawa barat",
	//}
	//
	//address2 := address1
	//address2.City = "jakarta"
	//fmt.Println("Address1", address1)
	//fmt.Println("Address2", address2)

	// dengan pointer
	var address1 Address = Address{
		City:     "bekasi",
		Province: "Jawa Barat",
		Country:  "Indonesia",
	}
	var address2 *Address = &address1
	address2.City = "jakarta"
	fmt.Println(address1)
	fmt.Println(address2)
}
