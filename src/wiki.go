/**
	WikiGo!
    Copyright (C) 2019  Iván Ávalos <ivan.avalos.diaz@hotmail.com>

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/
package main

import (
	"log"
	"net/http"
	"regexp"

	"github.com/dannyvankooten/extemplate"
	"github.com/jinzhu/gorm"
)

// Global variables
var db *gorm.DB
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
	// Initialize MySQL DB
	dbs := &DBSettings{
		Username: "wikigo",
		Password: "wikigo",
		Database: "wikigo",
		Host:     "localhost",
		Port:     "3306",
	}
	var err error
	db, err = dbs.initDb()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	db.AutoMigrate(&Page{})

	// Parse templates
	xt = extemplate.New()
	err = xt.ParseDir("tmpl/", []string{".tmpl"})
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
