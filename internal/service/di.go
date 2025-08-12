package service

import "go.uber.org/dig"

type (
	Holder struct {
		dig.In
		AuthorService  AuthorsService
		ArticleService ArticlesService
	}
)

func Register(container *dig.Container) error {
	if err := container.Provide(NewAuthorService); err != nil {
		return err
	}

	if err := container.Provide(NewArticlesService); err != nil {
		return err
	}

	return nil
}
