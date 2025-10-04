package view

import (
	"github.com/ubavic/lens/view/form"
	"maragu.dev/gomponents"
)

func RenderHomePage() gomponents.Node {
	form := form.Form{
		Fields: []form.Field{
			{
				Name:        "Name",
				Type:        form.FieldText,
				Required:    true,
				Description: "Name of the field",
			},
		},
	}

	return form.Node()
}
