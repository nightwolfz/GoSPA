package models

import (
	"bytes"
	"github.com/revel/revel"
	"github.com/jordan-wright/email"
	"net/smtp"
	"log"
)

type Email struct {
	To                  []string
	Subject, From, Body string
	TemplatePath        string
	Attachment			string
}

func (t Email) SendEmail(args map[string]interface{}) {
	e := email.NewEmail()
	e.From = t.From
	e.To = t.To
	e.Subject = t.Subject
	e.Text = []byte("Text Body is, of course, supported!")
	e.HTML = getViewTemplate(t.TemplatePath, args).Bytes()
	if t.Attachment != "" {
		e.AttachFile(t.Attachment)
	}
	e.Send("127.0.0.1:25", smtp.PlainAuth("", "", "", "127.0.0.1"))
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