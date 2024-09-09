package main

import "fmt"

func getHello() (string, string) {
	return "Hello", "hello"
}

func getFullName() (string, string) {
	return "joko", "anwar"
}

func getCompleteName() (firstName string, lastName string) {
	firstName = "jonathan"
	lastName = "joestar"
	return firstName, lastName
}

func main() {
	first, second := getHello()
	fmt.Println(first, second)

	firstname, _ := getFullName()
	fmt.Println(firstname)

	a, b := getCompleteName()
	fmt.Println(a, b)
}
