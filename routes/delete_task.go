package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/jufabeck2202/piScraper/messaging"
	"github.com/jufabeck2202/piScraper/messaging/types"
)

type DeleteTask struct {
	Recipient   types.Recipient `json:"recipient" validate:"dive,required"`
	Destination types.Platform  `json:"destination" validate:"required"`
	Capcha      string          `json:"captcha" validate:"required"`
}

var alertManager = messaging.NewAlerts()

func DeleteTaskController(c *fiber.Ctx) error {

	deleteTask := &DeleteTask{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(deleteTask); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	errors := Validate(*deleteTask)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}
	err := Capcha.Verify(deleteTask.Capcha)
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
	numberOfDeletedNotifications := alertManager.DeleteTask(Websites.GetAllUrls(), deleteTask.Recipient, deleteTask.Destination)
	log.Println("Removed Notification for ", deleteTask.Recipient)
	return c.JSON(numberOfDeletedNotifications)
}
