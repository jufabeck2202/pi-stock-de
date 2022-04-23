package webhook

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/jufabeck2202/piScraper/internal/core/domain"
)

type webhook struct {
}

func NewWebhook() webhook {
	return webhook{}
}

func (p webhook) Send(recipient domain.Recipient, item domain.Website) error {
	// struct to json
	json, err := json.Marshal(item)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", recipient.Webhook, bytes.NewBuffer(json))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
