# Type Assertion
Contoh implementasinya:
```go
package main

import "fmt"

func random() any {
	return "Ok"
}

func main() {
	var result any = random()
	var resultString string = result.(string)
	fmt.Println(resultString)

	// Error karena return nya adalah string
	//var resultInt int = result.(int)
	//fmt.Println(resultInt)

	// best handling
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown", value)
	}
}
```