package handler

import (
	"goasmf/component"
	"testing"
)

type TestHandlebarsComponentHandler struct {
}

func (this *TestHandlebarsComponentHandler) GetHTMLForTemplateFile(model component.Component, templateFileName string) string {
	return "myhtml-custom-template"
}

func (this *TestHandlebarsComponentHandler) GetHTMLForDefaultTemplateFile(model component.Component) string {
	return "myhtml-default"
}

func (this *TestHandlebarsComponentHandler) GetJavascript() string {
	return "myjavascript"
}

func (this *TestHandlebarsComponentHandler) GetCSS() string {
	return "mycss"
}

type testHandlebarsComponent3 struct {
}

type testHandlebarsComponent4 struct {
	testHandlebarsComponent3
}

func (this *testHandlebarsComponent3) GetName() string {
	return "my test component"
}

func (this *testHandlebarsComponent4) GetTemplateFileName() string {
	return "my custom test component"
}

func Test_defaultHtmlComponentHandler_GetHtml(t *testing.T) {
	type fields struct {
		handlebarsComponentHandler HandlebarsComponentHandler
	}
	type args struct {
		c component.Component
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "should render raw component",
			fields: fields{handlebarsComponentHandler: &TestHandlebarsComponentHandler{}},
			args:   args{c: &component.RawComponent{Html: "myhtml"}},
			want:   "myhtml",
		}, {
			name:   "should return default handlebars component",
			fields: fields{handlebarsComponentHandler: &TestHandlebarsComponentHandler{}},
			args:   args{c: &testHandlebarsComponent3{}},
			want:   "myhtml-default",
		}, {
			name:   "should return custom handlebars file component",
			fields: fields{handlebarsComponentHandler: &TestHandlebarsComponentHandler{}},
			args:   args{c: &testHandlebarsComponent4{}},
			want:   "myhtml-custom-template",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &defaultHtmlComponentHandler{
				handlebarsComponentHandler: tt.fields.handlebarsComponentHandler,
			}
			if got := this.GetHtml(tt.args.c); got != tt.want {
				t.Errorf("defaultHtmlComponentHandler.GetHtml() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_defaultHtmlComponentHandler_GetCSS(t *testing.T) {
// 	type fields struct {
// 		handlebarsComponentHandler *HandlebarsComponentHandler
// 	}
// 	type args struct {
// 		c component.Component
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 		want   string
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			this := &defaultHtmlComponentHandler{
// 				handlebarsComponentHandler: tt.fields.handlebarsComponentHandler,
// 			}
// 			if got := this.GetCSS(tt.args.c); got != tt.want {
// 				t.Errorf("defaultHtmlComponentHandler.GetCSS() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_defaultHtmlComponentHandler_GetJavascript(t *testing.T) {
// 	type fields struct {
// 		handlebarsComponentHandler *HandlebarsComponentHandler
// 	}
// 	type args struct {
// 		c component.Component
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 		want   string
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			this := &defaultHtmlComponentHandler{
// 				handlebarsComponentHandler: tt.fields.handlebarsComponentHandler,
// 			}
// 			if got := this.GetJavascript(tt.args.c); got != tt.want {
// 				t.Errorf("defaultHtmlComponentHandler.GetJavascript() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_defaultHtmlComponentHandler_getHandlebarsHtml(t *testing.T) {
// 	type fields struct {
// 		handlebarsComponentHandler *HandlebarsComponentHandler
// 	}
// 	type args struct {
// 		component component.HandlebarsComponent
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 		want   string
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			this := &defaultHtmlComponentHandler{
// 				handlebarsComponentHandler: tt.fields.handlebarsComponentHandler,
// 			}
// 			if got := this.getHandlebarsHtml(tt.args.component); got != tt.want {
// 				t.Errorf("defaultHtmlComponentHandler.getHandlebarsHtml() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_defaultHtmlComponentHandler_getCustomTemplateHandlebarsHtml(t *testing.T) {
// 	type fields struct {
// 		handlebarsComponentHandler *HandlebarsComponentHandler
// 	}
// 	type args struct {
// 		component component.CustomTemplateFileNameHandlebarsComponent
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 		want   string
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			this := &defaultHtmlComponentHandler{
// 				handlebarsComponentHandler: tt.fields.handlebarsComponentHandler,
// 			}
// 			if got := this.getCustomTemplateHandlebarsHtml(tt.args.component); got != tt.want {
// 				t.Errorf("defaultHtmlComponentHandler.getCustomTemplateHandlebarsHtml() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
