package controller

import (
	"articles/internal/service"
	"articles/shared"
	"articles/shared/dto"
	"articles/shared/errors"

	"github.com/labstack/echo"
)

type (
	//AuthorsController is
	AuthorsController struct {
		services service.Holder
		deps     shared.Deps
	}
)

// NewAuthorController is
func NewAuthorController(services service.Holder, deps shared.Deps) (*AuthorsController, error) {
	return &AuthorsController{
		services: services,
		deps:     deps,
	}, nil
}

// CreateAuthor is
func (ctrl *AuthorsController) CreateAuthor(ctx echo.Context) error {
	var (
		pctx    = shared.NewEmptyContext(ctx)
		context = pctx.Request().Context()
		request = new(dto.RequestAuthor)
	)

	if err := ctx.Bind(request); err != nil {
		return pctx.Fail(errors.ErrBindingRequest(err.Error()))
	}
	if err := ctx.Validate(request); err != nil {
		return pctx.Fail(errors.ErrValidationRequest(err.Error()))
	}

	err := ctrl.services.AuthorService.InsertAuthor(context, request)
	if err != nil {
		return pctx.Fail(err)
	}

	return pctx.Success(nil)
}
func (ctrl *AuthorsController) GetAllByName(ctx echo.Context) error {
	var (
		pctx    = shared.NewEmptyContext(ctx)
		context = pctx.Request().Context()
	)

	res, err := ctrl.services.AuthorService.GetAllByName(context, ctx.QueryParam("name"))
	if err != nil {
		return pctx.Fail(err)
	}

	return pctx.Success(res)
}
