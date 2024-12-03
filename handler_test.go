package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		// logic web
		fmt.Fprint(w, "Hello World")
	}

	server := http.Server{
		Addr:    "localhost:8081",
		Handler: handler,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

// untuk membuat endpoint lebih dari satu lebih baik menggunakan ServeMux karena fungsinya sama seperti router pada bahasa pemrograman lain
func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "Hello World")
	})
	mux.HandleFunc("/hi", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "Hi")
	})
	// jika pattern berakhiran "/" maka semua endpoint setelahnya akan tetap menggunakan func dari enpoint pertama
	mux.HandleFunc("/images/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "Images")
	})
	// jika enpoint lebih panjang maka yang dibaca adalah endpoint yang terpanjangnya
	mux.HandleFunc("/images/thumbnails/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "Images thumbnail")
	})

	server := http.Server{
		Addr:    "localhost:8082",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func TestRequest(t *testing.T){
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.Method)
		fmt.Fprintln(w, r.RequestURI)
	}

	server := http.Server{
		Addr: "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()

	if err != nil{
		panic(err)
	}
}