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
