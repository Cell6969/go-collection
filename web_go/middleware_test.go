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
