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

func (r *Repository) FindAll(ctx context.Context) ([]Book, error) {
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
		ORDER BY id DESC
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := make([]Book, 0)

	for rows.Next() {
		var book Book

		if err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.ISBN,
			&book.Description,
			&book.PublishedYear,
			&book.TotalCopies,
			&book.AvailableCopies,
			&book.CreatedAt,
			&book.UpdatedAt,
		); err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}