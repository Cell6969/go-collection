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
