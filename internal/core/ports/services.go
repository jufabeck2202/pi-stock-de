package ports

import (
	"github.com/jufabeck2202/piScraper/internal/core/domain"
)

type WebsiteService interface {
	GetList() []domain.Website
	GetAllUrls() []string
	GetUrls(shop string) []string
	GetItemById(url string) domain.Website
	UpdateItemInList(item domain.Website)
	Init()
	Load()
	Save()
}

type ValidateService interface {
	Validate(input interface{}) []*domain.ErrorResponse
}

type CaptchaService interface {
	Verify(captcha string) error
}

type AlertService interface {
	LoadAlerts(url string) []domain.Alert
	AddAlert(url string, alert domain.Alert)
	SaveAlerts(url string, alerts []domain.Alert)
	DeleteTask(urls []string, recipient domain.Recipient, platform domain.Platform) int
}

type MailService interface {
	IsVerified(email string) bool
	NewEmailSubscriber(email string) error
	Verify(email string) error
}

type MessagingPlatform interface {
	Send(recipient domain.Recipient, item domain.Website) error
}

type Adaptor interface {
	Run(list domain.Websites)
	Wait()
}
