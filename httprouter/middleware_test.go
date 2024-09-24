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
