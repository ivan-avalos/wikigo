package main

import (
	"io/ioutil"
	"os"
	"regexp"
)

type Page struct {
	Title string
	Body  []byte
}

type Pages struct {
	Pages []Page
}

func (p *Page) save() error {
	filename := "data/" + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := "data/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func loadPages() (*Pages, error) {
	files, err := ioutil.ReadDir("./data")
	if err != nil {
		return nil, err
	}
	pages := Pages{Pages: make([]Page, 0)}
	for _, f := range files {
		name := regexp.MustCompile("^([a-zA-Z0-9]+).txt$").FindStringSubmatch(f.Name())[1]
		page, err := loadPage(name)
		if err != nil {
			return nil, err
		}
		pages.Pages = append(pages.Pages, *page)
	}
	return &pages, nil
}

func deletePage(title string) error {
	return os.Remove("data/" + title + ".txt")
}
