package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/jufabeck2202/piScraper/internal/core/domain"
	"github.com/jufabeck2202/piScraper/internal/core/ports"
)

type AddTasks struct {
	Tasks  []domain.AlertTask `json:"tasks" validate:"dive,required"`
	Capcha string             `json:"captcha" validate:"required"`
}

type AlertHandler struct {
	websiteService   ports.WebsiteService
	validatorService ports.ValidateService
	captchaService   ports.CaptchaService
	alertService     ports.AlertService
}

func NewAlertHandler(websiteService ports.WebsiteService, validatorService ports.ValidateService, captchaService ports.CaptchaService, alertService ports.AlertService) *AlertHandler {

	return &AlertHandler{
		websiteService:   websiteService,
		validatorService: validatorService,
		captchaService:   captchaService,
		alertService:     alertService,
	}
}

func (hdl *AlertHandler) Post(c *fiber.Ctx) error {
	// Create new Book struct
	addTasks := &AddTasks{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(addTasks); err != nil {
		log.Println("error: ", err)
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	errors := hdl.validatorService.Validate(*addTasks)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}
	err := hdl.captchaService.Verify(addTasks.Capcha)
	if err != nil {
		log.Println("error: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	//check if the task is valid
	for _, t := range addTasks.Tasks {
		if t.Recipient.Pushover == "" && t.Recipient.Webhook == "" && t.Recipient.Email == "" || t.Destination > 3 || t.Destination < 1 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"msg":   "invalid task structure",
			})
		}
	}
	//Add task to alert
	for _, t := range addTasks.Tasks {
		hdl.alertService.AddAlert(t.Website.URL, domain.Alert{Recipient: t.Recipient.SanitizedRecipient(), Destination: t.Destination})
	}
	log.Println("Added new Notification")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "task added",
	})

}
