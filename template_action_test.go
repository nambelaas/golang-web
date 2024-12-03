package belajar_golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateFunctionIf(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/if.gohtml"))
	t.ExecuteTemplate(writer, "if.gohtml",Page{
		Title: "Test",
		Name: "Salman",
	})
}

func TestTeamplateAction(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet,"http://localhost:8001", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionIf(recorder, request)

	body, _ :=io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

func TemplateFunctionComparator(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/comparator.gohtml"))
	t.ExecuteTemplate(writer, "comparator.gohtml",map[string]interface{}{
		"Title":"Template Function Comparator",
		"FinalValue":90,
	})
}

func TestTeamplateComparator(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet,"http://localhost:8001", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionComparator(recorder, request)

	body, _ :=io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

func TemplateFunctionRange(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/range.gohtml"))
	t.ExecuteTemplate(writer, "range.gohtml",map[string]interface{}{
		"Title":"Template Action Range",
		"Hobbies": []string{
			"Games","Read",
		},
	})
}

func TestTeamplateRange(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet,"http://localhost:8001", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionRange(recorder, request)

	body, _ :=io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

// pada golang template tidak bisa melakukan looping dengan for/foreach sehingga dapat menggunakan with untuk slice/array
func TemplateFunctionWith(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/address.gohtml"))
	t.ExecuteTemplate(writer, "address.gohtml",map[string]interface{}{
		"Title":"Template Action With",
		"Name":"Salman",
		"Address": map[string]interface{}{
			"Street": "Jalan Belum Ada",
			"City": "Smg",
		},
	})
}

func TestTeamplateWith(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet,"http://localhost:8001", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionWith(recorder, request)

	body, _ :=io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}