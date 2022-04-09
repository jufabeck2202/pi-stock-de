package types

import "strings"

type Recipient struct {
	Webhook  string `json:"webhook" validate:"omitempty,max=255,min=1"`
	Pushover string `json:"pushover" validate:"omitempty,max=35,min=30"`
	Email    string `json:"email" validate:"omitempty,email,max=255"`
}

func (r Recipient) SanitizedRecipient() Recipient {
	return Recipient{
		Webhook:  strings.ToLower(r.Webhook),
		Pushover: r.Pushover,
		Email:    strings.ToLower(r.Email),
	}
}
