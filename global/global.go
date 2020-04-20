package global

import (
	"context"
	"goasmf/component"
)

var ComponentFactories map[string]func(context.Context) component.Component = make(map[string]func(context.Context) component.Component)

func RegisterComponent(name string, factoryFunction func(context.Context) component.Component) {
	ComponentFactories[name] = factoryFunction
}

func GetComponentFactoryByName(name string) func(ctx context.Context) component.Component {
	return ComponentFactories[name]
}

func CreateComponentByName(name string) component.Component {
	factory := GetComponentFactoryByName(name)
	if factory != nil {
		return factory(context.Background())
	} else {
		return nil
	}
}
