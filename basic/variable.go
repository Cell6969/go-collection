package main

import "fmt"

func main() {
	// Cara Pertama
	var name string

	name = "jonathan"
	fmt.Println(name)

	name = "aldo"
	fmt.Println(name)

	// Cara Kedua
	city := "jakarta"
	fmt.Println(city)

	// Deklarasi banyak variable
	var (
		firstname = "kaneda"
		lastname  = "1"
	)

	fmt.Println(firstname)
	fmt.Println(lastname)
}
