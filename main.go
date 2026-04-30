package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/smtp"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"

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

// RateLimiter stores IP access times
type RateLimiter struct {
	sync.Mutex
	ips map[string][]time.Time
}

var limiter = RateLimiter{
	ips: make(map[string][]time.Time),
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
			httpPort = "80"
		}
		log.Printf("Starting HTTP redirector on port %s...", httpPort)
		if err := http.ListenAndServe(":"+httpPort, certManager.HTTPHandler(nil)); err != nil {
			log.Printf("HTTP redirector failed: %v", err)
		}
	}()

	httpsPort := os.Getenv("PORT")
	if httpsPort == "" {
		httpsPort = "443"
	}

	server := &http.Server{
		Addr:         ":" + httpsPort,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
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

	// 1. Path Traversal Protection
	// Only allow alphanumeric characters, dashes, and underscores
	validParam := regexp.MustCompile(`^[a-zA-Z0-9\-_]+$`)
	if !validParam.MatchString(title) || !validParam.MatchString(date) {
		http.Error(w, "Invalid parameters", http.StatusBadRequest)
		return
	}

	fileName := date + "-" + title + ".md"
	filePath := filepath.Join("data/posts", fileName)

	// Additional check: Ensure the resulting path is still within data/posts
	absBase, _ := filepath.Abs("data/posts")
	absTarget, err := filepath.Abs(filePath)
	if err != nil || !strings.HasPrefix(absTarget, absBase) {
		http.Error(w, "Access denied", http.StatusForbidden)
		return
	}

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

	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	limiter.Lock()
	now := time.Now()
	var recent []time.Time
	for _, t := range limiter.ips[ip] {
		if now.Sub(t) < time.Hour {
			recent = append(recent, t)
		}
	}
	if len(recent) >= 3 {
		limiter.Unlock()
		http.Error(w, "Rate limit exceeded. Please try again in an hour.", http.StatusTooManyRequests)
		return
	}
	limiter.ips[ip] = append(recent, now)
	limiter.Unlock()

	var form ContactForm
	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// 2. SMTP Header Injection Protection
	// Strip any newline characters from user-provided email used in headers
	cleanEmail := strings.NewReplacer("\r", "", "\n", "").Replace(form.Email)

	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(cleanEmail) {
		http.Error(w, "Invalid email address", http.StatusBadRequest)
		return
	}

	if len(form.Message) < 10 || len(form.Message) > 1000 {
		http.Error(w, "Message must be between 10 and 1000 characters", http.StatusBadRequest)
		return
	}

	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASS")
	toEmail := os.Getenv("CONTACT_TO")

	if smtpHost == "" || smtpUser == "" || smtpPass == "" {
		log.Printf("SMTP missing. Message from %s: %s", cleanEmail, form.Message)
		w.WriteHeader(http.StatusOK)
		return
	}

	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)
	msg := []byte(fmt.Sprintf("From: %s\r\n"+
		"To: %s\r\n"+
		"Subject: Portfolio Contact from %s\r\n"+
		"Content-Type: text/plain; charset=UTF-8\r\n"+
		"\r\n"+
		"You received a new message from your portfolio website.\n\n"+
		"Visitor Email: %s\n"+
		"Message:\n%s\r\n", smtpUser, toEmail, cleanEmail, cleanEmail, form.Message))

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, smtpUser, []string{toEmail}, msg)
	if err != nil {
		log.Printf("Failed to send email: %v", err)
		http.Error(w, "Failed to send message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
