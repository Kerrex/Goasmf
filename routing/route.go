package routing

import component "goasmf/component"

// Route interface for different route types, eg. different renderer, etc.
type Route interface {
	GetPath() string
	GetComponentName() string
	GetChildren() []Route
}

type BaseRoute struct {
	Path string
	ComponentName string
	Children []Route
}

func (b *BaseRoute) GetPath() string {
	return b.Path
}

func (b *BaseRoute) GetComponentName() string {
	return b.ComponentName
}

func (b *BaseRoute) GetChildren() []Route {
	return b.Children
}

type RenderedRoute interface {
	GetRoute() Route
	GetComponent() component.Component
}

type BaseRenderedRoute struct {
	route     Route
	component component.Component
}

func (b BaseRenderedRoute) GetRoute() Route {
	return b.route
}

func (b BaseRenderedRoute) GetComponent() component.Component {
	return b.component
}

