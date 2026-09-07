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


func (r *Repository) FindByID(ctx context.Context, id int64) (*Book, error) {
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

	var book Book

	err := r.db.QueryRowContext(ctx, query, id).Scan(
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

	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (r *Repository) Update(ctx context.Context, id int64, book *Book) error {
	query := `
		UPDATE books
		SET
			title = ?,
			isbn = ?,
			description = ?,
			published_year = ?,
			total_copies = ?,
			available_copies = ?,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = ?
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
		id,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	book.ID = id

	return r.findByID(ctx, book)
}

func (r *Repository) Delete(ctx context.Context, id int64) error {
	query := `
		DELETE FROM books
		WHERE id = ?
	`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}