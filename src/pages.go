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
	"github.com/jinzhu/gorm"
)

// Page represents a single article
type Page struct {
	gorm.Model
	Title string
	Body  []byte `gorm:"type:text"`
}

// Pages represents a collection of articles
type Pages struct {
	Pages []Page
}

func (p *Page) save() error {
	var page Page
	if err := db.First(&page, "title = ?", p.Title).Error; err != nil {
		// Does not exist yet… create!
		if gorm.IsRecordNotFoundError(err) {
			if err := db.Create(p).Error; err != nil {
				return err
			}
			return nil
		}
		return err
	}

	// Already exists… modify!
	if err := db.Model(&page).Update("body", p.Body).Error; err != nil {
		return err
	}
	return nil
}

func loadPage(title string) (*Page, error) {
	var page Page
	if err := db.First(&page, "title = ?", title).Error; err != nil {
		return nil, err
	}
	return &page, nil
}

func loadPages() (*Pages, error) {
	var pages []Page
	if err := db.Find(&pages).Error; err != nil {
		return nil, err
	}
	return &Pages{Pages: pages}, nil
}

func deletePage(title string) error {
	var page Page
	if err := db.Find(&page, "title = ?", title).Error; err != nil {
		return err
	}
	if err := db.Delete(&page).Error; err != nil {
		return err
	}
	return nil
}
