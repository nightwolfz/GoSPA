package controllers

import (
	"github.com/revel/revel"
	"sequoia/app/models"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {

	greeting := "SEQUOIA"

	return c.Render(greeting)
}

func (c App) Mail(post models.Talent) revel.Result {

	c.Validation.Required(post.Name).Message("Please include your name.")
	c.Validation.Required(post.Phone).Message("Please include your phone.")
	c.Validation.Required(post.Email).Message("Please include your email.")
	/*c.Validation.Required(post.Reference).Message("Please cite who has refered you to us.")
	c.Validation.Required(post.CV).Message("Please attach your CV.")*/

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}

	c.Flash.Success("we_got_mail")

	//return c.Render(post.Email)
	return c.Redirect(App.Index)
}
