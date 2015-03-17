package models

import (
	"bytes"
	"github.com/revel/revel"
	//"github.com/jordan-wright/email"
	"github.com/SlyMarbo/gmail"
	//"net/smtp"
	"log"
)

type Email struct {
	To                  string
	Subject, From, Body string
	TemplatePath        string
	Attachment			string
}

func (t Email) SendEmail(args map[string]interface{}) {

	/*
	mTLSConfig := &tls.Config{
	    InsecureSkipVerify: true,
	    ServerName:         net.SplitHostPort("smtp.gmail.com:587"),
	}*/
	//com-basicem-smtp003.srv.combell-ops.net:25

	body := string(getViewTemplate(t.TemplatePath, args).Bytes()[:])

	e := gmail.Compose(t.Subject, body)
	e.From = ""
	e.Password = ""
	e.ContentType = "text/html; charset=utf-8"
	e.AddRecipient(t.To)
	e.Attach(t.Attachment)
	err := e.Send()
	if err != nil {
		log.Print(err)
	}
}

// Returns a parsed view template
func getViewTemplate(templateFilePath string, args map[string]interface{}) *bytes.Buffer {
	// Get the Template.
	template, err := revel.MainTemplateLoader.Template(templateFilePath+".html")
	if err != nil {
		log.Print(err)
		return nil
	}

	var b bytes.Buffer

	err = template.Render(&b, args)
	if err != nil {
		log.Print(err)
		return nil
	}

	return &b
}