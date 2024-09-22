package web_go

import (
	"fmt"
	"net/http"
	"testing"
)

func RedirectTo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Redirect")
}

func RedirectFrom(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/redirect-to", http.StatusTemporaryRedirect)
}

func TestRedirect(t *testing.T) {
	var mux *http.ServeMux = http.NewServeMux()
	mux.HandleFunc("/redirect-from", RedirectFrom)
	mux.HandleFunc("/redirect-to", RedirectTo)

	var server http.Server = http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	var err error = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
