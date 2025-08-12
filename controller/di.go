package controller

import (
	"articles/shared"
	"articles/shared/errors"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.uber.org/dig"
)

type (
	Holder struct {
		dig.In

		Deps shared.Deps

		AuthorController  *AuthorsController
		ArticleController *ArticlesController
	}
)

func Register(container *dig.Container) error {
	if err := container.Provide(NewAuthorController); err != nil {
		return err
	}

	if err := container.Provide(NewArticlesController); err != nil {
		return err
	}

	return nil
}

func (h *Holder) SetupRoutes(app *echo.Echo) {

	app.Validator = h.Deps.CustomValidator
	app.HTTPErrorHandler = h.ErrorHandler

	app.Use(middleware.Recover())
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	v1 := app.Group("/v1")

	authorRoute := v1.Group("/authors")
	authorRoute.POST("/create", h.AuthorController.CreateAuthor)
	authorRoute.GET("/get-all", h.AuthorController.GetAllByName)

	articlesRoute := v1.Group("/articles")
	articlesRoute.POST("/create", h.ArticleController.CreateArticles)
	articlesRoute.PUT("/update/:article_id", h.ArticleController.UpdateArticle)
	articlesRoute.GET("/articles", h.ArticleController.GetAllArticles)
}

func (h *Holder) ErrorHandler(err error, ctx echo.Context) {
	var (
		sctx, ok = ctx.(*shared.ArticlesContext)
	)

	if !ok {
		sctx = shared.NewEmptyContext(ctx)
	}

	e, ok := err.(*echo.HTTPError)
	if ok {
		msg, ok := e.Message.(string)
		if !ok {
			msg = err.Error()
		}
		err = errors.ErrBase.New(msg).WithProperty(errors.ErrCodeProperty, e.Code).WithProperty(errors.ErrHttpCodeProperty, e.Code)
	}

	h.Deps.Logger.Errorf(
		"path=%s error=%s",
		sctx.Path(),
		err,
	)

	_ = sctx.Fail(err)
}
