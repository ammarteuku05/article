package service

import (
	"articles/internal/entity"
	"articles/internal/repository"
	"articles/shared"
	"articles/shared/dto"
	"context"
	"time"

	"github.com/google/uuid"
)

//go:generate mockery --name AuthorsService --case snake --output ./mocks --disable-version-string
type (
	// AuthorsService is
	AuthorsService interface {
		InsertAuthor(ctx context.Context, req *dto.RequestAuthor) error
		GetAllByName(ctx context.Context, name string) (*[]entity.Author, error)
	}

	authorImpl struct {
		repo repository.Holder
		deps shared.Deps
	}
)

// NewAuthorService is
func NewAuthorService(repo repository.Holder, deps shared.Deps) AuthorsService {
	return &authorImpl{
		repo: repo,
		deps: deps,
	}
}

func (s *authorImpl) InsertAuthor(ctx context.Context, req *dto.RequestAuthor) error {
	return s.repo.AuthorRepository.InsertAuthor(ctx, &entity.Author{
		ID:        uuid.NewString(),
		Name:      req.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
}

func (s *authorImpl) GetAllByName(ctx context.Context, name string) (*[]entity.Author, error) {
	return s.repo.AuthorRepository.GetAllByName(ctx, name)
}
