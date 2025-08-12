package repository

import "go.uber.org/dig"

type (
	//Holder is
	Holder struct {
		dig.In

		AuthorRepository
		ArticlesRepository
	}
)

// Register is
func Register(container *dig.Container) error {
	if err := container.Provide(NewAuthorRepository); err != nil {
		return err
	}

	if err := container.Provide(NewArticlesRepository); err != nil {
		return err
	}

	return nil
}
