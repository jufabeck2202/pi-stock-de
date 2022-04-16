package mailsrv

import (
	"log"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"

	"github.com/jufabeck2202/piScraper/internal/core/ports"
)

type service struct {
	redisRepository ports.RedisRepository
	server          *mail.SMTPServer
}

func New(redisRepository ports.RedisRepository) *service {
	return &service{
		redisRepository: redisRepository,
	}
}

func (srv *service) IsVerified(email string) bool {
	return srv.redisRepository.GetBool(email)
}

func (srv *service) NewEmailSubscriber(email string) error {
	err := srv.createRedisEntry(email)
	if err != nil {
		return err
	}
	mail := mail.Mail{}
	err = mail.SendVerificationMail(email)
	if err != nil {
		return err
	}
	return nil

}

func (srv *service) createRedisEntry(email string) error {
	err := srv.redisRepository.Set(email, false, 300*time.Second)
	if err != nil {
		log.Println("error Mail Verifier empty: ", err)
		return err
	}
	return nil
}

func (srv *service) Verifiy(email string) error {
	err := srv.redisRepository.Set(email, true, 0)
	if err != nil {
		log.Println("error Mail Verifier empty: ", err)
		return err
	}
	return nil
}
