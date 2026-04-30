package main

import (
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"os"

	"golang.org/x/crypto/acme/autocert"
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
	mux := http.NewServeMux()

	// API routes
	mux.HandleFunc("/api/projects", projectsHandler)
	mux.HandleFunc("/api/blog", blogHandler)

	// Static files service
	staticPath := "./static"
	if _, err := os.Stat(staticPath); os.IsNotExist(err) {
		log.Printf("Static directory '%s' not found, API only mode", staticPath)
	} else {
		fs := http.FileServer(http.Dir(staticPath))
		mux.Handle("/", http.StripPrefix("/", fs))
	}

	// HTTPS Configuration with autocert
	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("wzlin.top", "www.wzlin.top"),
		Cache:      autocert.DirCache("certs"),
	}

	// Redirect HTTP to HTTPS
	go func() {
		httpPort := os.Getenv("HTTP_PORT")
		if httpPort == "" {
			httpPort = "80" // Standard HTTP port
		}
		log.Printf("Starting HTTP redirector on port %s...", httpPort)
		// autocert.Manager.HTTPHandler returns a handler that handles ACME challenges
		// and redirects all other traffic to HTTPS.
		if err := http.ListenAndServe(":"+httpPort, certManager.HTTPHandler(nil)); err != nil {
			log.Printf("HTTP redirector failed: %v", err)
		}
	}()

	httpsPort := os.Getenv("PORT")
	if httpsPort == "" {
		httpsPort = "443" // Standard HTTPS port
	}

	server := &http.Server{
		Addr:    ":" + httpsPort,
		Handler: mux,
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
			MinVersion:     tls.VersionTLS12,
		},
	}

	log.Printf("Starting HTTPS server on port %s...", httpsPort)
	if err := server.ListenAndServeTLS("", ""); err != nil {
		log.Fatalf("HTTPS server failed to start: %v", err)
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
