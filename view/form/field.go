package form

import (
	"github.com/ubavic/lens/view/component"
	"github.com/ubavic/lens/view/tooltip"
	"maragu.dev/gomponents"
	ghtml "maragu.dev/gomponents/html"
)

type FieldType uint

const (
	FieldText FieldType = iota + 1
	FieldInteger
	FieldDecimal
	FieldSelection
	FieldDate
	FieldTime
	FieldDateTime
	FieldBoolean
)

type Field struct {
	uid         component.Uid
	Type        FieldType
	Name        string
	Description string
	Value       string
	ReadOnly    bool
	Required    bool
	Changeable  bool
	Error       string
	Validate    func(string) error

	MinValue *string
	MaxValue *string
}

func (f *Field) Node() gomponents.Node {

	var entry gomponents.Node
	switch f.Type {
	case FieldText:
		entry = f.renderTextInput()
	case FieldInteger:
		entry = f.renderIntegerInput()
	case FieldDecimal:
		entry = f.renderDecimalInput()
	case FieldSelection:
		entry = f.renderDecimalInput()
	case FieldDate:
		entry = f.renderDateInput()
	case FieldTime:
		entry = f.renderDecimalInput()
	case FieldDateTime:
		entry = f.renderDateTimeInput()
	case FieldBoolean:
		entry = f.renderBooleanInput()
	}

	return ghtml.Div(
		gomponents.Attr("id", f.uid),
		ghtml.Class("field"),
		ghtml.Div(
			ghtml.Class("labelRow"),
			ghtml.Label(gomponents.Text(f.Name)),
			gomponents.If(f.Required, ghtml.Span(ghtml.Class("required"), gomponents.Raw("&lowast;"))),
			gomponents.If(f.Description != "", gomponents.Raw(tooltip.Render(f.Description))),
		),
		entry,
		gomponents.If(f.Error != "",
			ghtml.Div(ghtml.Class("error-message"), gomponents.Text(f.Error)),
		),
	)
}

func (f *Field) renderTextInput() gomponents.Node {
	return ghtml.Input(
		gomponents.Attr("type", "text"),
		gomponents.Attr("lens-target-id", f.uid),
		gomponents.Attr("onchange", "valueChanged(this)"),
		gomponents.Attr("autocomplete", "off"),
		ghtml.Value(f.Value),
		gomponents.If(f.ReadOnly, gomponents.Attr("readonly")),
		gomponents.If(f.Error != "", ghtml.Class("error")),
	)
}

func (f *Field) renderIntegerInput() gomponents.Node {
	return ghtml.Input(
		gomponents.Attr("type", "number"),
		gomponents.Attr("lens-target-id", f.uid),
		gomponents.Attr("onchange", "valueChanged(this)"),
		gomponents.Attr("autocomplete", "off"),
		ghtml.Value(f.Value),
		gomponents.Iff(f.MinValue != nil, func() gomponents.Node { return gomponents.Attr("min", *f.MinValue) }),
		gomponents.Iff(f.MaxValue != nil, func() gomponents.Node { return gomponents.Attr("max", *f.MaxValue) }),
		gomponents.If(f.ReadOnly, gomponents.Attr("readonly")),
		gomponents.If(f.Error != "", ghtml.Class("error")),
	)
}

func (f *Field) renderDecimalInput() gomponents.Node {
	return ghtml.Input(
		gomponents.Attr("type", "number"),
		gomponents.Attr("lens-target-id", f.uid),
		gomponents.Attr("onchange", "valueChanged(this)"),
		gomponents.Attr("autocomplete", "off"),
		ghtml.Value(f.Value),
		gomponents.Iff(f.MinValue != nil, func() gomponents.Node { return gomponents.Attr("min", *f.MinValue) }),
		gomponents.Iff(f.MaxValue != nil, func() gomponents.Node { return gomponents.Attr("max", *f.MaxValue) }),
		gomponents.If(!f.ReadOnly, gomponents.Attr("readonly")),
		gomponents.If(f.Error != "", ghtml.Class("error")),
	)
}

func (f *Field) renderBooleanInput() gomponents.Node {
	return ghtml.Input(
		gomponents.Attr("type", "checkbox"),
		gomponents.Attr("lens-target-id", f.uid),
		gomponents.Attr("onchange", "valueChanged(this)"),
		gomponents.Attr("autocomplete", "off"),
		ghtml.Value(f.Value),
		gomponents.If(f.ReadOnly, gomponents.Attr("readonly")),
		gomponents.If(f.Error != "", ghtml.Class("error")),
	)
}

func (f *Field) renderDateInput() gomponents.Node {
	return ghtml.Input(
		gomponents.Attr("type", "date"),
		gomponents.Attr("lens-target-id", f.uid),
		gomponents.Attr("onchange", "valueChanged(this)"),
		gomponents.Attr("autocomplete", "off"),
		ghtml.Value(f.Value),
		gomponents.Iff(f.MinValue != nil, func() gomponents.Node { return gomponents.Attr("min", *f.MinValue) }),
		gomponents.Iff(f.MaxValue != nil, func() gomponents.Node { return gomponents.Attr("max", *f.MaxValue) }),
		gomponents.If(f.ReadOnly, gomponents.Attr("readonly")),
		gomponents.If(f.Error != "", ghtml.Class("error")),
	)
}

func (f *Field) renderDateTimeInput() gomponents.Node {
	return ghtml.Input(
		gomponents.Attr("type", "datetime-local"),
		gomponents.Attr("lens-target-id", f.uid),
		gomponents.Attr("onchange", "valueChanged(this)"),
		gomponents.Attr("autocomplete", "off"),
		ghtml.Value(f.Value),
		gomponents.Iff(f.MinValue != nil, func() gomponents.Node { return gomponents.Attr("min", *f.MinValue) }),
		gomponents.Iff(f.MaxValue != nil, func() gomponents.Node { return gomponents.Attr("max", *f.MaxValue) }),
		gomponents.If(f.ReadOnly, gomponents.Attr("readonly")),
		gomponents.If(f.Error != "", ghtml.Class("error")),
	)
}

func (f *Field) Resolve(id component.Uid, hm component.HandlerMap) {
	hm[id] = func(a component.Action) component.Component {
		f.Error = ""

		f.Value = a.Data

		if f.Validate != nil {
			e := f.Validate(f.Value)
			if e != nil {
				f.Error = e.Error()
			}
		}

		return f
	}

	f.uid = id
}
