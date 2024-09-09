package main

import (
	"basic/helper"
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	// import from helper
	result := helper.SayHello("doni")
	fmt.Println(result)

	// import from accessMod
	fmt.Println(helper.Application)
}
