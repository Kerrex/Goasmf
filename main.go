package main

import (
	"goasmf/component"
	"goasmf/component/handler"
	"reflect"

	"github.com/aymerick/raymond"
	"github.com/dennwc/dom"
	"github.com/iancoleman/strcase"
)

type TestComponent struct {
	TestVariable string
}

func (this *TestComponent) GetName() string {
	return "testComponent"
}

type TestComponent2 struct {
	TestVariable string
}

func (this *TestComponent2) GetName() string {
	return "testComponent2"
}

func main() {
	componentRepository := make(map[string]interface{})
	htmlHandler := handler.NewHtmlComponentHandler()
	println("starting my application!!")
	raymond.RegisterHelper("renderComponent", func(componentName string, options *raymond.Options) raymond.SafeString {
		componentToRender := componentRepository[componentName]
		if componentToRender == nil {
			panic("cannot render nonexistent component")
		}
		if comp, ok := componentToRender.(component.Component); ok {
			componentCopy := reflect.New(reflect.ValueOf(comp).Elem().Type()).Interface().(component.Component)
			inputParameters := options.Hash()
			for key, val := range inputParameters {
				ps := reflect.ValueOf(componentCopy)
				field := ps.Elem().FieldByName(strcase.ToCamel(key))
				if field.IsValid() && field.CanAddr() && field.CanSet() {
					field.Set(reflect.ValueOf(val))
				}
			}
			return raymond.SafeString(htmlHandler.GetHtml(componentCopy))
		}
		return raymond.SafeString("")
	})

	myComponent := &TestComponent{}
	myComponent.TestVariable = "dupadupa"

	myComponent2 := &TestComponent2{}
	myComponent2.TestVariable = "calkowicieinnykomponent"

	componentRepository["testComponent"] = myComponent
	componentRepository["testComponent2"] = myComponent2
	dom.Body.SetInnerHTML(htmlHandler.GetHtml(myComponent))

}
