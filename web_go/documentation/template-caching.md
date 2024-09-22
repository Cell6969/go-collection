# Template Caching
Pada real casenya, template akan dicaching sehingga tidak perlu melakukan parsing berulang -ulang. Hal ini membuat web semakin lebih cepat.

Sebagai contoh:
```go
package web_go

import (
	"embed"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

//go:embed templates/*.gohtml
var templates embed.FS

var myTemplates *template.Template = template.Must(template.ParseFS(templates, "templates/*.gohtml"))

func TemplateCaching(w http.ResponseWriter, r *http.Request) {
	myTemplates.ExecuteTemplate(w, "simple.gohtml", nil)
}

func TestTemplateCaching(t *testing.T) {
	var request *http.Request = httptest.NewRequest(http.MethodGet, "/", nil)
	var recorder *httptest.ResponseRecorder = httptest.NewRecorder()

	TemplateCaching(recorder, request)

	var body, _ = io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
```

dengan demikian parsing hanya dilakukan sekali jadi aman dilakukan berulang ulang.