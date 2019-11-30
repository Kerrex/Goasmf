package routing

// Route interface for different route types, eg. different renderer, etc.
// should provide html, javascript and css ready to inject directly into HTML
type Route struct {
	Path string
	ComponentName string
	Children []Route
}
