package web_go

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(w http.ResponseWriter, r *http.Request) {
	var name string = r.URL.Query().Get("name")

	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "name is empty")
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello %s", name)
	}
}

func TestResponseCodeInvalid(t *testing.T) {
	var request *http.Request = httptest.NewRequest(http.MethodGet, "/", nil)
	var recorder *httptest.ResponseRecorder = httptest.NewRecorder()

	ResponseCode(recorder, request)

	var response *http.Response = recorder.Result()

	var body []byte
	var _ error
	body, _ = io.ReadAll(response.Body)

	fmt.Println(string(body))
	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
}

func TestResponseCodeValid(t *testing.T) {
	var request *http.Request = httptest.NewRequest(http.MethodGet, "/?name=aldo", nil)
	var recorder *httptest.ResponseRecorder = httptest.NewRecorder()

	ResponseCode(recorder, request)

	var response *http.Response = recorder.Result()

	var body []byte
	var _ error
	body, _ = io.ReadAll(response.Body)

	fmt.Println(string(body))
	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
}
