# Intro

Httprouter merupakan library untuk melakukan routing secara minimalis dan cepat. Berbeda dengan servemux, httprouter perlu diinstall packagenya terlebih dahulu.

untuk instalasi httprouter dan testify:
```shell
go get github.com/julienschmidt/httprouter

go get github.com/stretchr/testify
```

## Router
Untuk menjalankan webserver dari httprouter, kita hanya perlu memanggil router.

Contoh implementasinya:
```go
func main() {
	var router *httprouter.Router = httprouter.New()

	var server http.Server = http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe()
}
```

Cukup sederhana untuk menjalankan webserver dengan httprouter.

### HTTP Method
Pada router dapat method untuk mendefine method dari suatu endpoint. Sebagai contoh:
```go
func main() {
	var router *httprouter.Router = httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Hello HttpRouter")
	})

	var server http.Server = http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe()
}
```
Jadi mirip dengan serve mux, namun terdapat 3 parameter yaitu request, response dan parameter.

### Unit Test
Bisa juga dilakukan unit test menggunakan httprouter, sebagai contoh:
```go
package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter(t *testing.T) {
	// Initiate route
	var router *httprouter.Router = httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Hello world")
	})

	// initiate request
	var request *http.Request = httptest.NewRequest(http.MethodGet, "/", nil)
	var recorder *httptest.ResponseRecorder = httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	var response *http.Response = recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
	assert.Equal(t, "Hello world", string(body))
}
```
Cukup sederhana, mirip ketika menggunakan serve mux.

## Params
Params disini berfungsi untuk melakukan dynamic url. Hal ini tidak bisa dilakukan oleh serve mux, namun bisa dilakukan oleh httprouter. Berbeda dengan query parameter, parameter ini lebih ke dynamic url misal /product/1, /product/2, dst.

Contoh implementasi:
```go
func TestParams(t *testing.T) {
	var router *httprouter.Router = httprouter.New()
	router.GET("/products/:id", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		var id string = params.ByName("id")
		var text string = "Product " + id
		fmt.Fprintf(writer, text)
	})

	var request *http.Request = httptest.NewRequest("GET", "/products/123", nil)
	var recorder *httptest.ResponseRecorder = httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	var response = recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Product 123", string(body))
}
```

## Router Pattern
Pola router yang sering dibuat adalah named parameter dan catch all. Singkatnya named itu pola parameter dengan menggunakan nama sedangkan catch all itu menangkap semua parameter pada url. Named bisa ditaruh ditengah dan diakhir, sedangkan catch all hanya bisa ditaruh dibagian akhir.

## Serve File
Httprouter mendukung untuk melakukan serve file. 

Contoh implementasi:
```go
package main

import (
	"embed"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed resources
var resources embed.FS

func TestServeFileHello(t *testing.T) {
	var router *httprouter.Router = httprouter.New()
	directory, _ := fs.Sub(resources, "resources")
	router.ServeFiles("/files/*filepath", http.FS(directory))

	var request *http.Request = httptest.NewRequest(http.MethodGet, "/files/hello.txt", nil)
	var recorder *httptest.ResponseRecorder = httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	var response *http.Response = recorder.Result()
	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Hello World", string(body))
}

func TestServeFileGoodBye(t *testing.T) {
	var router *httprouter.Router = httprouter.New()
	directory, _ := fs.Sub(resources, "resources")
	router.ServeFiles("/files/*filepath", http.FS(directory))

	var request *http.Request = httptest.NewRequest(http.MethodGet, "/files/goodbye.txt", nil)
	var recorder *httptest.ResponseRecorder = httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	var response *http.Response = recorder.Result()
	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Good Bye", string(body))
}
```

## Panic Handler
httprouter juga menyediakan panic handler yakni handler yang khusus menangani error ketika terjadi panic.

Contoh implementasi:
```go
func TestPanicHandler(t *testing.T) {
	var router *httprouter.Router = httprouter.New()

	// handling error
	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, err interface{}) {
		fmt.Fprint(w, "Panic: ", err.(string))
	}

	// set router
	router.GET("/error", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		panic("ini error")
	})

	var request *http.Request = httptest.NewRequest(http.MethodGet, "/error", nil)
	var recorder *httptest.ResponseRecorder = httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	var response *http.Response = recorder.Result()

	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Panic: ini error", string(body))
}
```

## Notfound Handler
Selain panic handler, terdapat juga notfound handler. Artinya handler tersebut untuk berfungsi untuk menghandle ketika terjadi request namun tidak match url sehingga notfound. Di httprouter juga disediakan untuk menghandle custom notfound.

Contoh implementasi:
```go
func TestNotFound(t *testing.T) {
	var router *httprouter.Router = httprouter.New()
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "This is 404")
	})

	var request *http.Request = httptest.NewRequest(http.MethodGet, "/", nil)
	var recorder *httptest.ResponseRecorder = httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	var response *http.Response = recorder.Result()

	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "This is 404", string(body))
}
```

## Method Not Allowed Handler
Di httprouter, kita bisa menentukan method apa aja yang terjadi pada tiap route. Jikalau user menggunaman method yang tidak sesuai dengann yang di define maka bisa dilarikan ke not allowed handler. Not Allowed handler juga bisa dicustom sesuai keinginan. 

Contoh implementasi:
```go
func TestMethodNotAllowed(t *testing.T) {
	var router *httprouter.Router = httprouter.New()
	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Method tidak diizinkan")
	})

	router.POST("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "POST")
	})

	var request *http.Request = httptest.NewRequest(http.MethodGet, "/", nil)
	var recorder *httptest.ResponseRecorder = httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	var response *http.Response = recorder.Result()

	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Method tidak diizinkan", string(body))
}
```

## Middleware
Di httprouter tidak ada fitur spesifik untuk membuat middleware. Namun karena httprouter merupakan implementasi httphandler dan middleware merupakan handler, maka bisa dibuat secara manual.

Sebagai contoh:
```go
package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type LogMiddleware struct {
	http.Handler
}

func (middleware LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Receive Request")
	//forward request
	middleware.Handler.ServeHTTP(w, r)
}

func TestMiddleware(t *testing.T) {
	var router *httprouter.Router = httprouter.New()

	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Hello Word")
	})

	// masukkan router ke middleware
	var middleware LogMiddleware = LogMiddleware{
		router,
	}

	// buat request
	var request *http.Request = httptest.NewRequest("GET", "/", nil)
	var recorder *httptest.ResponseRecorder = httptest.NewRecorder()
	middleware.ServeHTTP(recorder, request)

	var response = recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Hello Word", string(body))
}
```
Jadi mirip dengan yang dilakukan pada serve mux. Secara urutan pattern:

server request -> middleware -> router

Jikalau ada middleware lagi maka logMiddleware dibungkus lagi dengann middleware lain dst.