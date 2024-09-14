# Request
Request di golang merupakan struct yang merepresentasikan HTTP Request. Beberapa componen dari HTTP Request seperti method, header, body, dll

Contoh penggunaan:
```go
var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.Method)
		fmt.Fprintln(w, r.RequestURI)
	}

	server := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
```

## Query Parameter
query parameter biasanya ditempatkan pada URL. Artinya bisa diakses menggunakan Request.

Sebagai contoh, buat handler func nya
```go
func SayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name") // query dari url dan memiliki tipe data map
	if name == "" {
		fmt.Fprint(w, "Hello")
	} else {
		fmt.Fprintf(w, "Hello %s", name)
	}
}
```

Kemudian pada unit test:
```go
func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/hello?name=aldo", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
```

### Multiple Query Parameter
Contoh implementasi:
```go
func MultipleQueryParameter(w http.ResponseWriter, r *http.Request) {
	firstname := r.URL.Query().Get("firstname")
	lastname := r.URL.Query().Get("lastname")

	fmt.Fprintf(w, "Hello %s %s", firstname, lastname)
}

func TestMultipleQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/hello?firstname=aldo&lastname=aldo", nil)
	recorder := httptest.NewRecorder()

	MultipleQueryParameter(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
```


### Multiple Value Query Parameter
```go
func MultpleValueQuery(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	names := query["name"]
	fmt.Fprint(w, strings.Join(names, " "))
}

func TestMultpleValueQuery(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/hello?name=aldo&name=var", nil)
	recorder := httptest.NewRecorder()

	MultpleValueQuery(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
```

## Header
Header adalah informasi tambahan ketika mengirim request dan response. Di golang kita bisa menangkap header dan juga mengatur header untuk dikembalikan. Sebagai contoh:
```go
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
```

Kemudian contoh untuk mengirim header dari server
```go
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
```