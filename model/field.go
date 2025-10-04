package model

type Field[T any] struct {
	Name        string
	Description string

	FieldType FieldType
	Nullable  bool
	Array     bool

	MinValue      T
	MaxValue      T
	MinLength     uint
	MaxLength     uint
	InitialValue  T
	AllowedValues []T

	Regex string

	Relation  FieldRelation
	Aggregate AggregateFormula

	Editable bool
	Visible  bool
}

type FieldI interface {
	ColumName() string
	Type() FieldType
}

func (f Field[T]) ColumName() string {
	return ""
}

func (f Field[T]) Type() FieldType {
	return f.FieldType
}
