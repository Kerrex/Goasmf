package main

import (
	"context"
	"github.com/dennwc/dom"
	"goasmf/component"
	"goasmf/component/handler"
	"goasmf/global"
	"goasmf/routing"
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

	htmlHandler := handler.NewHtmlComponentHandler()
	router := routing.Router{MainRoute: &routing.BaseRoute{
		Path:          "",
		ComponentName: "",
		Children: []routing.Route{
			&routing.BaseRoute{
				Path:          "test1",
				ComponentName: "testComponent",
				Children:      []routing.Route{
					&routing.BaseRoute{
						Path:          "tesst3",
						ComponentName: "testComponent2",
						Children:      nil,
					},
				},
			},
			&routing.BaseRoute{
				Path:          "test2",
				ComponentName: "testComponent2",
				Children:      nil,
			},
		},
	}}
	routingService := routing.Routing{Router: router}
	routingService.InitRoutingModule(htmlHandler)

	mainComponent := &MainComponent{}
	dom.Body.SetInnerHTML(htmlHandler.GetHtml(mainComponent))

}
