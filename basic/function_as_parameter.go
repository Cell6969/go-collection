package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

// function yang dijadikan parameter harus sesuai kontrak dengan parameter yang ditentukan
func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	// Implementasi cara ke 1
	sayHelloWithFilter("dona", spamFilter)

	// Implementasi cara ke 2
	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)
}
