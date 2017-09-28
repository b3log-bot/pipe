// Solo.go - A small and beautiful blogging platform written in golang.
// Copyright (C) 2017, b3log.org
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package model

// Tag model.
type Tag struct {
	Model

	Title                 string `gorm:"size:128" json:"title"`
	ArticleCount          int    `json:"articleCount"`          // including drafts and published articles
	PublishedArticleCount int    `json:"publishedArticleCount"` // just including published articles

	BlogID uint `json:"blogID"`
}