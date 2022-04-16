package mailsrv

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"text/template"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"

	"github.com/jufabeck2202/piScraper/internal/core/domain"
	"github.com/jufabeck2202/piScraper/internal/core/ports"
	"github.com/jufabeck2202/piScraper/utils"
)

type service struct {
	redisRepository ports.RedisRepository
	server          *mail.SMTPServer
}

type VerifyEmailTemplate struct {
	Email string
	Link  string
}

func New(redisRepository ports.RedisRepository) *service {
	server := mail.NewSMTPClient()
	if n, err := strconv.Atoi(os.Getenv("MAIL_PORT")); err == nil {
		server.Port = n
		server.Username = os.Getenv("MAIL_USERNAME")
		server.Password = os.Getenv("MAIL_PASSWORD")
		server.Encryption = mail.EncryptionSSL

	}
	server.Host = os.Getenv("MAIL_HOST")
	return &service{
		redisRepository: redisRepository,
		server:          server,
	}
}

func (srv *service) Verify(email string) error {
	err := srv.redisRepository.Set(email, true, 0)
	if err != nil {
		log.Println("error Mail Verifier empty: ", err)
		return err
	}
	return nil
}

func (srv *service) IsVerified(email string) bool {
	return srv.redisRepository.GetBool(email)
}

func (srv *service) NewEmailSubscriber(email string) error {
	err := srv.createRedisEntry(email)
	if err != nil {
		return err
	}
	err = srv.SendVerificationMail(email)
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

func (srv *service) unsubscribeLinkBuilder(email string) string {

	return "https://" + os.Getenv("HOST_URL") + "/unsubscribe/" + utils.Encrypt(email)

}
func (srv *service) verifyLinkBuilder(email string) string {

	return "https://" + os.Getenv("HOST_URL") + "/verify/" + utils.Encrypt(email)

}
func (srv *service) Send(recipient domain.Recipient, item domain.Website) error {
	lock.Lock()
	smtpClient, err := srv.server.Connect()
	if err != nil {
		return err
	}
	defer lock.Unlock()
	defer smtpClient.Close()
	// Create email
	email := mail.NewMSG()
	email.SetFrom("PI-Stock <" + os.Getenv("MAIL_USERNAME") + ">")
	email.AddTo(recipient.Email)
	email.SetSubject(item.Name + " is in Stock!")
	var tpl bytes.Buffer

	temp := template.New("t1")
	temp.Parse(inStockBody)
	err = temp.Execute(&tpl, item)
	if err != nil {
		return err
	}
	email.SetBody(mail.TextHTML, tpl.String())

	// Send email^
	err = email.Send(smtpClient)
	if err != nil {
		return err
	}
	fmt.Println("Email sent! to " + recipient.Email)
	return nil
}

var lock = sync.Mutex{}

func (srv *service) SendVerificationMail(newEmail string) error {
	lock.Lock()
	smtpClient, err := srv.server.Connect()
	if err != nil {
		return err
	}
	defer lock.Unlock()
	defer smtpClient.Close()
	// Create email
	email := mail.NewMSG()
	email.SetFrom("PI-Stock <" + os.Getenv("MAIL_USERNAME") + ">")
	email.AddTo(newEmail)
	email.SetSubject("Verify your email for PI-Stock")
	var tpl bytes.Buffer

	temp := template.New("t1")
	temp.Parse(verifyEmail)
	templateData := VerifyEmailTemplate{
		Email: newEmail,
		Link:  srv.verifyLinkBuilder(newEmail),
	}
	err = temp.Execute(&tpl, templateData)
	if err != nil {
		return err
	}
	email.SetBody(mail.TextHTML, tpl.String())

	// Send email^
	err = email.Send(smtpClient)
	if err != nil {
		return err
	}
	fmt.Println("Email sent! to " + newEmail)
	return nil
}

var inStockBody = `
<html>
<head>
   <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
   <title>The Pi {{.Name}} is availabe!</title>
</head>
<body>
<h1>The Pi {{.Name}} is availabe for {{.PriceString}}!</h1>
   <p>Check it out here: {{.URL}}</p>
   <P> To Unsubscribe visit https://pi.juli.sh, click the Unsubscribe-Button and enter your Email adress</P>
</body>
`

var verifyEmail = `
<html>
<head>
   <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
   <title>Verify Your Email for Pi-Stock</title>
</head>
<body>
<h4>Please Verify your Email {{.Email}} for Pi-Stock</h4>
   <P>To Verify click the following link: {{.Link}}</P>
</body>
`
