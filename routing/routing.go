package routing

import (
	"context"
	"fmt"
	"goasmf/component"
	"goasmf/component/handler"
	"goasmf/global"
	"reflect"

	"github.com/aymerick/raymond"
	"github.com/iancoleman/strcase"
)

type Routing struct {
}

func InitRoutingModule(handler handler.HtmlComponentHandler) {
	raymond.RegisterHelper("renderComponent", func(componentName string, options *raymond.Options) raymond.SafeString {
		componentInstanceId := options.HashStr("componentInstanceId")
		knownComponentInstances := global.GetGlobalContext().Value(global.CurrentlyRenderedComponentInstances).(map[string]component.Component)
		if comp, ok := knownComponentInstances[componentInstanceId]; ok {
			componentCopy := reflect.New(reflect.ValueOf(comp).Elem().Type()).Interface().(component.Component)
			inputParameters := options.Hash()
			for key, val := range inputParameters {
				ps := reflect.ValueOf(componentCopy)
				field := ps.Elem().FieldByName(strcase.ToCamel(key))
				if field.IsValid() && field.CanAddr() && field.CanSet() {
					field.Set(reflect.ValueOf(val))
				}
			}
			return raymond.SafeString(fmt.Sprintf("<div componentInstanceId=\"%s\">%s</div>", componentInstanceId, handler.GetHtml(componentCopy)))
		}
		return ""
	})

	raymond.RegisterHelper("routerOutlet", func(options *raymond.Options) raymond.SafeString {
		mainComponent := global.ComponentFactories["testComponent"](context.Background())
		return raymond.SafeString(handler.GetHtml(mainComponent))
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
