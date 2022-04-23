package domain

type AlertTask struct {
	Website     Website   `json:"website" validate:"dive,required"`
	Recipient   Recipient `json:"recipient" validate:"dive,required"`
	Destination Platform  `json:"destination" validate:"required,min=1"`
}
