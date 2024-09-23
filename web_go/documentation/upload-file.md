# Upload File

Golang mendukung untuk input berupa file dengan menggunakan Multipart. Contoh implementasinya:

buat file html form dan success:

form.gohtml
```gotemplate
<html>
<head>
    <meta charset="UTF-8">
    <title>Form Upload File</title>
</head>
<body>
<h1>Upload File</h1>
<form action="/upload" method="post" enctype="multipart/form-data">
    <label>Name: <input type="text" name="name"/></label><br>
    <label>File: <input type="file" name="file"/></label>
    <input type="submit" value="Upload"/>
</form>
</body>
</html>
```

success.gohtml
```gotemplate
<html>
<head>
    <meta charset="UTF-8">
    <title>Success</title>
</head>
<body>
<h1>{{.Name}}</h1>
<a href="{{.File}}">File</a>
</body>
</html>
```

Kemudian untuk code golang:
```go
package web_go

import (
	"io"
	"net/http"
	"os"
	"testing"
)

func UploadForm(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "upload.form.gohtml", nil)
}

func Upload(w http.ResponseWriter, r *http.Request) {
	//r.ParseMultipartForm(32 << 20)
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}

	fileDestination, err := os.Create("./resources/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(fileDestination, file)
	if err != nil {
		panic(err)
	}

	var name string = r.PostFormValue("name")
	myTemplates.ExecuteTemplate(w, "upload.success.gohtml", map[string]string{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})
}

func TestUploadForm(t *testing.T) {
	var mux *http.ServeMux = http.NewServeMux()
	mux.HandleFunc("/", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources"))))

	var server http.Server = http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	var err error = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
```

## Unit test upload file
bisa juga melakukan unit test untuk upload file, sebagai contoh:
```go
//go:embed resources/main.pdf
var uploadFileTest []byte

func TestUploadFile(t *testing.T) {
	// initiate body for input
	var body = new(bytes.Buffer)
	var writer = multipart.NewWriter(body)
	writer.WriteField("name", "aldo")
	file, _ := writer.CreateFormFile("file", "contohupload.png")
	file.Write(uploadFileTest)
	writer.Close()

	// make request
	var request *http.Request = httptest.NewRequest(http.MethodPost, "/upload", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	var recorder *httptest.ResponseRecorder = httptest.NewRecorder()

	Upload(recorder, request)

	bodyResponse, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(bodyResponse))
}
```