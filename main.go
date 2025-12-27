package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"ascii_art_web/helper"
)

type PageData struct {
	Text    string
	Banner  string
	Content string
	Error   string
}

var tmpl *template.Template

func main() {
	var err error

	tmpl, err = template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal("Error loading template:", err)
	}

	http.HandleFunc("/", homeHandler)

	http.HandleFunc("/ascii-art", asciiArtHandler)

	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 page not found", 404)
		return
	}

	if r.Method == "GET" {
		data := PageData{}
		renderTemplate(w, data, http.StatusOK)
		return
	} else {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "400 Bad Request: "+err.Error(), http.StatusBadRequest)
		return
	}

	data := PageData{}
	text := r.FormValue("text")

	banner := r.FormValue("banner")

	if text == "" || banner == "" {
		data.Error = "Please enter some text and a banner"
		fmt.Println("no text or banner")
		renderTemplate(w, data, http.StatusBadRequest)
		return
	}
	// check if the files is bigger than 15k
	if len(text) > 15*1024 {
		data.Error = "Please enter a text lower than 15kb"
		renderTemplate(w, data, http.StatusBadRequest)
		return
	}

	// Call helper function
	result, err := helper.MainHelper(text, banner)
	if err != nil {
		log.Printf("Error generating ASCII art: %v", err)
		data.Error = "Failed to generate ASCII art: " + err.Error()
		renderTemplate(w, data, http.StatusBadRequest)
		return
	}

	// Success
	data.Text = "\n" + text
	data.Banner = banner
	data.Content = "\n" + result + "\n"

	renderTemplate(w, data, http.StatusOK)
}

func renderTemplate(w http.ResponseWriter, data PageData, status int) {
	var buf bytes.Buffer

	err := tmpl.Execute(&buf, data)
	if err != nil {
		log.Printf("Template error: %v", err)
		// Only now send 500 if template failed
		http.Error(w, "500 Internal Server Error: error rendering template please conact devs", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	buf.WriteTo(w)
}
