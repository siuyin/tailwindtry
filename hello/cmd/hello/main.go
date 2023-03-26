package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//go:embed all:static all:tmpl
var content embed.FS

func main() {
	//tmpl := template.Must(template.ParseGlob("./tmpl/*.html"))
	tmpl := template.Must(template.ParseFS(content, "tmpl/*.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "main", map[string]string{
			"title":   "My Title",
			"content": "The quick brown fox"})
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "about", map[string]string{"title": "About Us"})
	})

	http.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `User-agent: *
		Disallow: /login/`)
	})

	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.Handle("/static/", http.FileServer(http.FS(content)))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
