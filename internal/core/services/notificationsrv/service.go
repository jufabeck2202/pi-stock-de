package notificationsrv

import (
	"fmt"
	"log"
	"sync"

	"github.com/jufabeck2202/piScraper/internal/core/domain"
	"github.com/jufabeck2202/piScraper/internal/core/ports"
)

type service struct {
	mailVerifierService ports.MailService
	pushoverClient      ports.MessagingPlatform
	webhookClient       ports.MessagingPlatform
}

type AlertTaskQueue struct {
	Website     domain.Website
	Recipient   domain.Recipient
	Destination domain.Platform
	Err         error
}

func NewNotificationService(mailVerifierService ports.MailService, pushoverClient ports.MessagingPlatform, webhookClient ports.MessagingPlatform) *service {
	return &service{
		mailVerifierService: mailVerifierService,
		pushoverClient:      pushoverClient,
		webhookClient:       webhookClient,
	}
}
func (srv *service) send(item AlertTaskQueue, wg *sync.WaitGroup) {
	defer wg.Done()
	switch item.Destination {

	case domain.PushHover:
		err := srv.pushoverClient.Send(item.Recipient, item.Website)
		if err != nil {
			fmt.Println("failed to send pushhover: ", err)
			item.Err = fmt.Errorf("failed to send pushhover: %v", err)
		}
	case domain.Mail:
		//check if email is verified
		verified := srv.mailVerifierService.IsVerified(item.Recipient.Email)
		if !verified {
			fmt.Println("Email not verified", item.Recipient.Email)
			item.Err = fmt.Errorf("email not verified")
			return
		}
		err := srv.mailVerifierService.Send(item.Recipient, item.Website)
		if err != nil {
			fmt.Println("failed to send pushhover: ", err)
			item.Err = fmt.Errorf("failed to send email: %v", err)
		}
	case domain.Webhook:
		err := srv.webhookClient.Send(item.Recipient, item.Website)
		if err != nil {
			fmt.Println("failed to send webhook: ", err)
			item.Err = fmt.Errorf("failed to send webhook: %v", err)
		}
	}

}

func (srv *service) Notifiy(alerts []domain.AlertTask) {
	alertQueue := make([]AlertTaskQueue, len(alerts))
	for i, alert := range alerts {
		alertQueue[i] = AlertTaskQueue{
			Website:     alert.Website,
			Recipient:   alert.Recipient,
			Destination: alert.Destination,
		}
	}
	var wg sync.WaitGroup
	for _, v := range alertQueue {
		wg.Add(1)
		go srv.send(v, &wg)
	}
	wg.Wait()

	var successCount, failCount int
	for _, v := range alertQueue {
		if v.Err == nil {
			successCount++
		} else {
			failCount++
		}
	}

	if failCount == 0 {
		log.Printf("Send all notifications successfully: %d \n", successCount)
	} else {
		log.Printf("Finished - but %d errors and %d successes\n", failCount, successCount)
	}

}
