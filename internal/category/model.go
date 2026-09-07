package category

import "time"

type Book struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	ISBN  string `json:"isbn"`
}

type Category struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Books       []Book    `json:"books"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}