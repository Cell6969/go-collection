# Request
Request di golang merupakan struct yang merepresentasikan HTTP Request. Beberapa componen dari HTTP Request seperti method, header, body, dll

Contoh penggunaan:
```go
var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.Method)
		fmt.Fprintln(w, r.RequestURI)
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

