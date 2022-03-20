package platforms

import (
	"fmt"
	"os"
	"time"

	"github.com/gregdel/pushover"

	"github.com/jufabeck2202/piScraper/messaging/types"
	"github.com/jufabeck2202/piScraper/utils"
)

type Pushover struct {
	client         *pushover.Pushover
	localRecipient *pushover.Recipient
}

func NewPushover(recipient types.Recipient) Pushover {

	client := pushover.New(os.Getenv("PUSHOVER_CLIENT"))
	localRecipient := pushover.NewRecipient(recipient.Pushover)
	return Pushover{client: client, localRecipient: localRecipient}
}

func (p Pushover) Send(item utils.Website) error {
	input := &pushover.Message{
		Message:   "Pi is in Stock:" + item.Name + "for" + item.PriceString + "at" + item.URL,
		URL:       item.URL,
		Timestamp: time.Now().Unix(),
	}
	response, err := p.client.SendMessage(input, p.localRecipient)
	fmt.Println(response)
	return err
}
