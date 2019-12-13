package main

import (
	"log"
	"net/http"
	"regexp"

	"github.com/dannyvankooten/extemplate"
)

// Global variables
var xt *extemplate.Extemplate
var validPath = regexp.MustCompile("^/(edit|save|view|delete)/([a-zA-Z0-9]+)$")

// Network handlers

func homeHandler(w http.ResponseWriter, r *http.Request) {
	pages, err := loadPages()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	renderTemplate(w, "home", pages)
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func deleteHandler(w http.ResponseWriter, r *http.Request, title string) {
	err := deletePage(title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func main() {
	// Parse templates
	xt = extemplate.New()
	err := xt.ParseDir("tmpl/", []string{".tmpl"})
	if err != nil {
		log.Fatal(err.Error())
	}

	// Handle HTTP requests
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	http.HandleFunc("/delete/", makeHandler(deleteHandler))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
