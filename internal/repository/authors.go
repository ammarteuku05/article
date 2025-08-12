package repository

import (
	"articles/internal/entity"
	"articles/shared"
	"articles/shared/errors"
	"context"
	"database/sql"
)

//go:generate mockery --name AuthorRepository --case snake --output ./mocks --disable-version-string
type (
	AuthorRepository interface {
		InsertAuthor(ctx context.Context, req *entity.Author) error
		GetByIdAuthor(ctx context.Context, id string) (*entity.Author, bool, error)
		GetAllByName(ctx context.Context, name string) (*[]entity.Author, error)
	}

	authorImpl struct {
		shared.Deps
	}
)

// NewAuthorRepository is
func NewAuthorRepository(deps shared.Deps) AuthorRepository {
	return &authorImpl{Deps: deps}
}

func (s *authorImpl) InsertAuthor(ctx context.Context, req *entity.Author) error {
	query := `
		INSERT INTO authors (id, name, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
	`
	_, err := s.DB.Exec(query, req.ID, req.Name, req.CreatedAt, req.UpdatedAt)
	return err
}

func (s *authorImpl) GetByIdAuthor(ctx context.Context, id string) (*entity.Author, bool, error) {
	var author entity.Author

	query := `
		SELECT id, name, created_at, updated_at
		FROM authors
		WHERE id = $1
	`

	err := s.DB.QueryRowContext(ctx, query, id).Scan(
		&author.ID,
		&author.Name,
		&author.CreatedAt,
		&author.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, false, errors.ErrRecordNotFoundAuthor // not found
		}
		return nil, false, err // other error
	}

	return &author, true, nil
}

func (s *authorImpl) GetAllByName(ctx context.Context, name string) (*[]entity.Author, error) {
	authors := make([]entity.Author, 0)

	baseQuery := `
		SELECT id, name, created_at, updated_at
		FROM authors
	`
	var rows *sql.Rows
	var err error

	if name != "" {
		query := baseQuery + ` WHERE name ILIKE $1`
		rows, err = s.DB.QueryContext(ctx, query, "%"+name+"%")
	} else {
		rows, err = s.DB.QueryContext(ctx, baseQuery)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var author entity.Author
		if err := rows.Scan(
			&author.ID,
			&author.Name,
			&author.CreatedAt,
			&author.UpdatedAt,
		); err != nil {
			return nil, err
		}
		authors = append(authors, author)
	}

	return &authors, nil
}
