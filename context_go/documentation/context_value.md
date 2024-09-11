# Context WIth Value

Kita bisa menambahkan value context dengan aturan pair (key-value). Saat penambahan value pada context maka akan terbuat context baru sedangkan context original nya tidak berubah.

Contoh implementasinya:
```go
func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)
}
```

Hasil outputnya:
```shell
context.Background
context.Background.WithValue(type string, val B)
context.Background.WithValue(type string, val C)
context.Background.WithValue(type string, val B).WithValue(type string, val D)
context.Background.WithValue(type string, val B).WithValue(type string, val E)
context.Background.WithValue(type string, val C).WithValue(type string, val F)
```
Dari hasil terlihat bahwa context A sebagai parent tidak memiliki value , kemudia context B memiliki value b dan context C memiliki value c. Context D dan E menurunkan dari context B sehingga membawa value B dan masing masing D dan E. Namun context B tidak memiliki value D dan E. Singkatnya semakin child context dia mendapatkan value dari parent -parentnya sedangkan parent tidak memiliki value dari child.
