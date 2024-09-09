package main

import "fmt"

type Address struct {
	City     string
	Country  string
	Province string
}

func ChanceCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	var address *Address = &Address{}
	ChanceCountryToIndonesia(address)
	fmt.Println(address)

	// jika variable awal yang dibuat bukan pointer
	var address2 Address = Address{}
	ChanceCountryToIndonesia(&address2)
	fmt.Println(address2)
}
