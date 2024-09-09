# Package

Di golang, ketika membuat suatu folder maka folder tersebut menjadi parent package. Package adalah tempat yang bisa digunakan untuk mengorganisir koder program yang dibuat.

Contoh, buat folder **helper** dan buat file helper.go
```go
package helper

func SayHello(name string) string {
	return "Hello " + name
}
```

Kemudian untuk mengimport:
```go
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
}
```
## Access Modifier
By default ketika membuat function atau variable pada package, huruf kapital bisa diakses dari luar, sedangkan huruf kecil tidak bisa.

## Package Initialization
Contoh implementasi:

mysql.go
```go
package database

var connection string

func init() {
	connection = "MySql"
}

func GetDatabase() string {
	return connection
}
```

untuk implementasi:
```go
package main

import (
	"basic/database"
	"fmt"
)

func main() {
	fmt.Println(database.GetDatabase())
}
```

Pada code diatas, init berfungsi menjalan function pertama kali namun bersifat private. Pada kondisi tersebut, init akan jalan ketika package dipanggil. Untuk menjalankan init tanpa harus memanggil function pada package tersebut
```go
import (
	"basic/database"
	_ "basic/internal"
	"fmt"
)
```