# Before dan After Test
Biasanya pada ketika melakukan testing, terdapat kondisi before dan after untuk setup dan finishing after test. Hal ini bisa dilakukan dengan menggunakan testing.M.
Contoh implementasinya:
```go
func TestMain(m *testing.M) {
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	fmt.Println("AFTER UNIT TEST")
}
```

Secara otomatis, semua unit test akan di run pada m.Run()
