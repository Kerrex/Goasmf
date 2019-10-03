package global

import (
	"context"
	"goasmf/component"
	"goasmf/rendering"
)

var ComponentFactories map[string]func(context.Context) component.Component = make(map[string]func(context.Context) component.Component)
var CurrentRouteComponents map[string]component.Component = make(map[string]component.Component)
var MainRenderer rendering.Renderer = rendering.NewHandlebarsRenderer()

func RegisterComponent(name string, factoryFunction func(context.Context) component.Component) {
	ComponentFactories[name] = factoryFunction
}
