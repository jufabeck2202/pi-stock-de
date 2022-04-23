package domain

type Alert struct {
	Recipient   Recipient `json:"recipient"`
	Destination Platform  `json:"destination"`
}
