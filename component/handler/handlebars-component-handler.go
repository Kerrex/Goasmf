package handler

import (
	"errors"
	"goasmf/component"
	"goasmf/helpers"
	"goasmf/rendering"
	"strings"

	"github.com/iancoleman/strcase"
)

type HandlebarsComponentHandler interface {
	GetHTMLForTemplateFile(model component.Component, templateFileName string) string
	GetHTMLForDefaultTemplateFile(model component.Component) string
	GetJavascript() string
	GetCSS() string
}

type defaultHandlebarsComponentHandler struct {
	fetcher  helpers.Fetcher
	renderer rendering.Renderer
}

func NewHandlebarsComponentHandler(fetcher helpers.Fetcher, renderer rendering.Renderer) HandlebarsComponentHandler {
	return &defaultHandlebarsComponentHandler{fetcher: fetcher, renderer: renderer}
}

func NewCustomTemplateHandlebarsComponentHandler(templateFileName string, fetcher helpers.Fetcher, renderer rendering.Renderer) HandlebarsComponentHandler {
	return &defaultHandlebarsComponentHandler{fetcher: fetcher, renderer: renderer}
}

// GetTemplateFileName gets valid template file name to fetch later
func getTemplateFileName(name string) string {
	return strcase.ToCamel(name) + ".hbm"
}

func getValidatedTemplateFileName(fileName string) string {
	err := validateTemplateFileName(fileName)
	if err != nil {
		panic(err)
	}
	return fileName
}

func validateTemplateFileName(fileName string) error {
	if !strings.HasSuffix(fileName, ".hbm") {
		return errors.New("handlebars template file must end with '.hbm'")
	}
	return validateFileName(fileName)
}

func validateFileName(fileName string) error {
	if strings.ContainsAny(fileName, " \n\t\r") {
		return errors.New("filename can't contain whitespaces")
	}
	return nil
}

func (this *defaultHandlebarsComponentHandler) GetHTMLForTemplateFile(model component.Component, templateFileName string) string {
	err := validateTemplateFileName(templateFileName)
	if err != nil {
		panic(err)
	}

	template, err := this.fetcher.FetchTemplate(templateFileName)
	if err != nil {
		return ""
	}

	return this.renderer.Render(template, model)
}

func (this *defaultHandlebarsComponentHandler) GetHTMLForDefaultTemplateFile(model component.Component) string {
	templateFileName := getTemplateFileName(model.GetName())
	return this.GetHTMLForTemplateFile(model, templateFileName)
}

func (this *defaultHandlebarsComponentHandler) GetJavascript() string {
	panic("not implemented")
}

func (this *defaultHandlebarsComponentHandler) GetCSS() string {
	panic("not implemented")
}
