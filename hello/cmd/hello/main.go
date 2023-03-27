package main

import (
	"embed"
	"fmt"
	"html"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//go:embed all:static all:tmpl
var content embed.FS

func main() {
	//tmpl := template.Must(template.ParseGlob("./tmpl/*.html"))
	tmpl := template.Must(template.ParseFS(content, "tmpl/*.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "main", map[string]interface{}{
			"title":   "Gerbau",
			"content": func() string { return "brown fox" }(),
			"btn1": struct {
				ID      string
				Text    string
				Color   string
				BG      string
				OnClick template.JS
			}{"btn1", "Lazy Dog", "text-gray-800", "bg-gray-100", template.JS(`document.getElementById("status").innerHTML="You clicked the button"`)},
			"btn2": struct {
				ID      string
				Text    string
				Color   string
				BG      string
				OnClick template.JS
			}{"btn2", "Clear", "text-gray-800", "bg-gray-100", template.JS(`document.getElementById("status").innerHTML=""`)},
		})
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "about", map[string]string{"title": "About Us"})
	})

	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "contact", map[string]string{"title": "Contact"})
	})

	http.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `User-agent: *
		Disallow: /login/`)
	})

	http.HandleFunc("/api/v1/github/", func(w http.ResponseWriter, r *http.Request) {
		p := strings.Split(html.EscapeString(r.URL.Path), "/")
		usr := p[len(p)-1]
		resp, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s", usr))
		if err != nil {
			log.Println(err)
			return
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			return
		}
		w.Write(body)
	})

	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.Handle("/static/", http.FileServer(http.FS(content)))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
