package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/jufabeck2202/piScraper/internal/core/ports"
	"github.com/jufabeck2202/piScraper/messaging/types"
)

type DeleteHandler struct {
	websiteService   ports.WebsiteService
	validatorService ports.ValidateService
	captchaService   ports.CaptchaService
	alertService     ports.AlertService
}

type DeleteTask struct {
	Recipient   types.Recipient `json:"recipient" validate:"dive,required"`
	Destination types.Platform  `json:"destination" validate:"required"`
	Capcha      string          `json:"captcha" validate:"required"`
}

func NewDeleteHandler(websiteService ports.WebsiteService, validatorService ports.ValidateService, captchaService ports.CaptchaService, alertService ports.AlertService) *DeleteHandler {

	return &DeleteHandler{
		websiteService:   websiteService,
		validatorService: validatorService,
		captchaService:   captchaService,
		alertService:     alertService,
	}
}

func (hdl *DeleteHandler) Delete(c *fiber.Ctx) error {

	deleteTask := &DeleteTask{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(deleteTask); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	errors := hdl.validatorService.Validate(*deleteTask)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}
	err := hdl.captchaService.Verify(deleteTask.Capcha)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}
	//check if the task is valid
	if deleteTask.Recipient.Pushover == "" && deleteTask.Recipient.Webhook == "" && deleteTask.Recipient.Email == "" || deleteTask.Destination > 3 || deleteTask.Destination < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   "invalid task structure",
		})
	}
	numberOfDeletedNotifications := hdl.alertService.DeleteTask(hdl.websiteService.GetAllUrls(), deleteTask.Recipient, deleteTask.Destination)
	log.Println("Removed Notification for ", deleteTask.Recipient)
	return c.JSON(numberOfDeletedNotifications)
}
