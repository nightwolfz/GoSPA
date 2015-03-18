package controllers

import (
	"github.com/revel/revel"
	"sequoia/app/models"
	"fmt"
	"log"
	"io"
	"os"
	"mime/multipart"
)

type Index struct {
	*revel.Controller
}

func (c Index) Index() revel.Result {

	greeting := "SEQUOIA"

	return c.Render(greeting)
}

func GetUploadedFilePath(m *multipart.Form) string {
	for fname, _ := range m.File {
		fheaders := m.File[fname]
		for i, _ := range fheaders {
			file, err := fheaders[i].Open()
			defer file.Close()
			if err != nil {
				log.Print(err)
			}
			dst_path := "/tmp/" + fheaders[i].Filename
			dst, err := os.Create(dst_path)
			defer dst.Close() 
			defer os.Chmod(dst_path, (os.FileMode)(0644))
			if err != nil {
				log.Print(err)
			}
			if _, err := io.Copy(dst, file); err != nil {
				log.Print(err)
			}
			return dst_path
		}
	}
	return ""
}

func (c Index) Contact(post models.Contact) revel.Result {

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Index.Index)
	}

	var args = make(map[string]interface{})
	args["info"] = post
	args["Attachment"] = "In attachment (if any)"

	email := new(models.Email)

	if post.IsClient {
		email.Subject = fmt.Sprintf("Client: %s", post.Name)
	} else {
		email.Subject = fmt.Sprintf("Talent: %s", post.Name)
	}

	email.From = "Sales <sequoia.projects@gmail.com>"
	email.To = "sales@sequoia-projects.be"
	email.TemplatePath = "Contact/Email.template"
	email.Attachment = GetUploadedFilePath(c.Request.MultipartForm)
	email.SendEmail(args)

	return c.Render(post)
}

/*func (c Index) Contact(post models.Contact) revel.Result {

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Index.Index)
	}

	var args = make(map[string]interface{})
	args["info"] = post
	args["Attachment"] = "In attachment (if any)"


	email := new(models.Email)

	if post.IsClient {
		email.Subject = fmt.Sprintf("Client: %s", post.Name)
	} else {
		email.Subject = fmt.Sprintf("Talent: %s", post.Name)
	}

	email.From = "Sales <sequoia.projects@gmail.com>"
	email.To = "sales@sequoia-projects.be"
	email.TemplatePath = "Contact/Email.template"
	email.Attachment = GetUploadedFilePath(c.Request.MultipartForm)
	email.SendEmail(args)

	return c.Render(post)
}*/
