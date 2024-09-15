package web_go

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

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
