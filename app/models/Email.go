package models

import (
	"bytes"
	"github.com/revel/revel"
	"github.com/nightwolfz/email"
	"net/smtp"
	"log"
)

type Email struct {
	To                  string
	Subject, From, Body string
	TemplatePath        string
	Attachment			string
}


func (t Email) SendEmail(args map[string]interface{}) {
	e := email.NewEmail()
	e.From = t.From
	e.To = []string{ t.To }
	e.Subject = t.Subject
	e.Text = []byte("Text Body is, of course, supported!")
	e.HTML = getViewTemplate(t.TemplatePath, args).Bytes()
	if t.Attachment != "" {
		e.AttachFile(t.Attachment)
	}
	
	auth := smtp.PlainAuth("", "gospa@gmail.com", "aaaaaa7*", "smtp.gmail.com")
	e.Send("smtp.gmail.com:587", auth)
}

/*func (t Email) SendEmail(args map[string]interface{}) {
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
}*/

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
