package book

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
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

func (h *Handler) GetBookByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	book, err := h.service.GetBookByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Book not found", http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to retrieve book", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(book)
}

func (h *Handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	var request CreateBookRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	book, err := h.service.UpdateBook(r.Context(), id, request)
	if err != nil {
		switch {
		case errors.Is(err, ErrInvalidTitle),
			errors.Is(err, ErrInvalidISBN),
			errors.Is(err, ErrInvalidTotalCopies):
			http.Error(w, err.Error(), http.StatusBadRequest)

		case errors.Is(err, sql.ErrNoRows):
			http.Error(w, "Book not found", http.StatusNotFound)

		default:
			http.Error(w, "Failed to update book", http.StatusInternalServerError)
		}

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(book)
}

func (h *Handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	err = h.service.DeleteBook(r.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Book not found", http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to delete book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) AddAuthor(w http.ResponseWriter, r *http.Request) {
	bookIDStr := r.PathValue("id")

	bookID, err := strconv.ParseInt(bookIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	var request AddAuthorRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if request.AuthorID <= 0 {
		http.Error(w, "Invalid author ID", http.StatusBadRequest)
		return
	}

	if err := h.service.AddAuthor(
		r.Context(),
		bookID,
		request.AuthorID,
	); err != nil {
		http.Error(
			w,
			"Failed to add author to book",
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (r *Repository) findCategoriesByBookID(
	ctx context.Context,
	bookID int64,
) ([]Category, error) {
	query := `
		SELECT
			c.id,
			c.name
		FROM categories c
		INNER JOIN book_categories bc
			ON c.id = bc.category_id
		WHERE bc.book_id = ?
		ORDER BY c.id
	`

	rows, err := r.db.QueryContext(ctx, query, bookID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := make([]Category, 0)

	for rows.Next() {
		var category Category

		if err := rows.Scan(
			&category.ID,
			&category.Name,
		); err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}