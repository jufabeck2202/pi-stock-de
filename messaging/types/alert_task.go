package types

import "github.com/jufabeck2202/piScraper/utils"

type AlertTask struct {
	Website     utils.Website `json:"website" validate:"dive,required"`
	Recipient   Recipient     `json:"recipient" validate:"dive,required"`
	Destination Platform      `json:"destination" validate:"required,min=1"`
}
