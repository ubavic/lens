package view

import "maragu.dev/gomponents"

type Component interface {
	Node() gomponents.Node
	ResolveIds(string)
}
