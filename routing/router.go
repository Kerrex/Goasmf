package routing

import (
	"goasmf/component"
	"regexp"
	"strings"
)

type Router struct {
	mainRoute Route
}

func (this *Router) FindComponentByPath(path string) component.Component {
	normalizedPath := normalizePath(path)
	potentialRouteSegments := strings.Split(normalizedPath, "/")
	componentTree := make([]component.Component, 0)

	for routeSegment := range potentialRouteSegments {

	}

}

func findRouteWithRouteSegment(segment string, []Route routesToSearchIn) Route {

}

func normalizePath(pathToNormalize string) string {
	re := regexp.MustCompile("/{2,}")
	result := re.ReplaceAll([]byte(pathToNormalize), []byte("/"))
	resultStr := string(result)

	return strings.Trim(resultStr, "/")
}
