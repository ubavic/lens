package model

type FieldType uint

const (
	Boolean FieldType = iota + 1
	Integer
	Float
	Fixed
	Text
	Date
	Time
	DateTime
	Binary
	UUID
	JSON
	XML
)
