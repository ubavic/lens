package view

import "strings"

type Component interface {
	Render(sb *strings.Builder) error
}
