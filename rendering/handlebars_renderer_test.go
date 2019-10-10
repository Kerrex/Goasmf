package rendering

import (
	"goasmf/component"
	"testing"
)

func TestHandlebarsRenderer_Render(t *testing.T) {
	type args struct {
		hbm   string
		model component.Component
	}
	testContexts := []struct {
		Value interface{}
	}{
		{
			Value: "myValue",
		}, {
			Value: func() string {
				return "myValueFromFunc"
			},
		},
	}
	tests := []struct {
		name     string
		renderer Renderer
		args     args
		want     string
	}{
		{
			name:     "should render html without context",
			renderer: NewHandlebarsRenderer(),
			args:     args{hbm: "<div>it's html</div>"},
			want:     "<div>it's html</div>",
		}, {
			name:     "should return empty string if hbm is invalid",
			renderer: NewHandlebarsRenderer(),
			args:     args{hbm: "<div>{{value</div>"},
			want:     "",
		}, {
			name:     "should return html without variables if they are not in context",
			renderer: NewHandlebarsRenderer(),
			args:     args{hbm: "<div>{{value}}</div>"},
			want:     "<div></div>",
		}, {
			name:     "should return html with variables if it's string in context",
			renderer: NewHandlebarsRenderer(),
			args:     args{hbm: "<div>{{value}}</div>", model: testContexts[0]},
			want:     "<div>myValue</div>",
		}, {
			name:     "should return html with variables if it's function in context",
			renderer: NewHandlebarsRenderer(),
			args:     args{hbm: "<div>{{value}}</div>", model: testContexts[1]},
			want:     "<div>myValueFromFunc</div>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			renderer := &HandlebarsRenderer{}
			if got := renderer.Render(tt.args.hbm, tt.args.model); got != tt.want {
				t.Errorf("HandlebarsRenderer.Render() = %v, want %v", got, tt.want)
			}
		})
	}
}
