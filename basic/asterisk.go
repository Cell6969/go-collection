package main

import "fmt"

type Address struct {
	City     string
	Province string
	Country  string
}

func main() {
	var adress1 Address = Address{
		Country:  "Indonesia",
		City:     "Jakarta",
		Province: "Jakarta",
	}
	var adress2 *Address = &adress1

	adress2.City = "bandung"
	fmt.Println(adress1)
	fmt.Println(adress2)

	//adress2 = &Address{
	//	City:     "Papua",
	//	Province: "Jakarta",
	//	Country:  "Indonesia",
	//}
	*adress2 = Address{
		City:     "Papua",
		Province: "Jakarta",
		Country:  "Indonesia",
	}
	fmt.Println(adress1)
	fmt.Println(adress2)
}
