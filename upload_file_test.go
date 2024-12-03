package belajar_golang_web

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	_ "embed"
)

func UploadForm(w http.ResponseWriter, r *http.Request) {
	err := myTemplates.ExecuteTemplate(w, "upload.form.gohtml",nil)
	if err != nil{
		panic(err)
	}
}

func Upload(w http.ResponseWriter, r *http.Request){
	// ambil file dari input
	file,fileHeader,err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	// create destinasi penyimpanan
	fileDestination,err := os.Create("./resources/"+fileHeader.Filename)
	if err != nil {
		panic(err)
	}
	// simpan file ke destinasi penyimpanannya
	_, err = io.Copy(fileDestination, file)
	if err != nil {
		panic(err)
	}
	name := r.PostFormValue("name")
	myTemplates.ExecuteTemplate(w, "upload.success.gohtml", map[string]interface{}{
		"Name":name,
		"File": "/static/"+fileHeader.Filename,
	})
}

func TestUploadForm(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/",UploadForm)
	mux.HandleFunc("/upload",Upload)
	mux.Handle("/static",http.StripPrefix("/static",http.FileServer((http.Dir("./resources")))))

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources/abc.jpg
var uploadFileTest []byte

func TestUploadFile(t *testing.T)  {
	body:= new(bytes.Buffer)

	writer := multipart.NewWriter(body)
	writer.WriteField("name","Salman S")
	file, _:= writer.CreateFormFile("file", "abc2.jpg")
	file.Write(uploadFileTest)
	writer.Close()

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080",body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	recorder := httptest.NewRecorder()

	Upload(recorder, request)

	response := recorder.Result()
	bodyResponse,_ := io.ReadAll(response.Body)

	fmt.Println(string(bodyResponse))
}