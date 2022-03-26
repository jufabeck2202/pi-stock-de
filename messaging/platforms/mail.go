package platforms

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"strconv"
	"sync"

	mail "github.com/xhit/go-simple-mail/v2"

	"github.com/jufabeck2202/piScraper/messaging/types"
	"github.com/jufabeck2202/piScraper/utils"
)

type Mail struct {
}

func NewMail() Mail {

	return Mail{}
}

var lock = sync.Mutex{}

func (m Mail) Send(recipient types.Recipient, item utils.Website) error {
	server := mail.NewSMTPClient()
	if n, err := strconv.Atoi(os.Getenv("MAIL_PORT")); err == nil {
		server.Port = n
		server.Username = os.Getenv("MAIL_USERNAME")
		server.Password = os.Getenv("MAIL_PASSWORD")
		server.Encryption = mail.EncryptionSSL

	}
	server.Host = os.Getenv("MAIL_HOST")
	lock.Lock()
	smtpClient, err := server.Connect()
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
	temp.Parse(htmlBody)
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

var htmlBody = `
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
