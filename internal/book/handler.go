package book

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var request CreateBookRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	book, err := h.service.CreateBook(r.Context(), request)
	if err != nil {
		switch {
		case errors.Is(err, ErrInvalidTitle),
			errors.Is(err, ErrInvalidISBN),
			errors.Is(err, ErrInvalidTotalCopies):
			http.Error(w, err.Error(), http.StatusBadRequest)

		default:
			http.Error(w, "Failed to create book", http.StatusInternalServerError)
		}

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(book)
}

func (h *Handler) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.service.GetBooks(r.Context())
	if err != nil {
		http.Error(w, "Failed to retrieve books", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(books)
}