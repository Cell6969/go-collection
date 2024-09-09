# Struct
Struct adalah template data yang digunakan untuk menggabungkan nol atau lebih tipe data (mirip object)
Contoh implementasinya:
```go
package main

import "fmt"

type Customer struct {
	Name    string
	Address string
	Age     int
}

func main() {
	// Membuat data untuk struct
	var jonathan Customer
	jonathan.Name = "Alphonse"
	jonathan.Address = "Kemang"
	jonathan.Age = 18
	fmt.Println(jonathan)

	// cara lain
	joko := Customer{
		Name:    "dona",
		Age:     10,
		Address: "Bekasi",
	}
	fmt.Println(joko)
}
```

## Struct Method
Karena struct itu adalah tipe data maka bisa juga digunakan sebagai parameter function. Sebagai contoh:
```go
package main

import "fmt"

type Customer struct {
	Name    string
	Address string
	Age     int
}

func (customer Customer) sayHello(name string) bool {
	fmt.Println("Hello", name, "my name is", customer.Name)
	return true
}

func main() {
	// Membuat data untuk struct
	var jonathan Customer
	jonathan.Name = "Alphonse"
	jonathan.Address = "Kemang"
	jonathan.Age = 18
	fmt.Println(jonathan)

	// cara lain
	joko := Customer{
		Name:    "dona",
		Age:     10,
		Address: "Bekasi",
	}
	fmt.Println(joko)
	joko.sayHello("albert")
}
```
