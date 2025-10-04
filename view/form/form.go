package form

import (
	"github.com/ubavic/lens/model"
	"maragu.dev/gomponents"
	ghtml "maragu.dev/gomponents/html"
)

type Form struct {
	Fields []Field
}

func NewForm(fields ...model.FieldI) Form {
	formFields := make([]Field, 0, len(fields))

	return Form{
		Fields: formFields,
	}
}

func (f *Form) Node() gomponents.Node {
	return ghtml.Form(
		gomponents.Map(f.Fields, func(f Field) gomponents.Node { return f.Node() }),
	)
}
