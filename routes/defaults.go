package routes

import (
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"gopkg.in/ezzarghili/recaptcha-go.v4"

	"github.com/jufabeck2202/piScraper/messaging"
	"github.com/jufabeck2202/piScraper/messaging/types"
	"github.com/jufabeck2202/piScraper/utils"
)

var Websites = utils.Websites{}
var Capcha, _ = recaptcha.NewReCAPTCHA(os.Getenv("RECAPTCHA_SECRET"), recaptcha.V3, 10*time.Second) // for v3 API use https://g.co/recaptcha/v3 (apperently the same admin UI at the time of writing)
var AlertManager = messaging.NewAlerts()

func Validate(input interface{}) []*types.ErrorResponse {
	var errors []*types.ErrorResponse
	err := validate.Struct(input)
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
