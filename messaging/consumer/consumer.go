package consumer

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/adjust/rmq/v4"

	"github.com/jufabeck2202/piScraper/messaging/platforms"
	"github.com/jufabeck2202/piScraper/messaging/types"
	"github.com/jufabeck2202/piScraper/services"
)

var mailVerifier = services.MailVerifier{}

type Consumer struct {
	Name    string
	Created time.Time
}

func NewConsumer(tag int) *Consumer {
	return &Consumer{
		Name:    fmt.Sprintf("consumer%d", tag),
		Created: time.Now(),
	}
}

func (consumer *Consumer) Consume(delivery rmq.Delivery) {
	var task types.AlertTask
	if err := json.Unmarshal([]byte(delivery.Payload()), &task); err != nil {
		// handle json error
		if err := delivery.Reject(); err != nil {
			fmt.Println("failed to reject: ", err)
		}
		return
	}
	send(task)
	if err := delivery.Ack(); err != nil {
		fmt.Println("failed to ack ", err)
	}
}

func send(item types.AlertTask) {
	switch item.Destination {

	case types.PushHover:
		err := platforms.NewPushover(item.Recipient).Send(item.Website)
		if err != nil {
			fmt.Println("failed to send pushhover: ", err)
		}
	case types.Mail:
		//check if email is verified
		verified := mailVerifier.IsVerified(item.Recipient.Email)
		if !verified {
			fmt.Println("Email not verified", item.Recipient.Email)
		}
		err := platforms.NewMail().Send(item.Recipient, item.Website)
		if err != nil {
			fmt.Println("failed to send pushhover: ", err)
		}
	case types.Webhook:
		err := platforms.NewWebhook().Send(item.Recipient, item.Website)
		if err != nil {
			fmt.Println("failed to send pushhover: ", err)
		}
	}

}
