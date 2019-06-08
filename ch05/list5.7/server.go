package main

import (
	"net/http"
	"text/template"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html")
	t.Execute(w, "hello")
}

func process2(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl2.html")
	t.Execute(w, "hello")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	http.HandleFunc("/process2", process2)
	server.ListenAndServe()
}
