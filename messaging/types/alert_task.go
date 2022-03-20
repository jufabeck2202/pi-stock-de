package types

import "github.com/jufabeck2202/piScraper/utils"

type AlertTask struct {
	Website     utils.Website `json:"website"`
	Recipient   Recipient     `json:"recipient"`
	Destination Platform      `json:"destination"`
}
