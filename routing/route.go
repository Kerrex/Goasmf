package routing

// Route interface for different route types, eg. different renderer, etc.
// should provide html, javascript and css ready to inject directly into HTML
type Route interface {
	GetHtml() string
	GetJavascript() string
	GetCss() string
	GetRoute() string
}

type HandlebarsRoute struct {
}
