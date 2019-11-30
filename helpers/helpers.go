package helpers

import (
	"io/ioutil"
	"net/http"
)

const templateDir = "/templates/"
var rawComponentTemplateCache = make(map[string]string)

type Fetcher interface {
	FetchTemplate(templateName string) (string, error)
}

type DefaultFetcher struct {
}

// FetchTemplate fetches template from remote server
func (f *DefaultFetcher) FetchTemplate(templateName string) (string, error) {
	if template, ok := rawComponentTemplateCache[templateName]; ok {
		return template, nil
	}

	request, err := http.Get(templateDir + templateName)
	if err != nil {
		return "", err
	}
	bodyBytes, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return "", err
	}

	rawComponentTemplateCache[templateName] = string(bodyBytes)
	return string(bodyBytes), nil
}

func NewFetcher() Fetcher {
	return &DefaultFetcher{}
}