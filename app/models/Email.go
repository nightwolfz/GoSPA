package models

import (
	"bytes"
	"log"
	"net/smtp"
)

type Email struct {
	Smtp, From, To, Body string
}

func (t Email) Send() {

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
}
