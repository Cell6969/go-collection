# Unit Test
Golang menyediakan fitur untuk melakukan unit test terhadap http request. Hal ini bisa berjalan menggunakan library dari net/http/httptest.

Sebagai contoh implementasinya:
```go
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
```