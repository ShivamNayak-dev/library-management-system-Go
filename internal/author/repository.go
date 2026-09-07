package author

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

func (r *Repository) Create(ctx context.Context, author *Author) error {
	query := `
		INSERT INTO authors (
			name,
			biography
		)
		VALUES (?, ?)
	`

	result, err := r.db.ExecContext(
		ctx,
		query,
		author.Name,
		author.Biography,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	author.ID = id

	return r.findByID(ctx, author)
}

func (r *Repository) findByID(ctx context.Context, author *Author) error {
	query := `
		SELECT
			id,
			name,
			biography,
			created_at,
			updated_at
		FROM authors
		WHERE id = ?
	`

	return r.db.QueryRowContext(ctx, query, author.ID).Scan(
		&author.ID,
		&author.Name,
		&author.Biography,
		&author.CreatedAt,
		&author.UpdatedAt,
	)
}

func (r *Repository) FindAll(ctx context.Context) ([]Author, error) {
	query := `
		SELECT
			id,
			name,
			biography,
			created_at,
			updated_at
		FROM authors
		ORDER BY id DESC
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	authors := make([]Author, 0)

	for rows.Next() {
		var author Author

		if err := rows.Scan(
			&author.ID,
			&author.Name,
			&author.Biography,
			&author.CreatedAt,
			&author.UpdatedAt,
		); err != nil {
			return nil, err
		}

		authors = append(authors, author)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return authors, nil
}

func (r *Repository) FindByID(ctx context.Context, id int64) (*Author, error) {
	query := `
		SELECT
			id,
			name,
			biography,
			created_at,
			updated_at
		FROM authors
		WHERE id = ?
	`

	var author Author

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&author.ID,
		&author.Name,
		&author.Biography,
		&author.CreatedAt,
		&author.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &author, nil
}

func (r *Repository) Update(ctx context.Context, id int64, author *Author) error {
	query := `
		UPDATE authors
		SET
			name = ?,
			biography = ?,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = ?
	`

	result, err := r.db.ExecContext(
		ctx,
		query,
		author.Name,
		author.Biography,
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

	author.ID = id

	return r.findByID(ctx, author)
}

func (r *Repository) Delete(ctx context.Context, id int64) error {
	query := `
		DELETE FROM authors
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