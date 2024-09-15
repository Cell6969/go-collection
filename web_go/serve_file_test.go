package web_go

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

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
