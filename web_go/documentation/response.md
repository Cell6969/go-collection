# Response

## Response Code
Http Response memiliki status code untuk memberi status hasil request. Di golang bisa menentukan response code atau status code yang akan dikirim kembalik ke client
```go
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

```

## Cookie
Golang bisa mengatur cookie (server) , sebagai contoh buat function untuk set cookie dan get cookie:
```go
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
```

Kemudian pada unit test:
```go
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
```