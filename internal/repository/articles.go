package repository

import (
	"articles/internal/entity"
	"articles/shared"
	"articles/shared/dto"
	"articles/shared/errors"
	"context"
	"database/sql"
	"fmt"
	"strings"
)

//go:generate mockery --name ArticlesRepository --case snake --output ./mocks --disable-version-string
type (
	ArticlesRepository interface {
		InsertArticle(ctx context.Context, req *entity.Article) error
		GetAllArticles(ctx context.Context, query *dto.QueryGetArticle) (*[]dto.ResponseGetArticles, int, error)
		UpdateArticle(ctx context.Context, req *entity.Article) error
		GetByIdArticle(ctx context.Context, articleId string) (*entity.Article, error)
	}

	articleImpl struct {
		shared.Deps
	}
)

// NewArticlesRepository is
func NewArticlesRepository(deps shared.Deps) ArticlesRepository {
	return &articleImpl{Deps: deps}
}

func (s *articleImpl) InsertArticle(ctx context.Context, req *entity.Article) error {
	query := `
		INSERT INTO articles (id, author_id, title, body, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := s.DB.Exec(query, req.ID, req.AuthorID, req.Title, req.Body, req.CreatedAt, req.UpdatedAt)
	return err
}

func (s *articleImpl) GetAllArticles(ctx context.Context, query *dto.QueryGetArticle) (*[]dto.ResponseGetArticles, int, error) {
	res := make([]dto.ResponseGetArticles, 0)

	// Base query
	sqlQuery := `
		SELECT 
			ar.id AS article_id,
			au.id AS author_id,
			ar.title,
			ar.body,
			au.name AS author_name,
			ar.created_at AS created_at
		FROM articles ar
		JOIN authors au ON au.id = ar.author_id
		WHERE 1=1
	`

	// Parameters slice for placeholders
	args := []interface{}{}
	argIndex := 1 // For positional params ($1, $2...) in PostgreSQL. If MySQL, just use `?`

	// Filters
	if query.AuthorName != "" {
		sqlQuery += fmt.Sprintf(" AND au.name ILIKE $%d", argIndex)
		args = append(args, "%"+query.AuthorName+"%")
		argIndex++
	}

	if query.Titles != "" {
		sqlQuery += fmt.Sprintf(" AND ar.title ILIKE $%d", argIndex)
		args = append(args, "%"+query.Titles+"%")
		argIndex++
	}

	if query.Body != "" {
		sqlQuery += fmt.Sprintf(" AND ar.body ILIKE $%d", argIndex)
		args = append(args, "%"+query.Body+"%")
		argIndex++
	}

	sortOrder := "DESC" // default newest first
	if strings.ToUpper(query.Sort) == "ASC" {
		sortOrder = "ASC"
	}
	sqlQuery += fmt.Sprintf(" ORDER BY ar.created_at %s", sortOrder)

	// Count query for total rows
	countQuery := "SELECT COUNT(*) FROM (" + sqlQuery + ") AS count_table"
	var total int
	if err := s.DB.QueryRowContext(ctx, countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	// Pagination
	if query.PerPage > 0 {
		sqlQuery += fmt.Sprintf(" LIMIT $%d", argIndex)
		args = append(args, query.PerPage)
		argIndex++
	}
	if query.Page > 0 {
		sqlQuery += fmt.Sprintf(" OFFSET $%d", argIndex)
		args = append(args, query.Page)
		argIndex++
	}

	// Execute query
	rows, err := s.DB.QueryContext(ctx, sqlQuery, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var r dto.ResponseGetArticles
		if err := rows.Scan(
			&r.ArticleID,
			&r.AuthorID,
			&r.Title,
			&r.Body,
			&r.AuthorName,
			&r.CreatedAt,
		); err != nil {
			return nil, 0, err
		}
		res = append(res, r)
	}

	return &res, total, nil
}

func (s *articleImpl) UpdateArticle(ctx context.Context, req *entity.Article) error {
	query := `
		UPDATE articles
		SET author_id = $2,
			title = $3,
			body = $4,
			created_at = $5,
			updated_at = $6
		WHERE id = $1
	`
	res, err := s.DB.Exec(query,
		req.ID,
		req.AuthorID,
		req.Title,
		req.Body,
		req.CreatedAt,
		req.UpdatedAt,
	)
	if err != nil {
		return err
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("no article found with id %v", req.ID)
	}

	return nil
}

func (s *articleImpl) GetByIdArticle(ctx context.Context, articleId string) (*entity.Article, error) {
	article := new(entity.Article)

	query := `
		SELECT id, author_id, title, body, created_at, updated_at
		FROM articles
		WHERE id = $1
	`

	err := s.DB.QueryRowContext(ctx, query, articleId).Scan(
		&article.ID,
		&article.AuthorID,
		&article.Title,
		&article.Body,
		&article.CreatedAt,
		&article.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrRecordNotFoundArticle // not found
		}
		return nil, err // other error
	}

	return article, nil
}
