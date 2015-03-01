package models

import (
	//"bytes"
	"github.com/revel/revel/mail"
	"log"
)

type Email struct {
	To                  []string
	Subject, From, Body string
	TemplatePath        string
}

func (t Email) SendEmail(args map[string]interface{}) {

	server := "127.0.0.1"
    /*server := "smtp-auth.mailprotect.be"
    username := "email address of sender"
    password := "password"*/
 
    /*auth := smtp.PlainAuth("",
        username,
        password,
        server,
    )*/


	mailer := mail.Mailer{Server: server, Port: 25/*, UserName:username, Password:password, Auth:auth*/}
	message := mail.Message{
		From:      t.From,
		To:        t.To,
		Subject:   t.Subject,
	}

	err := message.RenderTemplate(t.TemplatePath, args)
	if err != nil {
		log.Println(err)
	}

	err = mailer.SendMessage(&message)
	if err != nil {
		log.Fatal(err)
	}
}
