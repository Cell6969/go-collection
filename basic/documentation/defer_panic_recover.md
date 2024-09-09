# Defer
defer function adalah function yang bisa dijadwalkan untuk dieksekusi setelah sebuah function selesai di eksekusi.
Contoh implementasinya:
```go
package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	defer logging()
	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApplication()
}
```

# Panic
Panic function adalah function yang bisa digunakan untuk menghentikan program. Saat panic terjalan, defer tetap akan berjalan.
Contoh implementasinya
```go
package main

import "fmt"

func endApp() {
	fmt.Println("End App")
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("Erorr")
	}
}

func main() {
	runApp(false)
}
```
# Recover
Recover adalah function yang bisa digunakan untuk menangkap data panic.
```go

```