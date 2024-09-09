# Function

## Function tanpa parameter
```go
package main

import "fmt"

func sayHello() {
	fmt.Println("Hello World")
}

func main() {
	sayHello()
}
```
## Function dengan parameter
```go
package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	sayHelloTo("bob", "johnson")
}
```
## Function dengan Return
```go
package main

import "fmt"

func getHello(name string) string {
	return "Hello " + name
}

func main() {
	result := getHello("jonathan")
	fmt.Println(result)
}
```

## Function multiple value
```go
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
```

## Variadic Function
```go
package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	fmt.Println(sumAll(1, 2, 3, 4, 6))
	fmt.Println(sumAll(1, 2, 3))

	// dengan slice
	numbers := []int{10, 10, 10, 10}
	fmt.Println(sumAll(numbers...))
}
```

## Function as Value
```go
package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	contoh := getGoodBye
	misal := getGoodBye

	fmt.Println(contoh("azhar"))
	fmt.Println(misal("azhar"))
}
```
## Function as Parameter
```go
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
```

## 