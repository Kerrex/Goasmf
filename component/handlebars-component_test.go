package component

import (
	"testing"
)

func TestHandlebarsComponent_GetTemplateFileName(t *testing.T) {
	type fields struct {
		name               string
		templateFileName   string
		javascriptFileName string
		cssFileName        string
		model              map[string]interface{}
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
		}, {
			name:      "should return templateFileName if is present",
			fields:    fields{templateFileName: "templateFile.hbm"},
			want:      "templateFile.hbm",
			wantPanic: false,
		}, {
			name:      "should panic if templateFileName has whitespaces",
			fields:    fields{templateFileName: "template file.hbm"},
			wantPanic: true,
		}, {
			name:      "should panic if templateFileName has incorrect extension",
			fields:    fields{templateFileName: "templateFileNoExt"},
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &HandlebarsComponent{
				name:               tt.fields.name,
				templateFileName:   tt.fields.templateFileName,
				javascriptFileName: tt.fields.javascriptFileName,
				cssFileName:        tt.fields.cssFileName,
				model:              tt.fields.model,
			}
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

func getCorrectTestComponentWithModel() Component {
	comp := NewHandlebarsComponent("testComp").(*HandlebarsComponent)
	comp.fetcher = &TestFetcher{}

	model := make(map[string]interface{})
	model["name"] = "Tom"
	model["surname"] = "Jones"
	comp.SetModel(model)
	return comp
}

func getCorrectTestComponentWithoutModel() Component {
	comp := NewHandlebarsComponent("testComp").(*HandlebarsComponent)
	comp.fetcher = &TestFetcher{}
	comp.SetModel(nil)
	return comp
}
func TestHandlebarsComponent_GetHTML(t *testing.T) {
	tests := []struct {
		name      string
		component Component
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
			if got := this.GetHTML(); got != tt.want {
				kjh
				t.Errorf("HandlebarsComponent.GetHTML() = %v, want %v", got, tt.want)
			}
		})
	}
}
