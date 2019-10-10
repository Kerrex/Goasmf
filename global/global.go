package global

import (
	"context"
	"goasmf/component"
)

var componentFactories map[string]func(context.Context) component.Component = make(map[string]func(context.Context) component.Component)

func RegisterComponent(name string, factoryFunction func(context.Context) component.Component) {
	componentFactories[name] = factoryFunction
}

func GetComponentFactoryByName(name string) func(ctx context.Context) component.Component {
	return componentFactories[name]
}
