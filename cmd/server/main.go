package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
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

func renderTemplate(w http.ResponseWriter, r *http.Request, templateName string, data TemplateData) {
	// Create template with functions
	tmpl := template.New("").Funcs(funcMap)

	// Parse all template files (this part is the same)
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

	w.Header().Set("X-Title", data.Title)

	if r.Header.Get("HX-Request") == "true" {
		err = parsedTemplate.ExecuteTemplate(w, "content", data)
	} else {
		err = parsedTemplate.ExecuteTemplate(w, "layout", data)
	}

	if err != nil {
		http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	// Serve static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// registering handler
	registerProjectHandlers()

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

		renderTemplate(w, r, "index", data)
	})

	// FAQ page
	http.HandleFunc("/faq", func(w http.ResponseWriter, r *http.Request) {
		data := TemplateData{
			Title:       "FAQ - Deepesh Kalura",
			Active:      "faq",
			CurrentPath: "/faq",
			Data:        map[string]string{"message": "Frequently Asked Questions"},
		}

		renderTemplate(w, r, "faq", data)
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

		renderTemplate(w, r, "food", data)
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

		renderTemplate(w, r, "experience", data)
	})

	http.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		data := TemplateData{
			Title:       "Projects - Deepesh Kalura",
			Active:      "projects", // This is logical, even if not in the main nav
			CurrentPath: "/projects",
			Data:        map[string]string{"message": ""},
		}

		renderTemplate(w, r, "projects", data)
	})

	http.HandleFunc("/changelog", func(w http.ResponseWriter, r *http.Request) {
		// Read the markdown file from disk
		md, err := os.ReadFile("CHANGELOG.md")
		if err != nil {
			http.Error(w, "Could not read changelog file", http.StatusInternalServerError)
			return
		}

		// Configure the markdown parser
		extensions := parser.CommonExtensions | parser.AutoHeadingIDs
		p := parser.NewWithExtensions(extensions)

		// Convert the markdown content to HTML
		html := markdown.ToHTML(md, p, nil)

		data := TemplateData{
			Title:       "Changelog - Deepesh Kalura",
			Active:      "changelog",
			CurrentPath: "/changelog",
			// IMPORTANT: Wrap the HTML in template.HTML to prevent it from being escaped
			Data: template.HTMLEscaper(html),
		}

		renderTemplate(w, r, "changelog", data)
	})

	// Start server
	port := "8080"
	fmt.Printf("Server running at http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
