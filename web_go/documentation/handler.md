# Handler

Untuk membuat handler yang mengatur request dan response bisa menggunakan HandleFunc.

Contoh implementasi:
```go
var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		//logic web
		fmt.Fprint(w, "Hello World")
	}

	server := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
```

## Multiple Handler
Problem menggunakan HandlerFunc adalah dia hanya bisa membuat 1 endpoint saja. Sedangkan kasus real akan banyak endpoint yanh akan dibuat. Untuk itu maka perlu menggunakan ServeMux. ServeMux sendiri adalah implementasi Handler yang mendukung multiple endpoint

Contoh penggunaan:
```go
mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})
	mux.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hi")
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
```

## URL Pattern
Pada URL pattern jika kita tambahkan "/" pada akhir url maka semua url tersebut akan dieksekusi namun perlu diperhatikan eksekusi pertama dimulai dari url yang lebih panjang.
