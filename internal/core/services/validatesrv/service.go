package validatesrv

import (
	"github.com/go-playground/validator/v10"

	"github.com/jufabeck2202/piScraper/internal/core/domain"
)

type service struct {
	validator *validator.Validate
}

func New() *service {

	return &service{
		validator: validator.New(),
	}
}

func (srv *service) Validate(input interface{}) []*domain.ErrorResponse {
	var errors []*domain.ErrorResponse
	err := srv.validator.Struct(input)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element domain.ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
