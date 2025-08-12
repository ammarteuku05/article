package repository

import (
	"articles/internal/entity"
	"articles/shared/errors"
	"context"
	"database/sql"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestInsertAuthor(t *testing.T) {
	db, mock, repo := SetupTestAuthor(t)
	defer db.Close()

	author := &entity.Author{
		ID:        "123",
		Name:      "John Doe",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mock.ExpectExec(regexp.QuoteMeta(`
		INSERT INTO authors (id, name, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
	`)).
		WithArgs(author.ID, author.Name, author.CreatedAt, author.UpdatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.InsertAuthor(context.Background(), author)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetByIdAuthor_Found(t *testing.T) {
	db, mock, repo := SetupTestAuthor(t)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow("123", "John Doe", time.Now(), time.Now())

	mock.ExpectQuery(regexp.QuoteMeta(`
		SELECT id, name, created_at, updated_at
		FROM authors
		WHERE id = $1
	`)).
		WithArgs("123").
		WillReturnRows(rows)

	author, found, err := repo.GetByIdAuthor(context.Background(), "123")
	assert.NoError(t, err)
	assert.True(t, found)
	assert.Equal(t, "John Doe", author.Name)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetByIdAuthor_NotFound(t *testing.T) {
	db, mock, repo := SetupTestAuthor(t)
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(`
		SELECT id, name, created_at, updated_at
		FROM authors
		WHERE id = $1
	`)).
		WithArgs("999").
		WillReturnError(sql.ErrNoRows)

	author, found, err := repo.GetByIdAuthor(context.Background(), "999")
	assert.ErrorIs(t, err, errors.ErrRecordNotFoundAuthor)
	assert.False(t, found)
	assert.Nil(t, author)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAllByName_WithName(t *testing.T) {
	db, mock, repo := SetupTestAuthor(t)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow("123", "John Doe", time.Now(), time.Now()).
		AddRow("124", "Johnny", time.Now(), time.Now())

	mock.ExpectQuery(regexp.QuoteMeta(`
		SELECT id, name, created_at, updated_at
		FROM authors WHERE name ILIKE $1
	`)).
		WithArgs("%John%").
		WillReturnRows(rows)

	authors, err := repo.GetAllByName(context.Background(), "John")
	assert.NoError(t, err)
	assert.Len(t, *authors, 2)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAllByName_EmptyName(t *testing.T) {
	db, mock, repo := SetupTestAuthor(t)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).
		AddRow("125", "Alice", time.Now(), time.Now())

	mock.ExpectQuery(regexp.QuoteMeta(`
		SELECT id, name, created_at, updated_at
		FROM authors
	`)).
		WillReturnRows(rows)

	authors, err := repo.GetAllByName(context.Background(), "")
	assert.NoError(t, err)
	assert.Len(t, *authors, 1)
	assert.NoError(t, mock.ExpectationsWereMet())
}
