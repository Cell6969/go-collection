# Test Benchmark

Pada golang juga dilakukan test benchmark yang bertujuan untuk mengukur performa aplikasi. Test benchmar (Testing.B) memiliki syarat yaitu tidak boleh ada return.

Contoh implementasinya
```go
func BenchmarkHelloWord(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWord("jonathan")
	}
}
```

Untuk menjalankan perintah benchmark test bisa dengan command berikut:
```shell
go test -v -bench=.
```

untuk merunning hanya test benchmark saja bisa dengan command berikut:
```shell
go test -v -run=TestTidakAda -bench=enchmarkHelloWord
```

## Table benchmark
contoh implementasi
```go

```