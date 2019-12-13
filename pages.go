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
	"database/sql"
	"log"
)

// Page represents a single article
type Page struct {
	ID    int
	Title string
	Body  []byte
}

// Pages represents a collection of articles
type Pages struct {
	Pages []Page
}

func (p *Page) save() error {
	// Perform search
	sel, err := db.Prepare("SELECT title FROM pages WHERE title = ?")
	if err != nil {
		return err
	}
	defer sel.Close()
	var title string
	err = sel.QueryRow(p.Title).Scan(&title)
	if err != nil {
		log.Print("at least got here with", err.Error())
		if err == sql.ErrNoRows {
			// Does not exist
			ins, err := db.Prepare("INSERT INTO pages (title, body) VALUES(?, ?)")
			if err != nil {
				return err
			}
			defer ins.Close()
			_, err = ins.Exec(p.Title, p.Body)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}

	// Already exists
	upd, err := db.Prepare("UPDATE pages SET body = ? WHERE title = ?")
	if err != nil {
		return err
	}
	defer upd.Close()
	_, err = upd.Exec(p.Body, p.Title)
	if err != nil {
		return err
	}

	return nil
}

func loadPage(title string) (*Page, error) {
	sel, err := db.Prepare("SELECT * FROM pages WHERE title = ?")
	if err != nil {
		return nil, err
	}
	defer sel.Close()

	page := &Page{}
	err = sel.QueryRow(title).Scan(&page.ID, &page.Title, &page.Body)
	if err != nil {
		return nil, err
	}
	return page, nil
}

func loadPages() (*Pages, error) {
	sel, err := db.Prepare("SELECT * FROM pages")
	if err != nil {
		return nil, err
	}
	defer sel.Close()

	pages := Pages{Pages: make([]Page, 0)}
	var rows *sql.Rows
	rows, err = sel.Query()
	defer rows.Close()
	for rows.Next() {
		page := &Page{}
		err = rows.Scan(&page.ID, &page.Title, &page.Body)
		if err != nil {
			return nil, err
		}
		pages.Pages = append(pages.Pages, *page)
	}
	return &pages, nil
}

func deletePage(title string) error {
	del, err := db.Prepare("DELETE FROM pages WHERE title = ?")
	if err != nil {
		return err
	}
	defer del.Close()

	_, err = del.Exec(title)
	if err != nil {
		return err
	}

	return nil
}
