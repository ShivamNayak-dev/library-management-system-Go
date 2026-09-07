package category

import (
	"context"
	"errors"
	"strings"
)

var ErrInvalidName = errors.New("category name is required")

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) CreateCategory(
	ctx context.Context,
	request CreateCategoryRequest,
) (*Category, error) {
	if strings.TrimSpace(request.Name) == "" {
		return nil, ErrInvalidName
	}

	category := &Category{
		Name:        strings.TrimSpace(request.Name),
		Description: strings.TrimSpace(request.Description),
	}

	if err := s.repository.Create(ctx, category); err != nil {
		return nil, err
	}

	return category, nil
}

func (s *Service) GetCategories(
	ctx context.Context,
) ([]Category, error) {
	return s.repository.FindAll(ctx)
}

func (s *Service) GetCategoryByID(
	ctx context.Context,
	id int64,
) (*Category, error) {
	return s.repository.FindByID(ctx, id)
}

func (s *Service) UpdateCategory(
	ctx context.Context,
	id int64,
	request CreateCategoryRequest,
) (*Category, error) {
	if strings.TrimSpace(request.Name) == "" {
		return nil, ErrInvalidName
	}

	category := &Category{
		ID:          id,
		Name:        strings.TrimSpace(request.Name),
		Description: strings.TrimSpace(request.Description),
	}

	if err := s.repository.Update(ctx, id, category); err != nil {
		return nil, err
	}

	return category, nil
}

func (s *Service) DeleteCategory(
	ctx context.Context,
	id int64,
) error {
	return s.repository.Delete(ctx, id)
}