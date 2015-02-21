package models

import (
	"github.com/revel/revel"
)

type Talent struct {
	Name, Email, Phone, Reference string
	CV                            []byte
}

func (t *Talent) Validate(v *revel.Validation) {
	v.Check(t.Name,
		revel.Required{},
		revel.MaxSize{20},
		revel.MinSize{4},
	)

	v.Check(t.Email,
		revel.Required{},
		revel.MaxSize{20},
		revel.MinSize{4},
	)

	v.Check(t.Phone,
		revel.Required{},
		revel.MaxSize{20},
		revel.MinSize{4},
	)

	v.Check(t.Reference,
		revel.Required{},
		revel.MaxSize{20},
		revel.MinSize{4},
	)
}
