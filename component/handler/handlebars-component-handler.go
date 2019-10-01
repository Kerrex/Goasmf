package handler

import (
	"errors"
	"goasmf/helpers"
	"goasmf/rendering"
	"strings"

	"github.com/iancoleman/strcase"
)

type HandlebarsComponentHandler struct {
	name     string
	fetcher  helpers.Fetcher
	renderer rendering.Renderer
}

func NewHandlebarsComponentHandler(name string) *HandlebarsComponentHandler {
	return &HandlebarsComponentHandler{name: name, fetcher: helpers.NewFetcher(), renderer: rendering.NewHandlebarsRenderer()}
}

// GetTemplateFileName gets valid template file name to fetch later
func (this *HandlebarsComponentHandler) GetTemplateFileName() string {
	return strcase.ToCamel(this.GetName()) + ".hbm"
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

func (this *HandlebarsComponentHandler) GetName() string {
	if strings.TrimSpace(this.name) == "" {
		panic("component name cannot be empty")
	}
	return this.name
}

func (this *HandlebarsComponentHandler) GetHTML(model interface{}) string {
	templateFileName := this.GetTemplateFileName()
	template, err := this.fetcher.FetchTemplate(templateFileName)
	if err != nil {
		return ""
	}

	return this.renderer.Render(template, model)
}

func (this *HandlebarsComponentHandler) GetJavascript() string {
	panic("not implemented")
}

func (this *HandlebarsComponentHandler) GetCSS() string {
	panic("not implemented")
}
