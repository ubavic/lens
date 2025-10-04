package form

import (
	"fmt"
	"strings"

	"github.com/ubavic/lens/view/tooltip"
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
	ID          string
	Type        FieldType
	Name        string
	Description string
	Value       string
	Editable    bool
	Required    bool
	Changeable  bool
	Error       string
}

func (f *Field) Render(sb *strings.Builder) error {
	_, err := sb.WriteString(`<div class="field">`)
	if err != nil {
		return fmt.Errorf("writing field start: %w", err)
	}

	requiredString := ""
	if f.Required {
		requiredString = `<span class="required">&lowast;</span>`
	}

	tooltipString := ""
	if f.Description != "" {
		tooltipString = tooltip.Render(f.Description)
	}

	labelString := fmt.Sprintf(`<div class="labelRow"><label>%s</label>%s</div>%s`, f.Name, requiredString, tooltipString)
	_, err = sb.WriteString(labelString)
	if err != nil {
		return fmt.Errorf("writing field label: %w", err)
	}

	switch f.Type {
	case FieldText:
		err = f.renderTextInput(sb)
	case FieldInteger:
		err = f.renderIntegerInput(sb)
	case FieldDecimal:
		err = f.renderDecimalInput(sb)
	}
	if err != nil {
		return fmt.Errorf("writing field input: %w", err)
	}

	if f.Error != "" {
		errorString := fmt.Sprintf(`<div class="error-message">%s</div>`, f.Error)
		_, err = sb.WriteString(errorString)
		if err != nil {
			return fmt.Errorf("writing field error: %w", err)
		}
	}

	_, err = sb.WriteString(`</div>`)
	if err != nil {
		return fmt.Errorf("writing field input end: %w", err)
	}

	return nil
}

func (f *Field) renderTextInput(sb *strings.Builder) error {

	classes := []string{}
	if f.Error != "" {
		classes = append(classes, "error")
	}

	inputString := fmt.Sprintf(`<input type="text" name="%s" value="%s" class="%s">`, f.Name, f.Value, strings.Join(classes, " "))

	_, err := sb.WriteString(inputString)
	if err != nil {
		return fmt.Errorf("writing field input: %w", err)
	}

	return nil
}

func (f *Field) renderIntegerInput(sb *strings.Builder) error {
	classes := []string{}
	if f.Error != "" {
		classes = append(classes, "error")
	}

	inputString := fmt.Sprintf(`<input type="number" name="%s" value="%s" class="%s">`, f.Name, f.Value, strings.Join(classes, " "))

	_, err := sb.WriteString(inputString)
	if err != nil {
		return fmt.Errorf("writing field input: %w", err)
	}

	return nil
}

func (f *Field) renderDecimalInput(sb *strings.Builder) error {
	classes := []string{}
	if f.Error != "" {
		classes = append(classes, "error")
	}

	inputString := fmt.Sprintf(`<input type="number" name="%s" value="%s" class="%s">`, f.Name, f.Value, strings.Join(classes, " "))

	_, err := sb.WriteString(inputString)
	if err != nil {
		return fmt.Errorf("writing field input: %w", err)
	}

	return nil
}
