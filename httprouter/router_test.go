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
