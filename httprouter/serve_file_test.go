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
