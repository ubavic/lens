package form

import (
	"fmt"
	"strings"

	"github.com/ubavic/lens/model"
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

func (f *Form) Render(sb *strings.Builder) error {
	_, err := sb.WriteString(`<form>`)
	if err != nil {
		return fmt.Errorf("writing form start: %w", err)
	}

	for _, field := range f.Fields {
		err = field.Render(sb)
		if err != nil {
			return fmt.Errorf("writing field: %w", err)
		}
	}

	_, err = sb.WriteString(`</form>`)
	if err != nil {
		return fmt.Errorf("writing form end: %w", err)
	}

	return nil
}
