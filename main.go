package main

import (
	"html/template"
	"log"
	"net/http"
)

var pageView = template.Must(template.ParseGlob("web/*.html"))

func Hello(w http.ResponseWriter, r *http.Request) {
	pageView.ExecuteTemplate(w, "Index", nil)
}

func main() {
	http.HandleFunc("/", Hello)

	log.Println("porta: 8002")
	if err := http.ListenAndServe(":8002", nil); err != nil {
		log.Fatalf("Server exited with: %v", err)
	}
}
