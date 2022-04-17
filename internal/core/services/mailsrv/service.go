package mailsrv

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
	"text/template"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"

	"github.com/jufabeck2202/piScraper/internal/core/domain"
	"github.com/jufabeck2202/piScraper/internal/core/ports"
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

func (srv *service) Verify(email string) (string, error) {
	decytedEmail := Decrypt(email)
	exists := srv.redisRepository.Exists(decytedEmail)
	if !exists {
		return "", fmt.Errorf("email not found")
	}
	err := srv.redisRepository.Set(decytedEmail, true, 0)
	if err != nil {
		log.Println("error Mail Verifier empty: ", err)
		return "", err
	}
	return decytedEmail, nil
}

// Check if email is verified
func (srv *service) IsVerified(email string) bool {
	return srv.redisRepository.GetBool(email)
}

//Adds temporary redis entry with email and generates verificaiton email
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
	log.Println("creating redis entry", email)
	err := srv.redisRepository.Set(email, false, time.Hour*24*2)
	if err != nil {
		log.Println("error Mail Verifier empty: ", err)
		return err
	}
	return nil
}

func (srv *service) unsubscribeLinkBuilder(email string) string {

	return "https://" + os.Getenv("HOST_URL") + "/unsubscribe/" + Encrypt(email)

}
func (srv *service) verifyLinkBuilder(email string) string {

	return "https://" + os.Getenv("HOST_URL") + "/verify/" + Encrypt(email)

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
	log.Println("creating redis entry")
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

func Encrypt(input string) string {

	text := []byte(input)
	// generate a new aes cipher using our 32 byte long key
	c, err := aes.NewCipher([]byte(os.Getenv("ENCRYPTION_KEY")))
	// if there are any errors, handle them
	if err != nil {
		fmt.Println(err)
	}

	// gcm or Galois/Counter Mode, is a mode of operation
	// for symmetric key cryptographic block ciphers
	// - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	gcm, err := cipher.NewGCM(c)
	// if any error generating new GCM
	// handle them
	if err != nil {
		fmt.Println(err)
	}

	// creates a new byte array the size of the nonce
	// which must be passed to Seal
	nonce := make([]byte, gcm.NonceSize())
	// populates our nonce with a cryptographically secure
	// random sequence
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}

	// here we encrypt our text using the Seal function
	// Seal encrypts and authenticates plaintext, authenticates the
	// additional data and appends the result to dst, returning the updated
	// slice. The nonce must be NonceSize() bytes long and unique for all
	// time, for a given key.
	return base64.RawURLEncoding.EncodeToString(gcm.Seal(nonce, nonce, text, nil))
}

func Decrypt(input string) string {
	data, err := base64.RawURLEncoding.DecodeString(input)
	if err != nil {
		log.Fatal("error:", err)
	}
	cypher := []byte(data)

	c, err := aes.NewCipher([]byte(os.Getenv("ENCRYPTION_KEY")))
	if err != nil {
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}

	nonceSize := gcm.NonceSize()
	if len(cypher) < nonceSize {
		fmt.Println(err)
	}

	nonce, ciphertext := cypher[:nonceSize], cypher[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println(err)
	}
	return string(plaintext)
}
