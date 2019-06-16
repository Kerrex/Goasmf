package component

import (
	"errors"
	"goasmf/helpers"
	"goasmf/rendering"
	"strings"

	"github.com/iancoleman/strcase"
)

// HandlebarsComponent is model component rendered by handlebars
type HandlebarsComponent struct {
	name, templateFileName, javascriptFileName, cssFileName string
	model                                                   interface{}
	fetcher                                                 helpers.Fetcher
	renderer                                                rendering.Renderer
}

func NewHandlebarsComponent(name string) Component {
	return &HandlebarsComponent{name: name, fetcher: helpers.NewFetcher(), renderer: rendering.NewHandlebarsRenderer()}
}

// GetTemplateFileName gets valid template file name to fetch later
func (this *HandlebarsComponent) GetTemplateFileName() string {
	if this.templateFileName != "" {
		return getValidatedTemplateFileName(this.templateFileName)
	}

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

func (this *HandlebarsComponent) GetName() string {
	if strings.TrimSpace(this.name) == "" {
		panic("component name cannot be empty")
	}
	return this.name
}

func (this *HandlebarsComponent) GetHTML() string {
	templateFileName := this.GetTemplateFileName()
	template, err := this.fetcher.FetchTemplate(templateFileName)
	if err != nil {
		return ""
	}
	model := this.GetModel()
	return this.renderer.Render(template, model)
}

func (this *HandlebarsComponent) SetModel(model interface{}) {
	this.model = model
}

func (this *HandlebarsComponent) GetJavascript() string {
	panic("not implemented")
}

func (this *HandlebarsComponent) GetCSS() string {
	panic("not implemented")
}

func (this *HandlebarsComponent) GetModel() interface{} {
	return this.model
}
