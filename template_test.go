package belajar_golang_web

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SimpleHtml(writer http.ResponseWriter, request *http.Request)  {
	templateText:= `<html><body>{{.}}</body></html>`

	// t, err := template.New("SIMPLE").Parse(templateText)
	// if err != nil {
	// 	panic(err)
	// }

	t := template.Must(template.New("SIMPLE").Parse(templateText))

	t.ExecuteTemplate(writer, "SIMPLE", "Hello World")
}

func TestSimpleHtml(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080",nil)
	recorder := httptest.NewRecorder()

	SimpleHtml(recorder, request)

	response := recorder.Result()
	body,_ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func SimpleHtmlFile(writer http.ResponseWriter, request *http.Request)  {
	t := template.Must(template.ParseFiles("./templates/simple.gohtml"))

	t.ExecuteTemplate(writer, "simple.gohtml", "Hello World")
}

func TestSimpleHtmlFile(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080",nil)
	recorder := httptest.NewRecorder()

	SimpleHtmlFile(recorder, request)

	response := recorder.Result()
	body,_ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TemplateDirectory(writer http.ResponseWriter, request *http.Request)  {
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))

	t.ExecuteTemplate(writer, "simple.gohtml", "Hello World")
}

func TestTemplateDirectory(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080",nil)
	recorder := httptest.NewRecorder()

	TemplateDirectory(recorder, request)

	response := recorder.Result()
	body,_ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

//go:embed templates/*.gohtml
var templates embed.FS

func TemplateEmbed(writer http.ResponseWriter, request *http.Request)  {
	t, err:= template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(writer, "simple.gohtml", "Hello World")
}

func TestTemplateEmbed(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080",nil)
	recorder := httptest.NewRecorder()

	TemplateEmbed(recorder, request)

	response := recorder.Result()
	body,_ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}