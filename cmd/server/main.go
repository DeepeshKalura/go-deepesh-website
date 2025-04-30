package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

type TemplateData struct {
	Title       string
	Active      string
	CurrentPath string
	Data        interface{}
}

// Define template functions
var funcMap = template.FuncMap{
	// You can add template functions here if needed
	// For example: "formatDate": formatDateFunc,

	"isActive": func(currentPath, path string) bool {
		return currentPath == path
	},
}

func renderTemplate(w http.ResponseWriter, templateName string, data TemplateData) {
	// Create template with functions
	tmpl := template.New("").Funcs(funcMap)

	// Parse all template files
	templateFiles := []string{
		filepath.Join("templates", "layout.html"),
		filepath.Join("templates", "components", "nav.html"),
		filepath.Join("templates", "components", "footer.html"),
		filepath.Join("templates", "pages", templateName+".html"),
	}

	parsedTemplate, err := tmpl.ParseFiles(templateFiles...)
	if err != nil {
		http.Error(w, "Error parsing template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Set custom header for title
	w.Header().Set("X-Title", data.Title)

	// Execute the template
	err = parsedTemplate.ExecuteTemplate(w, "layout", data)
	if err != nil {
		http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
	}
}

func main() {
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
			Title:       "Deepesh Kalura - Software Developer",
			Active:      "home",
			CurrentPath: "/",
			Data:        map[string]string{"message": "Welcome to my website!"},
		}

		renderTemplate(w, "index", data)
	})

	// FAQ page
	http.HandleFunc("/faq", func(w http.ResponseWriter, r *http.Request) {
		data := TemplateData{
			Title:       "FAQ - Deepesh Kalura",
			Active:      "faq",
			CurrentPath: "/faq",
			Data:        map[string]string{"message": "Frequently Asked Questions"},
		}

		renderTemplate(w, "faq", data)
	})

	// Projects page
	http.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		data := TemplateData{
			Title:       "Projects Deepesh has communicated so far",
			Active:      "projects",
			CurrentPath: "/projects",
			Data:        map[string]string{"message": ""},
		}

		renderTemplate(w, "projects", data)
	})

	// Favicon
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		http.ServeFile(w, r, "static/img/favicon.ico")
	})

	// Food page
	http.HandleFunc("/food", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		data := TemplateData{
			Title:       "Cooking With Kalura Deepesh",
			Active:      "food",
			CurrentPath: "/food",
			Data:        map[string]string{"message": ""},
		}

		renderTemplate(w, "food", data)
	})

	// Experience page
	http.HandleFunc("/experience", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		data := TemplateData{
			Title:       "Experience - Deepesh Kalura",
			Active:      "experience",
			CurrentPath: "/experience",
			Data:        map[string]string{"message": ""},
		}

		renderTemplate(w, "experience", data)
	})

	// Start server
	port := "8080"
	fmt.Printf("Server running at http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
