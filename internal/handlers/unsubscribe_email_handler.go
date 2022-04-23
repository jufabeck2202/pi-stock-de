package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/jufabeck2202/piScraper/internal/core/ports"
)

type UnsubscribeEmailHandler struct {
	alertService ports.AlertService
	emailService ports.MailService
}

func NewUnsubscribeMailHandler(alertService ports.AlertService, emailService ports.MailService) *UnsubscribeEmailHandler {
	return &UnsubscribeEmailHandler{
		alertService: alertService,
		emailService: emailService,
	}
}

func (hdl *UnsubscribeEmailHandler) Get(c *fiber.Ctx) error {
	fmt.Println(c.Params("email"))
	email := c.Params("email")
	decytedEmail := hdl.emailService.Decrypt(email)
	hdl.alertService.RemoveEmailAlert(decytedEmail)
	err := hdl.emailService.Delete(decytedEmail)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"error": false,
		"msg":   decytedEmail,
	})
}
