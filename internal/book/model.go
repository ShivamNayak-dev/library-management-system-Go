package book

import "time"

type Book struct {
	ID             int64     `json:"id"`
	Title          string    `json:"title"`
	ISBN           string    `json:"isbn"`
	Description    string    `json:"description"`
	PublishedYear  int       `json:"published_year"`
	TotalCopies    int       `json:"total_copies"`
	AvailableCopies int      `json:"available_copies"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}