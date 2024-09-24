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

func TestPanicHandler(t *testing.T) {
	var router *httprouter.Router = httprouter.New()

	// handling error
	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, err interface{}) {
		fmt.Fprint(w, "Panic: ", err.(string))
	}

	// set router
	router.GET("/error", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		panic("ini error")
	})

	var request *http.Request = httptest.NewRequest(http.MethodGet, "/error", nil)
	var recorder *httptest.ResponseRecorder = httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	var response *http.Response = recorder.Result()

	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Panic: ini error", string(body))
}
