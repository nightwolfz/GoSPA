package models

import (
	"github.com/revel/revel"
)

type Contact struct {
	Name, Company, Job      string
	Email, Phone, Reference string
	CV                      []byte
	IsClient 				bool
}


func (t *Contact) Validate(v *revel.Validation) {

	v.Check(t.Name,
		revel.Required{},
		revel.MaxSize{20},
		revel.MinSize{4},
	)

	if t.Phone == "" {
		v.Check(t.Email,
			revel.Required{},
			revel.MaxSize{20},
			revel.MinSize{4},
		)
	}

	if t.Email == "" {
		v.Check(t.Phone,
			revel.Required{},
			revel.MaxSize{20},
			revel.MinSize{4},
		)
	}

}
