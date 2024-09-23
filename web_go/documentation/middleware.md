# Middleware
Golang juga menerapkan konsep middleware dalam pembuatan aplikasi web nya. Tetapi di golang tidak mengenal konsep middleware tetapi handler. Artinya tidak eksplisit menyebutkan middleware tetapi mengikuti interface handler.

Sebagai contoh:
```go
package web_go

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Middleware:", "Before execute handler")
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("Middleware:", "After execute handler")
}

func TestMiddleware(t *testing.T) {
	var mux = http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("handler executed")
		fmt.Fprint(writer, "Hello middleware")
	})

	mux.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Foo:handler executed")
		fmt.Fprint(writer, "Hello Foo")
	})

	// initiate middleware
	var middleware = &LogMiddleware{
		Handler: mux,
	}

	var server = http.Server{
		Addr:    ":8080",
		Handler: middleware,
	}

	var err error = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
```

## Error Handler
di middleware juga bisa melakukan error handler, yaitu dengan mengubah panic error menjadi response error.

Contoh implementasinya:
```go
package web_go

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Middleware:", "Before execute handler")
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("Middleware:", "After execute handler")
}

type ErrorHandler struct {
	Handler http.Handler
}

func (errorHandler *ErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Terjadi Error")
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "Error : %s", err)
		}
	}()
	errorHandler.Handler.ServeHTTP(writer, request)
}

func TestMiddleware(t *testing.T) {
	var mux = http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("handler executed")
		fmt.Fprint(writer, "Hello middleware")
	})

	mux.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Foo:handler executed")
		fmt.Fprint(writer, "Hello Foo")
	})

	mux.HandleFunc("/error", func(writer http.ResponseWriter, request *http.Request) {
		panic("Error sengaja")
	})
	// initiate middleware
	var logMiddleware = &LogMiddleware{
		Handler: mux,
	}

	var errorHandler = &ErrorHandler{
		Handler: logMiddleware,
	}

	var server = http.Server{
		Addr:    ":8080",
		Handler: errorHandler,
	}

	var err error = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

```

Jadi singkatnya kita bikin handler yang khusus menangani error kemudian, handler tersebut diregistrasikan ke middleware. Urutannya seperti code diatas:

server -> error handler -> log middleware -> mux.

Ketika terjadi error di mux , maka error naik ke log tetapi karena log tidak menghandle error maka naik lagi ke error handler.