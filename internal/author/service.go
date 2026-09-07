package author

import (
	"context"
	"errors"
	"strings"
)

var ErrInvalidName = errors.New("author name is required")

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) CreateAuthor(
	ctx context.Context,
	request CreateAuthorRequest,
) (*Author, error) {
	if strings.TrimSpace(request.Name) == "" {
		return nil, ErrInvalidName
	}

	author := &Author{
		Name:      strings.TrimSpace(request.Name),
		Biography: strings.TrimSpace(request.Biography),
	}

	if err := s.repository.Create(ctx, author); err != nil {
		return nil, err
	}

	return author, nil
}

func (s *Service) GetAuthors(ctx context.Context) ([]Author, error) {
	return s.repository.FindAll(ctx)
}

func (s *Service) GetAuthorByID(
	ctx context.Context,
	id int64,
) (*Author, error) {
	return s.repository.FindByID(ctx, id)
}

func (s *Service) UpdateAuthor(
	ctx context.Context,
	id int64,
	request CreateAuthorRequest,
) (*Author, error) {
	if strings.TrimSpace(request.Name) == "" {
		return nil, ErrInvalidName
	}

	author := &Author{
		ID:        id,
		Name:      strings.TrimSpace(request.Name),
		Biography: strings.TrimSpace(request.Biography),
	}

	if err := s.repository.Update(ctx, id, author); err != nil {
		return nil, err
	}

	return author, nil
}

func (s *Service) DeleteAuthor(ctx context.Context, id int64) error {
	return s.repository.Delete(ctx, id)
}