package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type TemplateData struct {
	Title  string
	Active string
	Data   interface{}
}

func main() {
	// Setup template parsing
	tmpl := template.Must(template.ParseGlob("templates/*.html"))

	// Serve static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Home page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		data := TemplateData{
			Title:  "Deepesh Kalura - Software Developer",
			Active: "home",
			Data:   map[string]string{"message": "Welcome to my website!"},
		}

		tmpl.ExecuteTemplate(w, "index.html", data)
	})

	// FAQ page
	http.HandleFunc("/faq", func(w http.ResponseWriter, r *http.Request) {
		data := TemplateData{
			Title:  "FAQ - Deepesh Kalura",
			Active: "faq",
			Data:   map[string]string{"message": "Frequently Asked Questions"},
		}

		tmpl.ExecuteTemplate(w, "faq.html", data)
	})

	// Start server
	port := "8080"
	fmt.Printf("Server running at http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
