package book

import (
	"context"
	"errors"
	"strings"
)

var (
	ErrInvalidTitle       = errors.New("title is required")
	ErrInvalidISBN        = errors.New("ISBN is required")
	ErrInvalidTotalCopies = errors.New("total copies must be greater than zero")
)

func validateCreateBookRequest(request CreateBookRequest) error {
	if strings.TrimSpace(request.Title) == "" {
		return ErrInvalidTitle
	}

	if strings.TrimSpace(request.ISBN) == "" {
		return ErrInvalidISBN
	}

	if request.TotalCopies <= 0 {
		return ErrInvalidTotalCopies
	}

	return nil
}

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) CreateBook(ctx context.Context, request CreateBookRequest) (*Book, error) {
	if err := validateCreateBookRequest(request); err != nil {
		return nil, err
	}

	book := &Book{
		Title:           strings.TrimSpace(request.Title),
		ISBN:            strings.TrimSpace(request.ISBN),
		Description:     strings.TrimSpace(request.Description),
		PublishedYear:   request.PublishedYear,
		TotalCopies:     request.TotalCopies,
		AvailableCopies: request.TotalCopies,
	}

	if err := s.repository.Create(ctx, book); err != nil {
		return nil, err
	}

	return book, nil
}

func (s *Service) GetBooks(ctx context.Context) ([]Book, error) {
	return s.repository.FindAll(ctx)
}