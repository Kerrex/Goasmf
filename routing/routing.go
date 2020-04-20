package routing

import (
	"context"
	"fmt"
	"github.com/aymerick/raymond"
	"github.com/dennwc/dom"
	"github.com/dennwc/dom/js"
	"github.com/iancoleman/strcase"
	"goasmf/component"
	"goasmf/component/handler"
	"goasmf/global"
	"goasmf/rendering"
	"reflect"
)

type Routing struct {
	Router Router
}

func (this *Routing) InitRoutingModule(handler handler.HtmlComponentHandler) {
	raymond.RegisterHelper("renderComponent", func(componentName string, options *raymond.Options) raymond.SafeString {
		componentInstanceId := options.HashStr("componentInstanceId")
		knownComponentInstances := global.GetGlobalContext().Value(global.CurrentlyRenderedComponentInstances).(map[string]component.Component)
		if comp, ok := knownComponentInstances[componentInstanceId]; ok {
			inputParameters := options.Hash()
			for key, val := range inputParameters {
				ps := reflect.ValueOf(comp)
				field := ps.Elem().FieldByName(strcase.ToCamel(key))
				if field.IsValid() && field.CanAddr() && field.CanSet() {
					field.Set(reflect.ValueOf(val))
				}
			}
			return raymond.SafeString(fmt.Sprintf("<div componentInstanceId=\"%s\">%s</div>", componentInstanceId, handler.GetHtml(comp)))
		}

		println("Component not registered:", componentName)
		println(len(knownComponentInstances))
		for key, val := range knownComponentInstances {
			println("key: " + key + " value " + val.GetName())
		}
		println(global.GetGlobalContext().Value(global.CurrentlyRenderedComponentInstances))
		return raymond.SafeString("")
	})

	raymond.RegisterHelper("routerOutlet", func(options *raymond.Options) raymond.SafeString {
		renderer := rendering.NewHandlebarsRenderer()
		renderedRoute := this.getRouteToRender()
		renderedString := renderer.Render(fmt.Sprintf("{{renderComponent \"%s\"}}", renderedRoute.GetComponent().GetName()), renderedRoute.GetComponent())

		return raymond.SafeString(renderedString)
	})
}

func (this *Routing) getRouteToRender() RenderedRoute {
	renderedRoute := this.getRenderedRoute()

	for renderedRoute.GetRoute().GetRedirectPath() != "" {
		pathToRedirectTo := renderedRoute.GetRoute().GetRedirectPath()
		js.Get("history").Call("pushState", nil, nil, pathToRedirectTo)
		renderedRoute = this.getRenderedRoute()
	}

	return renderedRoute
}

func (this *Routing) getRenderedRoute() RenderedRoute {
	currentUrl := getCurrentPath()
	renderedRoute := this.Router.FindRenderedRouteByPath(currentUrl)
	return renderedRoute
}

func getCurrentPath() string {
	return dom.GetWindow().JSValue().Get("location").Get("pathname").String()
}

func getRenderComponentContext(options *raymond.Options) context.Context {
	contextFromParam := options.Param(0)
	if contextFromParam == nil {
		return context.Background()
	}

	if ctx, ok := contextFromParam.(context.Context); ok {
		return ctx
	}

	return context.Background()
}
