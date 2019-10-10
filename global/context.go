package global

import (
	"context"
	"goasmf/component"
)

const CurrentRoutingTree = "CurrentRoutingTree"
const CurrentlyRenderedComponentInstances = "CurrentlyRenderedComponentInstances"

var globalContext context.Context

func InitGlobalContext() {
	globalContext = context.WithValue(context.Background(), CurrentlyRenderedComponentInstances, make(map[string]component.Component))
}

func GetGlobalContext() context.Context {
	return globalContext
}

func SetGlobalContext(newContext context.Context) {
	globalContext = newContext
}


