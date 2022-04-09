package services

import (
	"context"
	"log"
	"time"

	"github.com/jufabeck2202/piScraper/messaging/platforms"
	"github.com/jufabeck2202/piScraper/storage"
)

type MailVerifier struct{}

func (m MailVerifier) IsVerified(email string) bool {
	redis, err := storage.GetRedisConnection()
	if err != nil {
		log.Println("error Mail Verifier Connection: ", err)
		return false
	}
	data := redis.Get(context.Background(), email)
	verified, err := data.Bool()
	if err != nil {
		log.Println("error Mail Verifier empty: ", err)
		return false
	}
	return verified
}

func (m MailVerifier) NewEmailSubscriber(email string) error {
	err := m.createRedisEntry(email)
	if err != nil {
		return err
	}
	mail := platforms.Mail{}
	err = mail.SendVerificationMail(email)
	if err != nil {
		return err
	}
	return nil

}

func (m MailVerifier) createRedisEntry(email string) error {
	redis, err := storage.GetRedisConnection()
	if err != nil {
		log.Println("error Mail Verifier Connection: ", err)
		return err
	}
	data := redis.Set(context.Background(), email, false, 300*time.Second)
	if data.Err() != nil {
		log.Println("error Mail Verifier empty: ", data.Err())
		return data.Err()
	}
	return nil
}

func (m MailVerifier) Verifiy(email string) error {
	redis, err := storage.GetRedisConnection()
	if err != nil {
		log.Println("error Mail Verifier Connection: ", err)
		return err
	}
	data := redis.Set(context.Background(), email, true, 0)
	if data.Err() != nil {
		log.Println("error Mail Verifier empty: ", data.Err())
		return data.Err()
	}
	return nil
}
