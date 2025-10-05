package form

import (
	"strconv"

	"github.com/ubavic/lens/model"
	"maragu.dev/gomponents"
	ghtml "maragu.dev/gomponents/html"
)

type Form struct {
	uid    string
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
		gomponents.Attr("hx-post", "."),
		gomponents.Attr("autocomplete", "off"),
		gomponents.Map(f.Fields, func(f Field) gomponents.Node { return f.Node() }),
		ghtml.Input(gomponents.Attr("type", "submit"), gomponents.Attr("value", "save"), gomponents.Attr("name", f.uid)),
	)
}

func (f *Form) ResolveIds(id string) {
	f.uid = id

	for i, fi := range f.Fields {
		fi.ResolveIds(id + "-" + strconv.Itoa(i))
	}
}
