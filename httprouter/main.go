package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

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
