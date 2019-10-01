package handler

import (
	"goasmf/component"
)

type HtmlComponentHandler interface {
	GetHtml(c component.Component) string
	GetCSS(c component.Component) string
	GetJavascript(c component.Component) string
}

type defaultHtmlComponentHandler struct {
}

func NewHtmlComponentHandler() HtmlComponentHandler {
	return &defaultHtmlComponentHandler{}
}

func (this *defaultHtmlComponentHandler) GetHtml(c component.Component) string {
	switch comp := c.(type) {
	default:
		println("Error!")
		return ""
	case component.HandlebarsComponent:
		return getHandlebarsHtml(comp)
	case *component.RawComponent:
		return comp.GetHTML()

	}
}

func (this *defaultHtmlComponentHandler) GetCSS(c component.Component) string {
	panic("not implemented!")
}

func (this *defaultHtmlComponentHandler) GetJavascript(c component.Component) string {
	panic("not implemented!")
}

func getHandlebarsHtml(component component.HandlebarsComponent) string {
	handler := NewHandlebarsComponentHandler(component.GetName())
	return handler.GetHTML(component)
}
