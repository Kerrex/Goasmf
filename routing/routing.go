package routing

import (
	"context"
	"fmt"
	"github.com/aymerick/raymond"
	"github.com/iancoleman/strcase"
	"goasmf/component"
	"goasmf/component/handler"
	"goasmf/global"
	"goasmf/rendering"
	"reflect"
)

type Routing struct {
}

func InitRoutingModule(handler handler.HtmlComponentHandler) {
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
		mainComponent := global.ComponentFactories["testComponent"](context.Background())
		renderedString := renderer.Render("{{renderComponent \"testComponent\"}}", mainComponent)
		return raymond.SafeString(renderedString)
	})
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
