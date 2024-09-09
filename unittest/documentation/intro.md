# Unit Test

Pada golang terdapat 3 test yaitu testing.T, testing.M dan testing.B.

Contoh unit test:

hello_world.go
```go
package helper

func HelloWord(name string) string {
	return "Hello " + name
}
```

untuk unit test nya:
```go
package helper

import "testing"

func TestHelloWord(t *testing.T) {
	result := HelloWord("jonathan")
	if result != "Hello jonathan" {
		panic("Result it not match")
	}
}
```

Untuk menjalankan unit test bisa dengan command berikut:
```shell
go test

go test -v // untuk melihat lebih detail test

go test -v -run <TestNamaFunction>
```
