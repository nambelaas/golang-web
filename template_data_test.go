package belajar_golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateDataMap(writer http.ResponseWriter, request *http.Request){
	t:=template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(writer, "name.gohtml", map[string]interface{}{
		"Title": "Template Data Map",
		"Name": "Salman",
	})
}

func TestTemplateDataMap(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080",nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, request)

	response := recorder.Result()
	body,_ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

type Page struct {
	Title string
	Name string
}

func TemplateDataStruct(writer http.ResponseWriter, request *http.Request){
	t:=template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(writer, "name.gohtml", Page{
		Title: "Template Data Map",
		Name: "Salman",
	})
}

func TestTemplateDataStruct(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080",nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, request)

	response := recorder.Result()
	body,_ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}