package author

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

func (h *Handler) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var request CreateAuthorRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	author, err := h.service.CreateAuthor(r.Context(), request)
	if err != nil {
		if errors.Is(err, ErrInvalidName) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		http.Error(w, "Failed to create author", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(author)
}

func (h *Handler) GetAuthors(w http.ResponseWriter, r *http.Request) {
	authors, err := h.service.GetAuthors(r.Context())
	if err != nil {
		http.Error(w, "Failed to retrieve authors", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(authors)
}

func (h *Handler) GetAuthorByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid author ID", http.StatusBadRequest)
		return
	}

	author, err := h.service.GetAuthorByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Author not found", http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to retrieve author", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(author)
}

func (h *Handler) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid author ID", http.StatusBadRequest)
		return
	}

	var request CreateAuthorRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	author, err := h.service.UpdateAuthor(r.Context(), id, request)
	if err != nil {
		if errors.Is(err, ErrInvalidName) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Author not found", http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to update author", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(author)
}

func (h *Handler) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid author ID", http.StatusBadRequest)
		return
	}

	err = h.service.DeleteAuthor(r.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Author not found", http.StatusNotFound)
			return
		}

		http.Error(w, "Failed to delete author", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}