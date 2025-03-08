package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type TemplateData struct {
	Title string
	Data  interface{}
}

func main() {
	// Setup template parsing
	tmpl := template.Must(template.ParseGlob("templates/*.html"))

	// Serve static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Home page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TemplateData{
			Title: "Deepesh's Website",
			Data:  map[string]string{"message": "Welcome to my website!"},
		}
		tmpl.ExecuteTemplate(w, "index.html", data)
	})

	// Start server
	port := "8080"
	fmt.Printf("Server running at http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

	http.HandleFunc("/api/greet", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`
			<div class="bg-green-100 border border-green-400 text-green-700 px-4 py-3 rounded">
				<p class="font-bold">Hello from HTMX!</p>
				<p>This content was loaded dynamically without a page refresh.</p>
				<button
					hx-get="/api/greet"
					hx-trigger="click"
					hx-swap="outerHTML"
					class="mt-2 bg-green-500 hover:bg-green-600 text-white font-bold py-1 px-3 rounded text-sm">
					Reload
				</button>
			</div>
		`))
	})
}
