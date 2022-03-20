package platforms

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/jufabeck2202/piScraper/messaging/types"
	"github.com/jufabeck2202/piScraper/utils"
)

type webhook struct {
}

func NewWebhook() *webhook {
	return &webhook{}
}

func (p webhook) Send(recipient types.Recipient, item utils.Website) error {
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
