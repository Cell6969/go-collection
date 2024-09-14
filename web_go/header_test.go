package web_go

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(w http.ResponseWriter, r *http.Request) {
	var contenType string = r.Header.Get("content-type")
	fmt.Fprintln(w, contenType)
}

func TestRequestHeader(t *testing.T) {
	var request *http.Request = httptest.NewRequest(http.MethodPost, "/header", nil)
	request.Header.Add("content-type", "application/json")

	var recorder *httptest.ResponseRecorder = httptest.NewRecorder()

	RequestHeader(recorder, request)

	var response *http.Response = recorder.Result()
	var body []byte
	var _ error

	body, _ = io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func ResponseHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-Powered-By", "Go")
	fmt.Fprint(w, "ok")
}

func TestResponseHeader(t *testing.T) {
	var request *http.Request = httptest.NewRequest(http.MethodPost, "/header", nil)
	var recorder *httptest.ResponseRecorder = httptest.NewRecorder()

	ResponseHeader(recorder, request)

	var body []byte
	var _ error

	body, _ = io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))

	fmt.Println(recorder.Header().Get("X-Powered-By"))
}
