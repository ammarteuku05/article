package service

import (
	"articles/internal/entity"
	"articles/shared/dto"
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestIInsertAuthor_Success(t *testing.T) {
	req := &dto.RequestAuthor{
		Name: "name athor",
	}

	mockAuthor.On("InsertAuthor", mock.Anything, mock.AnythingOfType("*entity.Author")).
		Return(nil).Once()

	err := author.InsertAuthor(context.Background(), req)
	assert.NoError(t, err)
	mockAuthor.AssertExpectations(t)
}

func TestGetAllByName(t *testing.T) {
	expected := []entity.Author{
		{ID: uuid.NewString(), Name: "Name Author"},
	}

	mockAuthor.On("GetAllByName", mock.Anything, mock.Anything).
		Return(&expected, nil).Once()

	res, err := author.GetAllByName(context.Background(), "name_author")

	assert.NoError(t, err)
	assert.Equal(t, "Name Author", (*res)[0].Name)
	mockArticles.AssertExpectations(t)
}
