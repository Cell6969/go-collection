# Golang Web

Golang menyediakan library untuk menjalankan web server yaitu net/http. Dengan net/http golang bisa menjalankan web server tanpa eksternal web server. Namun perlu diingat bahwa untuk rekomendasi nya tetap menggunakan framework dikarenakan sudah dipermudah,

Untuk menjalankan server sederhana bisa seperti berikut:

```go
server := http.Server{
		Addr: "localhost:8080",
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
```

