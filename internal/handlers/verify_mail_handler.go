package handlers

import "github.com/jufabeck2202/piScraper/internal/core/ports"

type VerifMailHandler struct {
	mailService ports.MailService
}

func NewVerifMailHandler(mailService ports.MailService) *VerifMailHandler {
	return &VerifMailHandler{
		mailService: mailService,
	}
}
