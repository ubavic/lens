package component

import "maragu.dev/gomponents"

type Uid = string

type Action struct {
	Kind   uint8
	Caller Uid
	Data   string
}

type Handler func(Action) Component

type HandlerMap map[Uid]Handler

type Component interface {
	Node() gomponents.Node
	Resolve(Uid, HandlerMap)
}
