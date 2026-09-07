package book

import (
	"context"
	"errors"
)

var (
	ErrInvalidTitle         = errors.New("title is required")
	ErrInvalidISBN          = errors.New("ISBN is required")
	ErrInvalidTotalCopies   = errors.New("total copies must be greater than zero")
	ErrInvalidPublishedYear = errors.New("published year is invalid")
)

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) CreateBook(ctx context.Context, request CreateBookRequest) (*Book, error) {
	if request.Title == "" {
		return nil, ErrInvalidTitle
	}

	if request.ISBN == "" {
		return nil, ErrInvalidISBN
	}

	if request.TotalCopies <= 0 {
		return nil, ErrInvalidTotalCopies
	}

	book := &Book{
		Title:           request.Title,
		ISBN:            request.ISBN,
		Description:     request.Description,
		PublishedYear:   request.PublishedYear,
		TotalCopies:     request.TotalCopies,
		AvailableCopies: request.TotalCopies,
	}

	if err := s.repository.Create(ctx, book); err != nil {
		return nil, err
	}

	return book, nil
}