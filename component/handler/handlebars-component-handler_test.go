package handler

import (
	"goasmf/component"
	"goasmf/rendering"
	"testing"
)

type TestFetcher struct {
}

func (fetcher *TestFetcher) FetchTemplate(templateName string) (string, error) {
	if templateName == "myCustomTemplate.hbm" {
		return "<div>This is my custom template {{name}} {{surname}}</div>", nil
	}
	return "<div>This is my template: {{name}} {{surname}}</div>", nil
}

type testHandlebarsComponent struct {
	component.BaseComponent
	Name    string
	Surname string
}

func (this *testHandlebarsComponent) GetName() string {
	return "testComp"
}

func getCorrectTestComponentWithModel() component.Component {

	comp := testHandlebarsComponent{
		Name:    "Tom",
		Surname: "Jones",
	}
	return &comp
}

type testHandlebarsComponent2 struct {
	component.BaseComponent
}

func (this *testHandlebarsComponent2) GetName() string {
	return "testComp"
}

func getCorrectTestComponentWithoutModel() component.Component {
	comp := &testHandlebarsComponent2{}
	return comp
}

type testHandlebarsComponentNotExportedFields struct {
	component.BaseComponent
	name    string
	surname string
}

func (this *testHandlebarsComponentNotExportedFields) GetName() string {
	return "testComp"
}

func getCorrectTestComponentWithoutExportedFields() component.Component {

	comp := &testHandlebarsComponentNotExportedFields{
		name:    "Name",
		surname: "Surname",
	}
	return comp
}

func TestHandlebarsComponentHandler_GetHTMLForDefaultTemplateFile(t *testing.T) {
	tests := []struct {
		name      string
		component component.Component
		want      string
	}{
		{
			name:      "should render html correctly",
			component: getCorrectTestComponentWithModel(),
			want:      "<div>This is my template: Tom Jones</div>",
		}, {
			name:      "should put blank placeholders if model is nil",
			component: getCorrectTestComponentWithoutModel(),
			want:      "<div>This is my template:  </div>",
		}, {
			name:      "should put blank placeholders if model fields are not exported",
			component: getCorrectTestComponentWithoutExportedFields(),
			want:      "<div>This is my template:  </div>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := tt.component
			handler := defaultHandlebarsComponentHandler{
				fetcher:  &TestFetcher{},
				renderer: rendering.NewHandlebarsRenderer(),
			}
			if got := handler.GetHTMLForDefaultTemplateFile(this); got != tt.want {
				t.Errorf("HandlebarsComponent.GetHTMLForDefaultTemplateFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandlebarsComponentHandler_GetHTMLForTemplateFile(t *testing.T) {
	tests := []struct {
		name      string
		component component.Component
		want      string
	}{
		{
			name:      "should render custom template html correctly",
			component: getCorrectTestComponentWithModel(),
			want:      "<div>This is my custom template Tom Jones</div>",
		}, {
			name:      "should put blank placeholders if model is nil",
			component: getCorrectTestComponentWithoutModel(),
			want:      "<div>This is my custom template  </div>",
		}, {
			name:      "should put blank placeholders if model fields are not exported",
			component: getCorrectTestComponentWithoutExportedFields(),
			want:      "<div>This is my custom template  </div>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := tt.component
			handler := defaultHandlebarsComponentHandler{
				fetcher:  &TestFetcher{},
				renderer: rendering.NewHandlebarsRenderer(),
			}
			if got := handler.GetHTMLForTemplateFile(this, "myCustomTemplate.hbm"); got != tt.want {
				t.Errorf("HandlebarsComponent.GetHTMLForTemplateFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
