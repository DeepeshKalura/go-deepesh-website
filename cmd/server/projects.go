package main

import (
	"net/http"
	"path/filepath"
	"text/template"
)

// renderProjectTemplate uses a different layout for project pages.
func renderProjectTemplate(w http.ResponseWriter, r *http.Request, templateName string, data TemplateData) {
	// Create template with functions
	tmpl := template.New("").Funcs(funcMap)

	// Parse project-specific template files
	templateFiles := []string{
		filepath.Join("templates", "project_layout.html"),
		filepath.Join("templates", "components", "footer.html"), // Using project-specific nav
		filepath.Join("templates", "pages", templateName+".html"),
	}

	parsedTemplate, err := tmpl.ParseFiles(templateFiles...)
	if err != nil {
		http.Error(w, "Error parsing project template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the title header before rendering
	w.Header().Set("X-Title", data.Title)

	// HTMX-aware rendering logic
	if r.Header.Get("HX-Request") == "true" {
		// If it's an HTMX request from the main site, only send the content block
		// so it can be swapped into the #content-area div.
		err = parsedTemplate.ExecuteTemplate(w, "content", data)
	} else {
		// For a direct, full page load, send the entire project_layout
		err = parsedTemplate.ExecuteTemplate(w, "project_layout", data)
	}

	// Handle any execution errors from the if/else block
	if err != nil {
		http.Error(w, "Error executing project template: "+err.Error(), http.StatusInternalServerError)
	}
}

// registerProjectHandlers sets up the routes for the projects section.
func registerProjectHandlers() {
	// Note: The main projects overview page might still use the main layout.
	// This handler is for the individual project detail pages.
	http.HandleFunc("/project/ProjectOrca", func(w http.ResponseWriter, r *http.Request) {
		data := TemplateData{
			Title:       "Project Orca - Deepesh Kalura",
			CurrentPath: "/project/ProjectOrca",
			Data:        nil,
		}
		renderProjectTemplate(w, r, "project-orca", data)
	})

	http.HandleFunc("/project/ProjectOrca/baka_deepesh_news", func(w http.ResponseWriter, r *http.Request) {
		data := TemplateData{
			Title:       "Project Orca News",
			CurrentPath: "/project/ProjectOrca/baka_deepesh_news",
			Data:        nil,
		}
		renderProjectTemplate(w, r, "project-orca-news", data)

	})
}
