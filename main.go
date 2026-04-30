package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"path/filepath"

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

// ContactForm represents the structure of the contact message
type ContactForm struct {
	Email   string `json:"email"`
	Message string `json:"message"`
}

func main() {
	mux := http.NewServeMux()

	// API routes
	mux.HandleFunc("/api/projects", projectsHandler)
	mux.HandleFunc("/api/blog", blogHandler)
	mux.HandleFunc("/api/blog/content", blogContentHandler)
	mux.HandleFunc("/api/contact", contactHandler)

	// Static files service
	staticPath := "./static"
	if _, err := os.Stat(staticPath); os.IsNotExist(err) {
		log.Printf("Static directory '%s' not found, API only mode", staticPath)
	} else {
		fs := http.FileServer(http.Dir(staticPath))
		mux.Handle("/", fs)
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
	files, err := os.ReadDir("data/posts")
	if err != nil {
		http.Error(w, "Could not read blog posts", http.StatusInternalServerError)
		return
	}

	var blogs []Blog
	for i, f := range files {
		if f.IsDir() || filepath.Ext(f.Name()) != ".md" {
			continue
		}

		name := f.Name()
		title := name[:len(name)-3]
		date := "2026-04-30"
		if len(name) > 10 {
			date = name[:10]
			title = name[11 : len(name)-3]
		}

		blogs = append(blogs, Blog{
			ID:      i + 1,
			Title:   title,
			Summary: "Click to read the full article...",
			Date:    date,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blogs)
}

func blogContentHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	date := r.URL.Query().Get("date")
	if title == "" || date == "" {
		http.Error(w, "Missing title or date", http.StatusBadRequest)
		return
	}

	fileName := date + "-" + title + ".md"
	filePath := filepath.Join("data/posts", fileName)

	content, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, "Blog post not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/markdown")
	w.Write(content)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var form ContactForm
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Email configuration from environment variables
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASS")
	toEmail := os.Getenv("CONTACT_TO")

	if smtpHost == "" || smtpUser == "" || smtpPass == "" {
		log.Printf("SMTP configuration missing. Message from %s: %s", form.Email, form.Message)
		// Return success to user but log it locally if SMTP is not configured
		w.WriteHeader(http.StatusOK)
		return
	}

	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)
	msg := []byte(fmt.Sprintf("To: %s\r\n"+
		"Subject: New Contact Message from %s\r\n"+
		"\r\n"+
		"From: %s\n\n%s\r\n", toEmail, form.Email, form.Email, form.Message))

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, smtpUser, []string{toEmail}, msg)
	if err != nil {
		log.Printf("Failed to send email: %v", err)
		http.Error(w, "Failed to send message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
