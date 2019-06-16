package component

// Component is default interface for any component type
// Should provide ready to render HTML, Javascript and CSS
type Component interface {
	GetName() string
	GetHTML() string
	GetJavascript() string
	GetCSS() string
}

// RawComponent is simple component providing raw HTML, jS and CSS without any preprocessing
type RawComponent struct {
	name, html, javascript, css string
}

// GetHTML renders HTML string
func (component *RawComponent) GetHTML() string {
	return component.html
}

// GetJavascript renders JavaScript string
func (component *RawComponent) GetJavascript() string {
	return component.javascript
}

// GetCSS renders Css String
func (component *RawComponent) GetCSS() string {
	return component.css
}

func (component *RawComponent) getName() string {
	return component.name
}

// ModelComponent is Component containing model used to render HTML
type ModelComponent interface {
	Component
	GetModel() interface{}
	SetModel(model interface{})
}
