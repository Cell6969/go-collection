# FileServer

Golang memiliki fitur yang bernama FilServer. Dengan ini golang dapat membuat handler sebagai statci file server.

Contoh:
```go
func TestFileServer(t *testing.T) {
	var directory http.Dir = http.Dir("./resources")
	var fileServer http.Handler = http.FileServer(directory)

	var mux *http.ServeMux = http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer)) // prefix untuk url access

	var server http.Server = http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	var err error
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
```

Dengan demikian ketika mengakses /static bisa didapatkan file file yang bersifat static.

### go embed
Golang terbaru mendukung fitur embed yang memungkinkan untuk mengcompile file static. Contoh:
```go
//go:embed resources
var resources embed.FS

func TestFileServerGolangEmbed(t *testing.T) {
	var directory fs.FS
	var _ error
	directory, _ = fs.Sub(resources, "resources")
	var fileServer http.Handler = http.FileServer(http.FS(directory))

	var mux *http.ServeMux = http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer)) // prefix untuk url access

	var server http.Server = http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	var err error
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
```

## ServeFile
Ada kondisi dimana ingin menggunakan static file tetapi spesifik dan sesuai yang diinginkan. Hal ini bisa dilakukan menggunakan function http.ServeFile().

Contoh implementasi:
```go
func ServeFile(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("name") == "" {
		http.ServeFile(w, r, "./resources/notfound.html")
	} else {
		http.ServeFile(w, r, "./resources/ok.html")
	}
}

func TestServeFileServer(t *testing.T) {
	var server http.Server = http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	var err error = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
```

### go embed
Contoh implementasi ketika menggunakan golang embed:
```go
//go:embed resources/ok.html
var resourceOk string

//go:embed resources/notfound.html
var resourceNotFound string

func ServeFileEmbed(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("name") == "" {
		fmt.Fprint(w, resourceNotFound)
	} else {
		fmt.Fprint(w, resourceOk)
	}
}

func TestServeFileServerEmbed(t *testing.T) {
	var server http.Server = http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}

	var err error = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

```