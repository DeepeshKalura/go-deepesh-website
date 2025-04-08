# Deepesh's Go Website

A personal website built with Go, HTMX, and Tailwind CSS.

## Getting Started

1. Clone this repository
2. Run the application:
   ```
   go run cmd/server/main.go
   ```
3. Open your browser at http://localhost:8080

## Project Structure

- `/cmd/server`: Contains the main application entry point
- `/internal`: Application code
- `/static`: Static assets (CSS, JS, images)
- `/templates`: HTML templates
- `/assets`: Source files for assets (like SCSS, etc.)

## Technologies Used

- Go for backend
- HTMX for interactive UI without JavaScript
- Tailwind CSS for styling


for the deployment

```bash
go build -tags netgo -ldflags '-s -w' -o app ./cmd/server/
```
