# Interface
Interface adalah tipe data abstract dan tidak dapat diimplementasikan secara langsung. Biasany berisi method method dan kontrak
Contoh implementasi:
```go
package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

func main() {

}
```
Pada code diatas, tidak dapat langsung dieksekusi karena interface sendiri merupakan kontrak bukan data. Oleh karena itu perlu membuat struct serta mengimplementasikan method pada struct sesuai dengan interface
```go
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

func main() {
	person := Person{Name: "aldo"}
	SayHello(person)
}
```
Karena data struct sudah ada dan sudah mengimplementasikan method GetName sesuai dengan kontrak maka bisa dijalankan fungsi SayHello sesuai dengan inteface.
Contoh lain:
```go
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
```
## Interface Kosong
Seperti tipe data any pada typescript, tidak memiliki kontrak tertentu.