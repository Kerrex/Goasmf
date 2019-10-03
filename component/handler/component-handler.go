package handler

import (
	"goasmf/component"
	"goasmf/helpers"
	"goasmf/rendering"
)

type HtmlComponentHandler interface {
	GetHtml(c component.Component) string
	GetCSS(c component.Component) string
	GetJavascript(c component.Component) string
}

type defaultHtmlComponentHandler struct {
	handlebarsComponentHandler HandlebarsComponentHandler
}

func NewHtmlComponentHandler() HtmlComponentHandler {
	return &defaultHtmlComponentHandler{handlebarsComponentHandler: NewHandlebarsComponentHandler(helpers.NewFetcher(), rendering.NewHandlebarsRenderer())}
}

func (this *defaultHtmlComponentHandler) GetHtml(c component.Component) string {
	switch comp := c.(type) {
	default:
		println("Its actually unreachable")
		return ""
	case *component.RawComponent:
		return comp.GetHTML()
	case component.CustomTemplateFileNameHandlebarsComponent:
		return this.getCustomTemplateHandlebarsHtml(comp)
	case component.HandlebarsComponent:
		return this.getHandlebarsHtml(comp)
	}
}

func (this *defaultHtmlComponentHandler) GetCSS(c component.Component) string {
	panic("not implemented!")
}

func (this *defaultHtmlComponentHandler) GetJavascript(c component.Component) string {
	panic("not implemented!")
}

func (this *defaultHtmlComponentHandler) getHandlebarsHtml(component component.HandlebarsComponent) string {
	return this.handlebarsComponentHandler.GetHTMLForDefaultTemplateFile(component)
}

func (this *defaultHtmlComponentHandler) getCustomTemplateHandlebarsHtml(component component.CustomTemplateFileNameHandlebarsComponent) string {
	return this.handlebarsComponentHandler.GetHTMLForTemplateFile(component, component.GetTemplateFileName())
}
