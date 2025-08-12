package shared

import (
	"github.com/go-playground/validator/v10"
)

type (
	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func NewCustomValidator() (*CustomValidator, error) {
	customValidator := validator.New()

	return &CustomValidator{validator: customValidator}, nil
}
