package book

import (
	"context"
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(ctx context.Context, book *Book) error {
	query := `
		INSERT INTO books (
			title,
			isbn,
			description,
			published_year,
			total_copies,
			available_copies
		)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	result, err := r.db.ExecContext(
		ctx,
		query,
		book.Title,
		book.ISBN,
		book.Description,
		book.PublishedYear,
		book.TotalCopies,
		book.AvailableCopies,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	book.ID = id

	return r.findByID(ctx, book)
}

func (r *Repository) findByID(ctx context.Context, book *Book) error {
	query := `
		SELECT
			id,
			title,
			isbn,
			description,
			published_year,
			total_copies,
			available_copies,
			created_at,
			updated_at
		FROM books
		WHERE id = ?
	`

	return r.db.QueryRowContext(ctx, query, book.ID).Scan(
		&book.ID,
		&book.Title,
		&book.ISBN,
		&book.Description,
		&book.PublishedYear,
		&book.TotalCopies,
		&book.AvailableCopies,
		&book.CreatedAt,
		&book.UpdatedAt,
	)
}