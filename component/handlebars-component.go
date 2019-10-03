package component

// HandlebarsComponent is model component rendered by handlebars
type HandlebarsComponent interface {
	Component
}

type CustomTemplateFileNameHandlebarsComponent interface {
	HandlebarsComponent
	GetTemplateFileName() string
}
