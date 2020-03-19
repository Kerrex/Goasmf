package routing

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
