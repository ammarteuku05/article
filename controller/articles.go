package controller

import (
	"articles/internal/service"
	"articles/shared"
	"articles/shared/dto"
	"articles/shared/errors"

	"github.com/labstack/echo"
)

type (
	//ArticlesController is
	ArticlesController struct {
		services service.Holder
		deps     shared.Deps
	}
)

// NewArticlesController is
func NewArticlesController(services service.Holder, deps shared.Deps) (*ArticlesController, error) {
	return &ArticlesController{
		services: services,
		deps:     deps,
	}, nil
}

// CreateArticles is
func (ctrl *ArticlesController) CreateArticles(ctx echo.Context) error {
	var (
		pctx    = shared.NewEmptyContext(ctx)
		context = pctx.Request().Context()
		request = new(dto.RequestArticle)
	)

	if err := ctx.Bind(request); err != nil {
		return pctx.Fail(errors.ErrBindingRequest(err.Error()))
	}

	if err := ctrl.deps.CustomValidator.Validate(request); err != nil {
		return pctx.Fail(errors.ErrValidationRequest(err.Error()))
	}

	err := ctrl.services.ArticleService.InsertArticle(context, request)
	if err != nil {
		return pctx.Fail(err)
	}

	return pctx.Success(nil)
}

// GetAllArticles is
func (ctrl *ArticlesController) GetAllArticles(ctx echo.Context) error {
	var (
		pctx    = shared.NewEmptyContext(ctx)
		context = pctx.Request().Context()
		request = new(dto.QueryGetArticle)
	)
	pages := ctrl.deps.NewPaginationFromRequest(ctx.Request())

	if err := ctx.Bind(request); err != nil {
		return pctx.Fail(errors.ErrBindingRequest(err.Error()))
	}

	res, totalData, err := ctrl.services.ArticleService.GetAllArticle(context, request)
	if err != nil {
		return pctx.Fail(err)
	}
	pages.SetData(res, totalData)

	return pctx.Success(pages)
}

// UpdateArticle is
func (ctrl *ArticlesController) UpdateArticle(ctx echo.Context) error {
	var (
		pctx    = shared.NewEmptyContext(ctx)
		context = pctx.Request().Context()
		request = new(dto.RequestArticle)
	)

	if err := ctx.Bind(request); err != nil {
		return pctx.Fail(errors.ErrBindingRequest(err.Error()))
	}

	if err := ctrl.deps.CustomValidator.Validate(request); err != nil {
		return pctx.Fail(errors.ErrValidationRequest(err.Error()))
	}

	err := ctrl.services.ArticleService.UpdateArticle(context, ctx.Param("article_id"), request)
	if err != nil {
		return pctx.Fail(err)
	}

	return pctx.Success(nil)
}
