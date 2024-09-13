package web_go

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func TestHttp(t *testing.T) {
	// Buat request, url beserta method dan body
	request := httptest.NewRequest(http.MethodGet, "/hello", nil)
	// buat recorder
	recorder := httptest.NewRecorder()

	// implement request dan recorder pada Handler
	HelloHandler(recorder, request)

	// dapatkan response
	response := recorder.Result()

	// baca hasil response dalam bentuk byte
	body, _ := io.ReadAll(response.Body)

	// convert byte menjadi string
	bodyString := string(body)

	fmt.Println(bodyString)
}
