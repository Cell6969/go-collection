package web_go

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

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
