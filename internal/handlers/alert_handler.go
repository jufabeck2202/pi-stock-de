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
	emailService     ports.MailService
}

func NewAlertHandler(websiteService ports.WebsiteService, validatorService ports.ValidateService, captchaService ports.CaptchaService, alertService ports.AlertService, emailService ports.MailService) *AlertHandler {

	return &AlertHandler{
		websiteService:   websiteService,
		validatorService: validatorService,
		captchaService:   captchaService,
		alertService:     alertService,
		emailService:     emailService,
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
	//check if email should be verified
	containsEmail := false
	emails := make([]string, 0)
	for _, t := range addTasks.Tasks {
		if t.Recipient.Email != "" && t.Destination == domain.Mail {
			containsEmail = true
			emails = append(emails, t.Recipient.Email)
		}
	}
	//email needs to be verified
	if containsEmail {

		//check if all emails are the same
		if !sameEmails(emails) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"msg":   "all emails need to be the same",
			})
		}

		//check if email is verified
		if !hdl.emailService.IsVerified(emails[0]) {
			//email needs to be verified
			log.Println("error: ", "email needs to be verified")
			hdl.emailService.NewEmailSubscriber(emails[0])
		}

	}

	//Add task to alert
	for _, t := range addTasks.Tasks {
		hdl.alertService.AddAlert(t.Website.URL, domain.Alert{Recipient: t.Recipient.SanitizedRecipient(), Destination: t.Destination})
	}
	log.Println("Added new Notification")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "task added, will be notified, if email is verified",
	})

}

func sameEmails(a []string) bool {
	for i := 1; i < len(a); i++ {
		if a[i] != a[0] {
			return false
		}
	}
	return true
}
