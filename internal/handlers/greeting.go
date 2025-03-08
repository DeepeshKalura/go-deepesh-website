package handlers

import (
	"net/http"
)

// GreetingHandler responds with a greeting message
func GreetingHandler(w http.ResponseWriter, r *http.Request) {
	// You would use template rendering in a real app
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
}
