# Template Function
Selain value, di template juga bisa memanggil function entah dari struct atau anonymus function pada map. Jadi pada template html juga bisa menjalankan function dan value.

Sebagai contoh, pada struct berikut:
```go
type MyPage struct {
	Name string
}

func (myPage MyPage) SayHello(name string) string {
	return "Hello " + name + ", My Name is" + myPage.Name
}
```

lalu untuk test sederhana:
```go
func TemplateFunction(w http.ResponseWriter, r *http.Request) {
	var t *template.Template = template.Must(template.New("FUNCTION").Parse(`{{.SayHello "Aldo"}}`))
	t.ExecuteTemplate(w, "FUNCTION", MyPage{Name: "Ali"})
}

func TestTemplateFunction(t *testing.T) {
	var request *http.Request = httptest.NewRequest(http.MethodGet, "/", nil)
	var recorder *httptest.ResponseRecorder = httptest.NewRecorder()

	TemplateFunction(recorder, request)

	var body, _ = io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
```

## Global Function
Golang sudah menyediakan beberapa global function seperti len, if , dll. Sebagai contoh:
```go
func TemplateFunctionGlobal(w http.ResponseWriter, r *http.Request) {
	var t *template.Template = template.Must(template.New("FUNCTION").Parse(`{{len .Name}}`))
	t.ExecuteTemplate(w, "FUNCTION", MyPage{Name: "Ali"})
}

func TestTemplateFunctionGlobal(t *testing.T) {
	var request *http.Request = httptest.NewRequest(http.MethodGet, "/", nil)
	var recorder *httptest.ResponseRecorder = httptest.NewRecorder()

	TemplateFunctionGlobal(recorder, request)

	var body, _ = io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
```
Note : klo global function tidak perlu memakai titik didepan.

## Menambah global function
Bisa juga untuk menambah global function. Sebagai contoh:
```go
func TemplateFunctionCreateGlobal(w http.ResponseWriter, r *http.Request) {
	var t *template.Template = template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})

	t = template.Must(t.Parse(`{{ upper .Name}}`))
	t.ExecuteTemplate(w, "FUNCTION", MyPage{Name: "Aldo"})
}

func TestTemplateFunctionCreateGlobal(t *testing.T) {
	var request *http.Request = httptest.NewRequest(http.MethodGet, "/", nil)
	var recorder *httptest.ResponseRecorder = httptest.NewRecorder()
	TemplateFunctionCreateGlobal(recorder, request)

	var body, _ = io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
```
Jadi sebelum memparsing file, terlebih dahulu untuk meregristrasikan function tersebut.

## Function Pipelines
Function Pipelines merupakan function pada golang yang berfungsi untuk meneruskan dari 1 function ke function lain.
Sebagai contoh:
```go
func TemplateFunctionCreatePipeline(w http.ResponseWriter, r *http.Request) {
	var t *template.Template = template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"sayHello": func(name string) string {
			return "Hello Mr " + name
		},
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})

	t = template.Must(t.Parse(`{{ sayHello .Name | upper}}`))
	t.ExecuteTemplate(w, "FUNCTION", MyPage{Name: "Aldo"})
}

func TestTemplateFunctionCreatePipeline(t *testing.T) {
	var request *http.Request = httptest.NewRequest(http.MethodGet, "/", nil)
	var recorder *httptest.ResponseRecorder = httptest.NewRecorder()
	TemplateFunctionCreatePipeline(recorder, request)

	var body, _ = io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
```