package di

import (
	"articles/controller"
	"articles/internal/repository"
	"articles/internal/service"
	shared "articles/shared"
	"articles/shared/config"
	"fmt"

	"go.uber.org/dig"
)

var (
	Container *dig.Container = dig.New()
)

func init() {
	if err := Container.Provide(config.New); err != nil {
		panic(err)
	}

	if err := Container.Provide(NewDB); err != nil {
		fmt.Printf("err %v", err)
		panic(err)
	}

	if err := shared.Register(Container); err != nil {
		panic(err)
	}

	if err := Container.Provide(NewLogger); err != nil {
		panic(err)
	}

	if err := service.Register(Container); err != nil {
		panic(err)
	}

	if err := repository.Register(Container); err != nil {
		panic(err)
	}

	if err := controller.Register(Container); err != nil {
		panic(err)
	}
}
