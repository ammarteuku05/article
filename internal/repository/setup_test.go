package repository

import (
	"articles/shared"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func SetupTestArticle(t *testing.T) (*sql.DB, sqlmock.Sqlmock, *articleImpl) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	repo := &articleImpl{shared.Deps{
		DB: db,
	}}

	return db, mock, repo
}

func SetupTestAuthor(t *testing.T) (*sql.DB, sqlmock.Sqlmock, *authorImpl) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	repo := &authorImpl{shared.Deps{
		DB: db,
	}}

	return db, mock, repo
}
