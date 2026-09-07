package category

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

func (h *Handler) CreateCategory(
	w http.ResponseWriter,
	r *http.Request,
) {
	var request CreateCategoryRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(
			w,
			"Invalid request body",
			http.StatusBadRequest,
		)
		return
	}

	category, err := h.service.CreateCategory(
		r.Context(),
		request,
	)
	if err != nil {
		if errors.Is(err, ErrInvalidName) {
			http.Error(
				w,
				err.Error(),
				http.StatusBadRequest,
			)
			return
		}

		http.Error(
			w,
			"Failed to create category",
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(category)
}

func (h *Handler) GetCategories(
	w http.ResponseWriter,
	r *http.Request,
) {
	categories, err := h.service.GetCategories(
		r.Context(),
	)
	if err != nil {
		http.Error(
			w,
			"Failed to retrieve categories",
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(categories)
}

func (h *Handler) GetCategoryByID(
	w http.ResponseWriter,
	r *http.Request,
) {
	idStr := r.PathValue("id")

	id, err := strconv.ParseInt(
		idStr,
		10,
		64,
	)
	if err != nil {
		http.Error(
			w,
			"Invalid category ID",
			http.StatusBadRequest,
		)
		return
	}

	category, err := h.service.GetCategoryByID(
		r.Context(),
		id,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(
				w,
				"Category not found",
				http.StatusNotFound,
			)
			return
		}

		http.Error(
			w,
			"Failed to retrieve category",
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(category)
}

func (h *Handler) UpdateCategory(
	w http.ResponseWriter,
	r *http.Request,
) {
	idStr := r.PathValue("id")

	id, err := strconv.ParseInt(
		idStr,
		10,
		64,
	)
	if err != nil {
		http.Error(
			w,
			"Invalid category ID",
			http.StatusBadRequest,
		)
		return
	}

	var request CreateCategoryRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(
			w,
			"Invalid request body",
			http.StatusBadRequest,
		)
		return
	}

	category, err := h.service.UpdateCategory(
		r.Context(),
		id,
		request,
	)
	if err != nil {
		if errors.Is(err, ErrInvalidName) {
			http.Error(
				w,
				err.Error(),
				http.StatusBadRequest,
			)
			return
		}

		if errors.Is(err, sql.ErrNoRows) {
			http.Error(
				w,
				"Category not found",
				http.StatusNotFound,
			)
			return
		}

		http.Error(
			w,
			"Failed to update category",
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(category)
}

func (h *Handler) DeleteCategory(
	w http.ResponseWriter,
	r *http.Request,
) {
	idStr := r.PathValue("id")

	id, err := strconv.ParseInt(
		idStr,
		10,
		64,
	)
	if err != nil {
		http.Error(
			w,
			"Invalid category ID",
			http.StatusBadRequest,
		)
		return
	}

	err = h.service.DeleteCategory(
		r.Context(),
		id,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(
				w,
				"Category not found",
				http.StatusNotFound,
			)
			return
		}

		http.Error(
			w,
			"Failed to delete category",
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}