# Template Layout

Golang mendukung fitur template layout pada gohtml yakni dimana file gohtml bisa mengimport file html yang lain

Sebagai contoh:

buat header.gohtml
```gotemplate
<html>
<head>
    <meta charset="UTF-8">
    <title>{{.Title}}</title>
</head>
<body>
```

footer.gohtml
```gotemplate
</body>
</html>
```

Kemudian untuk content:
```gotemplate
{{template "header.gohtml"}}
<h1>Hello {{.Name}}</h1>
{{template "footer.gohtml"}}
```

Lalu pada code golang:
```go
func TemplateLayout(w http.ResponseWriter, r *http.Request) {
	var t *template.Template = template.Must(template.ParseFiles(
		"./templates/header.gohtml",
		"./templates/footer.gohtml",
		"./templates/layout.gohtml",
	))

	t.ExecuteTemplate(w, "layout.gohtml", map[string]interface{}{
		"Title": "Template layout",
		"Name":  "aldo",
	})
}

func TestTemplateLayout(t *testing.T) {
	var request *http.Request = httptest.NewRequest(http.MethodGet, "/", nil)
	var recorder *httptest.ResponseRecorder = httptest.NewRecorder()

	TemplateLayout(recorder, request)

	var body, _ = io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}
```
note: pada best practice nya jarang menggunakan parseFiles, langsung menggunakan globs.

## Template Name
Template html juga bisa menggunakan aliasing, artinya hanya dengan memanggil nama aliasing kitab bisa merender html tersebut. 

Sebagai contoh pada layout.gohtml
```gotemplate
{{define "layout"}}
{{template "header.gohtml" .}}
<h1>Hello {{.Name}}</h1>
{{template "footer.gohtml"}}
{{end}}
```

Kemudian pada unit test:
```go
func TemplateLayout(w http.ResponseWriter, r *http.Request) {
	var t *template.Template = template.Must(template.ParseFiles(
		"./templates/header.gohtml",
		"./templates/footer.gohtml",
		"./templates/layout.gohtml",
	))

	t.ExecuteTemplate(w, "layout", map[string]interface{}{
		"Title": "Template layout",
		"Name":  "aldo",
	})
}
```