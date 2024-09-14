package web_go

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(w http.ResponseWriter, r *http.Request) {
	var err error
	err = r.ParseForm()
	if err != nil {
		panic(err)
	}

	var firstName, lastName string
	firstName = r.PostForm.Get("first_name")
	lastName = r.PostForm.Get("last_name")

	fmt.Fprintf(w, "Hello %s %s", firstName, lastName)
}

func TestFormPost(t *testing.T) {
	var formBody *strings.Reader = strings.NewReader("first_name=doe&last_name=aldo")
	var request *http.Request = httptest.NewRequest(http.MethodPost, "/form", formBody)
	var recorder *httptest.ResponseRecorder = httptest.NewRecorder()
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	FormPost(recorder, request)

	var response *http.Response = recorder.Result()

	var body []byte
	var _ error

	body, _ = io.ReadAll(response.Body)

	fmt.Println(string(body))
}
