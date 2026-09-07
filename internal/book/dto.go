package book

type CreateBookRequest struct {
	Title          string `json:"title"`
	ISBN           string `json:"isbn"`
	Description    string `json:"description"`
	PublishedYear  int    `json:"published_year"`
	TotalCopies    int    `json:"total_copies"`
}