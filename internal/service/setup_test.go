package service

import (
	"articles/internal/repository"
	"articles/internal/repository/mocks"
	mocks_logger "articles/logger/mocks"
	"articles/shared"
)

var (
	mockAuthor   = new(mocks.AuthorRepository)
	mockArticles = new(mocks.ArticlesRepository)
	mockLogger   = new(mocks_logger.Logger)

	article = &articlesImpl{
		repo: repository.Holder{
			AuthorRepository:   mockAuthor,
			ArticlesRepository: mockArticles,
		},
		deps: shared.Deps{
			Logger: mockLogger,
		},
	}

	author = &authorImpl{
		repo: repository.Holder{
			AuthorRepository:   mockAuthor,
			ArticlesRepository: mockArticles,
		},
		deps: shared.Deps{
			Logger: mockLogger,
		},
	}
)
