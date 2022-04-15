package pushover

import (
	"fmt"
	"os"
	"time"

	"github.com/gregdel/pushover"

	"github.com/jufabeck2202/piScraper/internal/core/domain"
)

type Pushover struct {
	client *pushover.Pushover
}

func NewPushover() Pushover {
	client := pushover.New(os.Getenv("PUSHOVER_CLIENT"))
	return Pushover{client: client}
}

func (p Pushover) Send(recipient domain.Recipient, item domain.Website) error {
	localRecipient := pushover.NewRecipient(recipient.Pushover)
	input := &pushover.Message{
		Message:   "Pi is in Stock:" + item.Name + "for" + item.PriceString + "at" + item.URL,
		URL:       item.URL,
		Timestamp: time.Now().Unix(),
	}
	response, err := p.client.SendMessage(input, localRecipient)
	fmt.Println(response)
	return err
}
