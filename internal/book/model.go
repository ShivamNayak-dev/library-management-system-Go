package book

import "time"

type Author struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Book struct {
	ID              int64      `json:"id"`
	Title           string     `json:"title"`
	ISBN            string     `json:"isbn"`
	Description     string     `json:"description"`
	PublishedYear   int        `json:"published_year"`
	TotalCopies     int        `json:"total_copies"`
	AvailableCopies int        `json:"available_copies"`
	Authors         []Author   `json:"authors"`
	Categories      []Category `json:"categories"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}