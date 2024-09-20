package web_go

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TODO: ActionIf
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

// TODO: Operator

func TemplateActionOperator(w http.ResponseWriter, r *http.Request) {
	var t = template.Must(template.ParseFiles("./templates/comparator.gohtml"))
	t.ExecuteTemplate(w, "comparator.gohtml", map[string]interface{}{
		"Title":      "Template Operator",
		"FinalValue": 70,
	})
}

func TestTemplateActionOperator(t *testing.T) {
	var request = httptest.NewRequest(http.MethodGet, "/operator.gohtml", nil)
	var recorder = httptest.NewRecorder()

	TemplateActionOperator(recorder, request)

	var body, _ = io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

// TODO: Range

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

// TODO: WITH
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
