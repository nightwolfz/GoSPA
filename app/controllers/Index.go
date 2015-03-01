package controllers

import (
	"github.com/revel/revel"
	"sequoia/app/models"
	"fmt"
)

type Index struct {
	*revel.Controller
}

func (c Index) Index() revel.Result {

	greeting := "SEQUOIA"

	return c.Render(greeting)
}

func (c Index) Contact(post models.Contact) revel.Result {

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Index.Index)
	}

	var args = make(map[string]interface{})
	args["info"] = post

	email := new(models.Email)

	if post.IsClient {
		email.Subject = fmt.Sprintf("Client: %s", post.Name)
	} else {
		email.Subject = fmt.Sprintf("Talent: %s", post.Name)
	}

	email.From = "sales@sequoia-projects.be"
	email.To = []string{"ryan@megidov.com"}
	email.TemplatePath = "Contact/Email.template"
	email.SendEmail(args)

	return c.Render(post)
}
