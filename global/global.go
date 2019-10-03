package global

import (
	"context"
	"goasmf/component"
)

var ComponentFactories map[string]func(context.Context) component.Component = make(map[string]func(context.Context) component.Component)

func RegisterComponent(name string, factoryFunction func(context.Context) component.Component) {
	ComponentFactories[name] = factoryFunction
}

