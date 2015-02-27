package models

import (
	"bytes"
	"github.com/revel/revel/mail"
	//"log"
	//"net/smtp"
)

type Email struct {
	To                  []string
	Subject, From, Body string
	TemplatePath        string
}

func (t Email) SendEmail(args map[string]interface{}) {

	/*auth := smtp.PlainAuth("",
		username,
		password,
		server,
	)*/

	mailer := mail.Mailer{Server: "127.0.0.1", Port: 25 /*, UserName:username, Password:password, Auth:auth*/}
	message := mail.NewHtmlMessage(t.To, t.Subject, t.Body)
	err := message.RenderTemplate(t.TemplatePath, args)
	if err != nil {
		message.HtmlBody = bytes.NewBufferString(t.Body)
	}
	mailer.SendMessage(message)
}

/*func (t Email) Send2() {

	// Connect to the remote SMTP server.
	c, err := smtp.Dial("127.0.0.1:25")
	if err != nil {
		log.Fatal(err)
	}

	// Set the sender and recipient.
	c.Mail(t.From)
	c.Rcpt(t.To)

	// Send the email body.
	wc, err := c.Data()
	if err != nil {
		log.Fatal(err)
	}
	defer wc.Close()

	buf := bytes.NewBufferString(t.Body)
	if _, err = buf.WriteTo(wc); err != nil {
		log.Fatal(err)
	}
}*/
