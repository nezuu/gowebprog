package main

import (
	"html/template"
	"net/http"
)

// 表示するものがスライスの場合
func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html")
	daysOfWeek := []string{"月", "火", "水", "木", "金", "土", "日"}
	t.Execute(w, daysOfWeek)
}

// 表示するものがない場合
func process2(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl2.html")
	daysOfWeek := []string{}
	// daysOfWeek := []string{"月", "火", "水", "木", "金", "土", "日"}
	t.Execute(w, daysOfWeek)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	http.HandleFunc("/process2", process2)
	server.ListenAndServe()
}
