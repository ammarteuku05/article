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

//go:generate mockery --name ArticlesService --case snake --output ./mocks --disable-version-string
type (
	// articlesService is
	ArticlesService interface {
		InsertArticle(ctx context.Context, req *dto.RequestArticle) error
		GetAllArticle(ctx context.Context, query *dto.QueryGetArticle) (*[]dto.ResponseGetArticles, int, error)
		UpdateArticle(ctx context.Context, articleId string, req *dto.RequestArticle) error
	}

	articlesImpl struct {
		repo repository.Holder
		deps shared.Deps
	}
)

// NewArticlesService is
func NewArticlesService(repo repository.Holder, deps shared.Deps) ArticlesService {
	return &articlesImpl{
		repo: repo,
		deps: deps,
	}
}

func (s *articlesImpl) InsertArticle(ctx context.Context, req *dto.RequestArticle) error {
	_, _, err := s.repo.AuthorRepository.GetByIdAuthor(ctx, req.AuthorID)
	if err != nil {
		s.deps.Logger.Debugf("error get by id author : %v", err.Error())
		return err
	}

	return s.repo.ArticlesRepository.InsertArticle(ctx, &entity.Article{
		ID:        uuid.NewString(),
		AuthorID:  req.AuthorID,
		Body:      req.Body,
		Title:     req.Title,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
}

func (s *articlesImpl) GetAllArticle(ctx context.Context, query *dto.QueryGetArticle) (*[]dto.ResponseGetArticles, int, error) {
	return s.repo.ArticlesRepository.GetAllArticles(ctx, query)
}

func (s *articlesImpl) UpdateArticle(ctx context.Context, articleId string, req *dto.RequestArticle) error {
	article, err := s.repo.GetByIdArticle(ctx, articleId)
	if err != nil {
		s.deps.Logger.Debugf("error get by id article : %v", err.Error())
		return err
	}

	return s.repo.UpdateArticle(ctx, &entity.Article{
		ID:       article.ID,
		AuthorID: req.AuthorID,
		Title:    req.Title,
		Body:     req.Body,
	})
}
