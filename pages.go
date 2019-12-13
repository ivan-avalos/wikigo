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
