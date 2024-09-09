# Testify

Untuk mempermudah dalam proses testing, di golang terdapat package library yang bisa diinstal
```
github.com/stretchr/testify
```

Untuk install menggunakan command:
```shell
go get github.com/stretchr/testify
```

Contoh implementasi assertion menggunakan testify
```go
func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWord("jonathan")
	assert.Equal(t, "Hello jonathan", result)
}
```

## assert vs Require
Pada assert ketika suatu proses gagal maka untuk test selanjutnya akan dilanjut sedangkan require ketika 1 unit test gagal
maka unit test lain tidak akan dilanjut