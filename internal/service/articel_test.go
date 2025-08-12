package service

import (
	"articles/internal/entity"
	"articles/shared/dto"
	"context"
	"errors"

	"testing"

	errors_shared "articles/shared/errors"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInsertArticle_Success(t *testing.T) {
	req := &dto.RequestArticle{
		AuthorID: "auth-1",
		Title:    "My Title",
		Body:     "My Body",
	}

	mockAuthor.On("GetByIdAuthor", mock.Anything, "auth-1").
		Return(&entity.Author{ID: "auth-1"}, true, nil).Once()
	mockArticles.On("InsertArticle", mock.Anything, mock.AnythingOfType("*entity.Article")).
		Return(nil).Once()

	err := article.InsertArticle(context.Background(), req)
	assert.NoError(t, err)
	mockAuthor.AssertExpectations(t)
	mockArticles.AssertExpectations(t)
}

func TestInsertArticle_AuthorNotFound(t *testing.T) {

	req := &dto.RequestArticle{
		AuthorID: "auth-2",
		Title:    "Bad Author",
		Body:     "Nope",
	}

	mockAuthor.On("GetByIdAuthor", mock.Anything, "auth-2").
		Return(nil, false, errors.New("not found")).Once()
	mockLogger.On("Debugf", mock.Anything, mock.Anything).Once()

	err := article.InsertArticle(context.Background(), req)
	assert.Error(t, err)
	mockAuthor.AssertExpectations(t)
	mockLogger.AssertExpectations(t)
}

func TestGetAllArticle(t *testing.T) {
	query := &dto.QueryGetArticle{}
	expected := []dto.ResponseGetArticles{
		{ArticleID: uuid.NewString(), Title: "Title", Body: "Body", AuthorName: "Author"},
	}

	mockArticles.On("GetAllArticles", mock.Anything, query).
		Return(&expected, 1, nil).Once()

	res, count, err := article.GetAllArticle(context.Background(), query)
	assert.NoError(t, err)
	assert.Equal(t, 1, count)
	assert.Equal(t, "Title", (*res)[0].Title)
	mockArticles.AssertExpectations(t)
}

func TestUpdateArticle_Success(t *testing.T) {
	req := &dto.RequestArticle{
		AuthorID: "auth-1",
		Title:    "Updated Title",
		Body:     "Updated Body",
	}

	mockArticles.On("GetByIdArticle", mock.Anything, "art-1").
		Return(&entity.Article{
			ID: "art-1",
		}, nil).Once()
	mockArticles.On("UpdateArticle", mock.Anything, mock.Anything).Return(nil).Once()

	err := article.UpdateArticle(context.Background(), "art-1", req)
	assert.NoError(t, err)
}

func TestUpdateArticle_NotFound(t *testing.T) {
	mockArticles.On("GetByIdArticle", mock.Anything, "art-1").
		Return(nil, errors_shared.ErrRecordNotFoundArticle).Once()
	mockLogger.On("Debugf", mock.Anything, mock.Anything).Once()

	req := &dto.RequestArticle{
		AuthorID: "auth-1",
		Title:    "Updated Title",
		Body:     "Updated Body",
	}

	err := article.UpdateArticle(context.Background(), "art-1", req)
	assert.Error(t, err)
	mockLogger.AssertExpectations(t)
}
