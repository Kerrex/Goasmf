package main

import (
	"context"
	"goasmf/component"
	"goasmf/component/handler"
	"goasmf/global"
	"goasmf/routing"

	"github.com/dennwc/dom"
)

type TestComponent struct {
	component.BaseComponent
	TestVariable string
}

func (this *TestComponent) GetName() string {
	return "testComponent"
}

type TestComponent2 struct {
	component.BaseComponent
	TestVariable string
}

func (this *TestComponent2) GetName() string {
	return "testComponent2"
}

type MainComponent struct {
	component.BaseComponent
}

func (this *MainComponent) GetTemplateFileName() string {
	return "index.hbm"
}

func (this *MainComponent) GetName() string {
	return "main"
}

func main() {
	println("starting my application!!")
	global.InitGlobalContext()
	htmlHandler := handler.NewHtmlComponentHandler()
	routing.InitRoutingModule(htmlHandler)

	global.RegisterComponent("testComponent", func(ctx context.Context) component.Component {
		myComponent := &TestComponent{}
		myComponent.TestVariable = "dupadupa"
		return myComponent
	})

	global.RegisterComponent("testComponent2", func(ctx context.Context) component.Component {
		myComponent := &TestComponent2{}
		myComponent.TestVariable = "calkowicieinnykomponent"
		return myComponent
	})

	mainComponent := &MainComponent{}
	dom.Body.SetInnerHTML(htmlHandler.GetHtml(mainComponent))

}
