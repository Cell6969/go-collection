package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	jonathan := Man{"jonathan"}
	jonathan.Married()
	fmt.Println(jonathan.Name)
}
