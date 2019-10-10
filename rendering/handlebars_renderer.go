package rendering

import (
	"context"
	"fmt"
	"github.com/aymerick/raymond/ast"
	"goasmf/component"
	"reflect"
	"unsafe"

	"github.com/aymerick/raymond"
)

var instanceTemplateCache = make(map[string]*raymond.Template)

type Renderer interface {
	Render(template string, comp component.Component) string
}
type HandlebarsRenderer struct {
}

func (renderer *HandlebarsRenderer) Render(hbm string, comp component.Component) string {
	if comp == nil {
		panic("cannot render component with no component!")
	}

	if template, ok := instanceTemplateCache[comp.GetInstanceId()]; ok {
		html, err :=  template.Exec(comp)
		if err != nil {
			panic(err)
		}
		return html
	}
	parsed, err := raymond.Parse(hbm)
	if parsed == nil {
		panic("invalid hbm file: " + hbm)
	}
	generateInstanceIdsAndSaveToGlobalContext(parsed)

	println(parsed.PrintAST())

	html, err := parsed.Exec(comp)
	println("html " + html)
	if err != nil {
		fmt.Errorf("Unable to parse template: %s", err.Error())
		return ""
	}
	return html
}

func generateInstanceIdsAndSaveToGlobalContext(parsed *raymond.Template) {
	program := getAstProgram(parsed)
	visitor := NewRenderComponentVisitor(context.Background())
	program.Accept(visitor)
}

func getAstProgram(parsed *raymond.Template) *ast.Program {
	rs := reflect.ValueOf(parsed).Elem()
	rf := rs.FieldByName("program")
	rf = reflect.NewAt(rf.Type(), unsafe.Pointer(rf.UnsafeAddr())).Elem() // Strange hack to get unexported field
	program := rf.Interface().(*ast.Program)
	return program
}

func NewHandlebarsRenderer() Renderer {
	return &HandlebarsRenderer{}
}
