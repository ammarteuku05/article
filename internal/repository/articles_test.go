package repository

import (
	"articles/internal/entity"
	"articles/shared/dto"
	"context"
	"database/sql/driver"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestInsertArticle(t *testing.T) {
	// Create sqlmock DB
	db, mock, repo := SetupTestArticle(t)
	defer db.Close()

	// Prepare test data
	article := &entity.Article{
		ID:        "article-123",
		AuthorID:  "author-456",
		Title:     "Test Title",
		Body:      "Test Body",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Expect Exec to be called with correct args
	mock.ExpectExec(`INSERT INTO articles \(id, author_id, title, body, created_at, updated_at\)`).
		WithArgs(article.ID, article.AuthorID, article.Title, article.Body, article.CreatedAt, article.UpdatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1)) // last insert id, rows affected

	// Call repository.the function
	err := repo.InsertArticle(context.Background(), article)
	require.NoError(t, err)

	// Ensure all expectations were met
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAllArticles(t *testing.T) {
	// Create mock DB
	db, mock, repo := SetupTestArticle(t)
	defer db.Close()

	// Prepare test query parameters
	queryParams := &dto.QueryGetArticle{
		AuthorName: "John",
		Titles:     "Golang",
		Body:       "test body",
		Sort:       "DESC",
		PerPage:    2,
		Page:       0, // no offset param
	}

	// Arguments for count query (filters only)
	countArgs := []driver.Value{
		"%John%",
		"%Golang%",
		"%test body%",
	}

	// Arguments for main query (filters + limit)
	mainArgs := []driver.Value{
		"%John%",
		"%Golang%",
		"%test body%",
		2, // limit only, no offset since Page=0
	}

	// Mock count query
	mock.ExpectQuery(`SELECT COUNT\(\*\) FROM \(`).
		WithArgs(countArgs...).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(2))

	// Mock main query
	rows := sqlmock.NewRows([]string{
		"article_id", "author_id", "title", "body", "author_name", "created_at",
	}).AddRow(
		"1", "10", "Golang Rocks", "This is the body", "John Doe", time.Now(),
	).AddRow(
		"2", "10", "Advanced Go", "Another body", "John Doe", time.Now(),
	)

	mock.ExpectQuery(`SELECT\s+ar\.id AS article_id`).
		WithArgs(mainArgs...).
		WillReturnRows(rows)

	// Execute
	ctx := context.Background()
	result, total, err := repo.GetAllArticles(ctx, queryParams)

	require.NoError(t, err)
	require.Equal(t, 2, total)
	require.Len(t, *result, 2)
	require.Equal(t, "Golang Rocks", (*result)[0].Title)
	require.Equal(t, "Advanced Go", (*result)[1].Title)

	require.NoError(t, mock.ExpectationsWereMet())
}
