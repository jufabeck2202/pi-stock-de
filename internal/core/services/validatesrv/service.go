package validatesrv

import (
	"github.com/go-playground/validator/v10"

	"github.com/jufabeck2202/piScraper/messaging/types"
)

type service struct {
	validator *validator.Validate
}

func New() *service {

	return &service{
		validator: validator.New(),
	}
}

func (srv *service) Validate(input interface{}) []*types.ErrorResponse {
	var errors []*types.ErrorResponse
	err := srv.validator.Struct(input)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element types.ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
