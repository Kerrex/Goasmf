package routing

import (
	"context"
	"goasmf/global"
	"regexp"
	"strings"
)

type Router struct {
	MainRoute Route
}

func (this *Router) FindRenderedRouteByPath(path string) RenderedRoute {
	normalizedPath := normalizePath(path)
	potentialRouteSegments := strings.Split(normalizedPath, "/")
	renderingTree := make([]RenderedRoute, 0)
	println(len(potentialRouteSegments))

	routeToSearchInChildren := this.MainRoute
	if normalizedPath == this.MainRoute.GetPath() {
		renderingTree = append(renderingTree, getRenderedRoute(this.MainRoute))
	} else {
		for _, routeSegment := range potentialRouteSegments {
			route := findRouteWithRouteSegment(routeSegment, routeToSearchInChildren.GetChildren())
			if route != nil {
				renderedRoute := getRenderedRoute(route)
				renderingTree = append(renderingTree, renderedRoute)

				routeToSearchInChildren = renderedRoute.route
			} else {
				break
			}
		}
	}

	globalContext := global.GetGlobalContext()
	global.SetGlobalContext(context.WithValue(globalContext, global.CurrentRoutingTree, renderingTree))
	if len(renderingTree) == 0 {
		panic("unknown route " + path)
	}
	return renderingTree[(len(renderingTree) - 1)]
}

func getRenderedRoute(route Route) BaseRenderedRoute {
	return BaseRenderedRoute{route: route, component: global.CreateComponentByName(route.GetComponentName())}
}

func findRouteWithRouteSegment(segment string, routesToSearchIn []Route) Route {
	if routesToSearchIn == nil {
		return nil
	}

	for _, route := range routesToSearchIn {
		println(route.GetPath())
		if strings.ToUpper(segment) == strings.ToUpper(route.GetPath()) {
			return route
		}
	}
	return nil
}

func normalizePath(pathToNormalize string) string {
	re := regexp.MustCompile("/{2,}")
	result := re.ReplaceAll([]byte(pathToNormalize), []byte("/"))
	resultStr := string(result)

	return strings.Trim(resultStr, "/")
}
