# Pointer
Singkatnya adalah pointer adalah referensi objek/struct
Contoh :
```go
package main

import "fmt"

type Address struct {
	City     string
	Province string
	Country  string
}

func main() {
	address1 := Address{
		City:     "bekasi",
		Country:  "indonesia",
		Province: "jawa barat",
	}

	address2 := address1
	address2.City = "jakarta"
	fmt.Println("Address1", address1)
	fmt.Println("Address2", address2)
}
```
Pada code diatas, jika dieksekusi maka city pada address1 tidak akan berubah. Hal ini dikarenakan address2 adalah pass by value oleh address1.

Dengan pointer:
```go
var address1 Address = Address{
		City:     "bekasi",
		Province: "Jawa Barat",
		Country:  "Indonesia",
}
var address2 *Address = &address1
address2.City = "jakarta"
fmt.Println(address1)
fmt.Println(address2)
```
Pada code diatas, baik address1 atau address2 akan mengalami perubahan pada city.

## Asterisk
Pada code dibawah:
```go
package main

import "fmt"

type Address struct {
	City     string
	Province string
	Country  string
}

func main() {
	var adress1 Address = Address{
		Country:  "Indonesia",
		City:     "Jakarta",
		Province: "Jakarta",
	}
	var adress2 *Address = &adress1

	adress2.City = "bandung"
	fmt.Println(adress1)
	fmt.Println(adress2)

	adress2 = &Address{
		City:     "Papua",
		Province: "Jakarta",
		Country:  "Indonesia",
	}
	fmt.Println(adress1)
	fmt.Println(adress2)
}
```
jika dilihat, maka address1 tidak akan sama dengan address2. Karena address2 memulai melakukan reference baru. Untuk membuat address1 juga berubah karena perubahan address2 maka bisa dilakuakn dengan menambah asterisk:
```go
*adress2 = Address{
		City:     "Papua",
		Province: "Jakarta",
		Country:  "Indonesia",
	}
	fmt.Println(adress1)
	fmt.Println(adress2)
```
Dengan demikian siapapun yang memiliki pointer dari address2 akan berubah.