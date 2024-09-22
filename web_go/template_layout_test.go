package web_go

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

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

func TestTemplateLayout(t *testing.T) {
	var request *http.Request = httptest.NewRequest(http.MethodGet, "/", nil)
	var recorder *httptest.ResponseRecorder = httptest.NewRecorder()

	TemplateLayout(recorder, request)

	var body, _ = io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}
