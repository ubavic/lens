package model

type Model struct {
	Name        string
	Description string

	TableName string

	Virtual bool

	Fields []FieldI
}
