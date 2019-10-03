package global

import "context"

const CurrentRoutingTree = "CurrentRoutingTree"

var globalContext = context.Background()

func GetGlobalContext() context.Context {
	return globalContext
}

func SetGlobalContext(newContext context.Context) {
	globalContext = newContext
}


