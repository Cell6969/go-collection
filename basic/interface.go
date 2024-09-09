package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

type Person struct {
	Name string
}

func (p Person) GetName() string {
	return p.Name
}

type Animal struct {
	Name string
}

func (a Animal) GetName() string {
	return a.Name
}

func main() {
	person := Person{Name: "aldo"}
	SayHello(person)

	gajah := Animal{
		Name: "Gajah",
	}

	SayHello(gajah)
}
