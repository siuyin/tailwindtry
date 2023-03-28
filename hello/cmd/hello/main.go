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

	"github.com/nats-io/nuid"
	"github.com/siuyin/dflt"
)

//go:embed all:static all:tmpl
var content embed.FS

type Button struct {
	ID      string
	Text    string
	Color   string
	BG      string
	OnClick template.JS
}

func main() {
	//tmpl := template.Must(template.ParseGlob("./tmpl/*.html"))
	tmpl := template.Must(template.ParseFS(content, "tmpl/*.html"))

	rootHandler("/", tmpl)
	aboutHandler("/about", tmpl)
	contactHandler("/contact", tmpl)

	robotstxt("/robots.txt")

	apiV1Github("/api/v1/github/")
	apiV1NUID("/api/v1/nuid")

	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.Handle("/static/", http.FileServer(http.FS(content)))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func rootHandler(mnt string, tmpl *template.Template) {
	http.HandleFunc(mnt, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			w.WriteHeader(404)
			return
		}

		tmpl.ExecuteTemplate(w, "main", map[string]interface{}{
			"title":   "Gerbau",
			"content": func() string { return "brown fox" }(),
			"btn1":    Button{"btn1", "Lazy Dog", "text-gray-800", "bg-gray-100", template.JS(`document.getElementById("status").innerHTML="<p class='text-emerald-800'>Dog responds with a lazy woof!</p>"`)},
			"btn2":    Button{"btn2", "Clear", "text-gray-800", "bg-gray-100", template.JS(`document.getElementById("status").innerHTML=""`)},
			"mod1":    Button{"mod1", "Modal Demo", "text-gray-800", "bg-gray-100", template.JS(`document.getElementById("modal1").classList.remove("hidden")`)},
			"modal1": struct {
				ID  string
				Btn Button
			}{"modal1", Button{"modClose", "Close", "text-gray-800", "bg-gray-100", template.JS(`document.getElementById("modal1").classList.add("hidden")`)}},
			"incrBtn": Button{"incrBtn", "+1", "text-gray-800", "bg-gray-100", template.JS(``)},
			"decrBtn": Button{"decrBtn", "-1", "text-gray-800", "bg-gray-100", template.JS(``)},
		})
	})
}

func robotstxt(mnt string) {
	http.HandleFunc(mnt, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `User-agent: *
		Disallow: /login/`)
	})
}

func aboutHandler(mnt string, tmpl *template.Template) {
	http.HandleFunc(mnt, func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "about", map[string]string{
			"title": "About Us",
		})
	})
}

func contactHandler(mnt string, tmpl *template.Template) {
	http.HandleFunc(mnt, func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "contact", map[string]any{
			"title": "Contact",
			"list":  []string{"phone1", "phone2", "phone3"},
		})
	})
}

func apiV1Github(mnt string) {
	http.HandleFunc(mnt, func(w http.ResponseWriter, r *http.Request) {
		p := strings.Split(html.EscapeString(r.URL.Path), "/")
		usr := p[len(p)-1]
		client := &http.Client{}
		req, err := http.NewRequest("GET", fmt.Sprintf("https://api.github.com/users/%s", usr), nil)
		if err != nil {
			log.Println(err)
			return
		}

		req.Header.Add("Accept", "application/vnd.github+json")
		token := dflt.EnvString("GITHUB_TOKEN", "none")
		if token != "none" {
			req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
		}
		resp, err := client.Do(req)
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
}

func apiV1NUID(mnt string) {
	http.HandleFunc(mnt, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", nuid.Next())
	})
}
