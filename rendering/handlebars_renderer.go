package rendering

import (
	"fmt"

	"github.com/aymerick/raymond"
)

type Renderer interface {
	Render(template string, model interface{}) string
}
type HandlebarsRenderer struct {
}

func (renderer *HandlebarsRenderer) Render(hbm string, model interface{}) string {
	html, err := raymond.Render(hbm, model)
	if err != nil {
		fmt.Errorf("Unable to parse template: %s", err.Error())
		return ""
	}
	return html
}

func NewHandlebarsRenderer() Renderer {
	return &HandlebarsRenderer{}
}
