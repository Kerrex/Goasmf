package handler

import (
	"goasmf/component"
	"testing"
)

func TestHandlebarsComponentHandler_GetTemplateFileName(t *testing.T) {
	type fields struct {
		name string
	}
	tests := []struct {
		name      string
		fields    fields
		want      string
		wantPanic bool
	}{
		{
			name:      "should return camelcased name with extension if templateFileName is empty",
			fields:    fields{name: "Test HandleBars Component"},
			want:      "TestHandleBarsComponent.hbm",
			wantPanic: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewHandlebarsComponentHandler(tt.fields.name)
			if tt.wantPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("HandlebarsComponent.GetTemplateFileName() - expected panic")
					}
				}()
			}
			got := this.GetTemplateFileName()
			if got != tt.want && !tt.wantPanic {
				t.Errorf("HandlebarsComponent.GetTemplateFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}

type TestFetcher struct {
}

func (fetcher *TestFetcher) FetchTemplate(templateName string) (string, error) {
	return "<div>This is my template: {{name}} {{surname}}</div>", nil
}

type testHandlebarsComponent struct {
	name    string
	surname string
}

func (this *testHandlebarsComponent) GetName() string {
	return "testComp"
}

func getCorrectTestComponentWithModel() component.Component {

	comp := testHandlebarsComponent{
		name:    "Tom",
		surname: "Jones",
	}
	return &comp
}

type testHandlebarsComponent2 struct {
}

func (this *testHandlebarsComponent2) GetName() string {
	return "testComp"
}

func getCorrectTestComponentWithoutModel() component.Component {
	comp := &testHandlebarsComponent2{}
	return comp
}
func TestHandlebarsComponentHandler_GetHTML(t *testing.T) {
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
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := tt.component
			handler := NewHandlebarsComponentHandler(this.GetName())
			if got := handler.GetHTML(this); got != tt.want {
				t.Errorf("HandlebarsComponent.GetHTML() = %v, want %v", got, tt.want)
			}
		})
	}
}
