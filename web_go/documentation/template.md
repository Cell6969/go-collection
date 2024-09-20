# Template
golang mendukung untuk merender template engine. Untuk implementasi sederhananya:
```go
func SimpleHtml(w http.ResponseWriter, req *http.Request) {
	var templateText string = `<html><body>{{.}}</body></html>`
	var t *template.Template = template.Must(template.New("SIMPLE").Parse(templateText))
	t.ExecuteTemplate(w, "SIMPLE", "Hello HTML Template")
}

func TestSimpleHTML(t *testing.T) {
	var request *http.Request = httptest.NewRequest(http.MethodGet, "/", nil)
	var recorder *httptest.ResponseRecorder = httptest.NewRecorder()

	SimpleHtml(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
```

## Template File
Template juga bisa dalam file. Sebagai contoh buat file simple.gohtml
```gotemplate
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{.}}</title>
</head>
<body>
<h1>{{.}}</h1>
</body>
</html>
```

Kemudian pada code golang:
```go
func SimpleHtmlFile(w http.ResponseWriter, req *http.Request) {
	var t *template.Template = template.Must(template.ParseFiles("./templates/simple.gohtml"))
	t.ExecuteTemplate(w, "simple.gohtml", "Hello HTML Template")
}

func TestSimpleHTMLFile(t *testing.T) {
	var request *http.Request = httptest.NewRequest(http.MethodGet, "/", nil)
	var recorder *httptest.ResponseRecorder = httptest.NewRecorder()

	SimpleHtmlFile(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
```

## Template Directory
Jikalau ingin langsung meload template dari directory maka bisa dilakukan oleh golang. sebagai contoh:
```go
func TemplateDirectory(w http.ResponseWriter, req *http.Request) {
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))
	t.ExecuteTemplate(w, "simple.gohtml", "Hello HTML Template")
}

func TestTemplateDirectory(t *testing.T) {
	var request *http.Request = httptest.NewRequest(http.MethodGet, "/", nil)
	var recorder *httptest.ResponseRecorder = httptest.NewRecorder()
	TemplateDirectory(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
```

### Template go-embed
Contoh implementasi ketika menggunakan fitur embed:
```go
//go:embed templates/*.gohtml
var templates embed.FS

func TemplateEmbed(w http.ResponseWriter, req *http.Request) {
	var t *template.Template = template.Must(template.ParseFS(templates, "templates/*.gohtml"))
	t.ExecuteTemplate(w, "simple.gohtml", "Hello HTML Template")
}

func TestTemplateEmbed(t *testing.T) {
	var request *http.Request = httptest.NewRequest(http.MethodGet, "/", nil)
	var recorder *httptest.ResponseRecorder = httptest.NewRecorder()
	TemplateEmbed(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
```

# Template Data
Data yang dimasukkan dalam template bisa lebih dari 1 maka disarankan menamakan variable data di template html. Sebagai contoh:
```gotemplate
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{.Title}}</title>
</head>
<body>
<h1>Hello {{.Name}}</h1>
</body>
</html>
```

kemudian untuk memasukkan data:
```go
func TemplateDataMap(w http.ResponseWriter, r *http.Request) {
	var t *template.Template = template.Must(template.ParseFiles("./templates/name.gohtml"))

	//Memasukkan datanya menggunakan map
	t.ExecuteTemplate(w, "name.gohtml", map[string]interface{}{
		"Title": "Template Data Map",
		"Name":  "Aldo",
	})
}

func TestTemplateDataMap(t *testing.T) {
	var request = httptest.NewRequest(http.MethodGet, "/", nil)
	var recorder = httptest.NewRecorder()

	TemplateDataMap(recorder, request)

	var body, _ = io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}
```
Jadi dengan mengirimkan data map , kita bisa memasukkan data sesuai nama variable yang ada pada template html. Selain data map bisa juga memasukkan dari struct.

Contoh implementasi:
```go
type Page struct {
	Title string
	Name  string
}

func TemplateDataStruct(w http.ResponseWriter, r *http.Request) {
	var t = template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(w, "name.gohtml", Page{
		Title: "Ini Struct",
		Name:  "Aldo",
	})
}

func TestTemplateDataStruct(t *testing.T) {
	var request = httptest.NewRequest(http.MethodGet, "/", nil)
	var recorder = httptest.NewRecorder()
	TemplateDataStruct(recorder, request)

	var body, _ = io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
```

Jikalau ternyata berupa nested object misal:
```gotemplate
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{.Title}}</title>
</head>
<body>
<h1>Hello {{.Name}}</h1>
<h2>Alamat : {{.Address.Street}}</h2>
</body>
</html>
```

Kemudian pada code golang:
```go
type Address struct {
	Street string
}

type Page struct {
	Title   string
	Name    string
	Address Address
}

func TemplateDataStruct(w http.ResponseWriter, r *http.Request) {
	var t = template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(w, "name.gohtml", Page{
		Title: "Ini Struct",
		Name:  "Aldo",
		Address: Address{
			Street: "Belum Ada",
		},
	})
}

func TestTemplateDataStruct(t *testing.T) {
	var request = httptest.NewRequest(http.MethodGet, "/", nil)
	var recorder = httptest.NewRecorder()
	TemplateDataStruct(recorder, request)

	var body, _ = io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
```
Hal ini berlaku juga jika ingin memasukkan dengan map.

## Template Action
Template engine golang juga mendukung beberapa action seperti if-else, loop,dll.

### If-else
**gohtml**
```gotemplate
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{.Title}}</title>
</head>
<body>
{{if .Name}}
    <h1>Hello {{.Name}}</h1>
{{else}}
    <h1>Hello</h1>
{{end}}
</body>
</html>
```

Kemudian pada code golang:
```go
func TemplateActionIf(w http.ResponseWriter, r *http.Request) {
	var t = template.Must(template.ParseFiles("./templates/if.gohtml"))
	t.ExecuteTemplate(w, "if.gohtml", map[string]interface{}{
		"Title": "Template if Action",
	})
}

func TestTemplateActionIf(t *testing.T) {
	var request = httptest.NewRequest(http.MethodGet, "/if.gohtml", nil)
	var recorder = httptest.NewRecorder()
	TemplateActionIf(recorder, request)

	var body, _ = io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
```

### Operator Comparison
**gohtml**
```gotemplate
<html>
<head>
    <meta charset="UTF-8">
    <title>{{.Title}}</title>
</head>
<body>
    {{if ge .FinalValue 80}}
        <h1>Good</h1>
    {{else if go .FinalValue 60}}
        <h1>Nice Try</h1>
    {{else}}
        <h1>Try Again</h1>
    {{end}}
</body>
</html>
```

Kemudian pada code go lang:
```go
func TemplateActionOperator(w http.ResponseWriter, r *http.Request) {
	var t = template.Must(template.ParseFiles("./templates/comparator.gohtml"))
	t.ExecuteTemplate(w, "comparator.gohtml", map[string]interface{}{
		"Title":      "Template Operator",
		"FinalValue": 90,
	})
}

func TestTemplateActionOperator(t *testing.T) {
	var request = httptest.NewRequest(http.MethodGet, "/operator.gohtml", nil)
	var recorder = httptest.NewRecorder()

	TemplateActionOperator(recorder, request)

	var body, _ = io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
```

### Range
```gotemplate
<html>
<head>
    <meta charset="UTF-8">
    <title>{{.Title}}</title>
</head>
<body>
{{range $index,$element := .Hobbies}}
    <h1>Hobby : {{$element}}</h1>
{{else}}
    <h1>You don't have a hobbies</h1>
{{end}}
</body>
</html>
```

```go
func TemplateActionRange(w http.ResponseWriter, r *http.Request) {
	var t = template.Must(template.ParseFiles("./templates/range.gohtml"))
	t.ExecuteTemplate(w, "range.gohtml", map[string]interface{}{
		"Title": "Template Range",
		"Hobbies": []string{
			"Game", "Coding", "Reading",
		},
	})
}

func TestTemplateActionRange(t *testing.T) {
	var request = httptest.NewRequest(http.MethodGet, "/range.gohtml", nil)
	var recorder = httptest.NewRecorder()
	TemplateActionRange(recorder, request)

	var body, _ = io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
```

### With
```gotemplate
<html>
<head>
    <meta charset="UTF-8">
    <title>{{.Title}}</title>
</head>
<body>
<h1>Name : {{.Name}}</h1>
{{with .Address}}
    <h1>Address Street: {{.Street}}</h1>
    <h1>Address City: {{.City}}</h1>
{{end}}
</body>
</html>
```

```go
func TemplateActionWith(w http.ResponseWriter, r *http.Request) {
	var t *template.Template = template.Must(template.ParseFiles("./templates/with.gohtml"))
	t.ExecuteTemplate(w, "with.gohtml", map[string]interface{}{
		"Title": "Template With",
		"Name":  "Aldo",
		"Address": map[string]interface{}{
			"Street": "bintara",
			"City":   "bekasi",
		},
	})
}

func TestTemplateActionWith(t *testing.T) {
	var request *http.Request = httptest.NewRequest(http.MethodGet, "/with.gohtml", nil)
	var recorder *httptest.ResponseRecorder = httptest.NewRecorder()

	TemplateActionWith(recorder, request)

	var body, _ = io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
```