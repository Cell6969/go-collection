package main

import "fmt"

func main() {
	// Contoh pembuatan array
	var names = [...]string{"jonathan", "abdul", "daji"}
	fmt.Println(names)

	var values = [...]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 101}
	fmt.Println(values)

	// Len array
	fmt.Println(len(values))

	// mendapatkan element pada index tertentu
	fmt.Println(values[2])

	// replace element pada array
	values[3] = 1000
	fmt.Println(values)
}
