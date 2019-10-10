package component

// Component is default interface for any component type
type Component interface {
	GetName() string
	GetInstanceId() string
	SetInstanceId(componentInstanceId string)
}

// RawComponent is simple component providing raw HTML, jS and CSS without any preprocessing
type RawComponent struct {
	BaseComponent
	Name, Html, Javascript, Css string
}

// GetHTML renders HTML string
func (component *RawComponent) GetHTML() string {
	return component.Html
}

// GetJavascript renders JavaScript string
func (component *RawComponent) GetJavascript() string {
	return component.Javascript
}

// GetCSS renders Css String
func (component *RawComponent) GetCSS() string {
	return component.Css
}

func (component *RawComponent) GetName() string {
	return component.Name
}

type BaseComponent struct {
	instanceId string
}

func (b *BaseComponent) GetInstanceId() string {
	return b.instanceId
}

func (b *BaseComponent) SetInstanceId(newInstanceId string) {
	b.instanceId = newInstanceId
}



