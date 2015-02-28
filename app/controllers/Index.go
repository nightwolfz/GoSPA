package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"sequoia/app/models"
)

type Index struct {
	*revel.Controller
}

func (c Index) Index() revel.Result {

	greeting := "SEQUOIA"

	return c.Render(greeting)
}

func (c Index) Contact(post models.Contact) revel.Result {

	//c.Validation.Required(post.Name).Message("Please include your name.")
	//c.Validation.Required(post.Phone).Message("Please include your phone.")
	//c.Validation.Required(post.Email).Message("Please include your email.")
	/*c.Validation.Required(post.Reference).Message("Please cite who has refered you to us.")
	c.Validation.Required(post.CV).Message("Please attach your CV.")*/

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Index.Index)
	}

	post.Name = "money!"

	fmt.Printf("%+v", post.Name)

	//c.Flash.Success("we_got_mail")

	var args = make(map[string]interface{})
	args["FirstName"] = "Whizdumb"

	email := new(models.Email)
	email.Subject = "Welcome to Sequoia!"
	email.From = "noreply@ragefx.com"
	email.To = []string{"ryan@megidov.com"}
	email.Body = "This is a sample body!"
	email.TemplatePath = "Contact/Email.template"
	email.SendEmail(args)

	return c.Render(post)
	//return c.Redirect(Index.Contact, post)
}