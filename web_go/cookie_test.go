package web_go

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(w http.ResponseWriter, r *http.Request) {
	var cookieAc *http.Cookie = new(http.Cookie)
	cookieAc.Name = "_Secure-Ac"
	cookieAc.Value = r.URL.Query().Get("name")
	cookieAc.Path = "/"
	cookieAc.HttpOnly = true

	http.SetCookie(w, cookieAc)
	fmt.Fprint(w, "success create cookie")
}

func GetCookie(w http.ResponseWriter, r *http.Request) {
	var cookie *http.Cookie
	var err error
	cookie, err = r.Cookie("_Secure-Ac")
	if err != nil {
		fmt.Fprint(w, "No Cookie")
	} else {
		var name string = cookie.Value
		fmt.Fprintf(w, "Hello %s", name)
	}
}

func TestCookie(t *testing.T) {
	var mux *http.ServeMux = http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	var server http.Server = http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	var err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestSetCookie(t *testing.T) {
	var request *http.Request = httptest.NewRequest(http.MethodGet, "/?name=aldo", nil)
	var recorder *httptest.ResponseRecorder = httptest.NewRecorder()

	SetCookie(recorder, request)

	var cookies []*http.Cookie = recorder.Result().Cookies()

	for _, cookie := range cookies {
		fmt.Printf("Cookie Name: %s, Value: %s\n", cookie.Name, cookie.Value)
	}
}

func TestGetCookie(t *testing.T) {
	var request *http.Request = httptest.NewRequest(http.MethodGet, "/?name=aldo", nil)
	var cookie *http.Cookie = new(http.Cookie)
	cookie.Name = "_Secure-Ac"
	cookie.Value = "aldo"
	request.AddCookie(cookie)

	var recorder *httptest.ResponseRecorder = httptest.NewRecorder()
	GetCookie(recorder, request)

	var body []byte
	var _ error

	body, _ = io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}
