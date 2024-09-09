# Array

contoh immplementasi array
```go
package main

import "fmt"

func main() {
	// Contoh pembuatan array
	var names [3]string

	names[0] = "alfred"
	names[1] = "bob"
	names[2] = "david"

	fmt.Println(names)
	// Error array out bound
	//names[3] = "kei"

	// Deklarasi array secara langsung
	var values = [3]int{10, 20, 30}
	fmt.Println(values)
}
```

## Function array
