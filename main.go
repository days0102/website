package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

// Project represents a portfolio project
type Project struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
}

// Blog represents a blog post summary
type Blog struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Date    string `json:"date"`
}

func main() {
	// API routes
	http.HandleFunc("/api/projects", projectsHandler)
	http.HandleFunc("/api/blog", blogHandler)

	// Static files service
	staticPath := "./static"
	if _, err := os.Stat(staticPath); os.IsNotExist(err) {
		log.Printf("Static directory '%s' not found, API only mode", staticPath)
	} else {
		fs := http.FileServer(http.Dir(staticPath))
		http.Handle("/", http.StripPrefix("/", fs))
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s...", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func projectsHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("data/projects.json")
	if err != nil {
		http.Error(w, "Could not open projects data", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Type", "application/json")
	io.Copy(w, file)
}

func blogHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("data/blogs.json")
	if err != nil {
		http.Error(w, "Could not open blog data", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Type", "application/json")
	io.Copy(w, file)
}
