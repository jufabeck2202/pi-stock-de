package types

type Recipient struct {
	Webhook  string `json:"webhook"`
	Pushover string `json:"pushover"`
	Email    string `json:"email"`
}
