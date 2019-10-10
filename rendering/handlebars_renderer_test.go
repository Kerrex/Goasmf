package rendering

import (
	"goasmf/component"
	"testing"
)

type testComponent struct {
	component.BaseComponent
	Value string
}

func (t *testComponent) GetName() string {
	return ""
}

type testComponent2 struct {
	component.BaseComponent
}

func (t *testComponent2) GetName() string {
	return ""
}

func (t *testComponent2) Value() string {
	return "myValueFromFunc"
}

type args struct {
	hbm   string
	model component.Component
}

func TestHandlebarsRenderer_Render(t *testing.T) {
	testComponent := testComponent{Value: "myValue"}
	tests := []struct {
		name        string
		renderer    Renderer
		args        args
		want        string
		shouldPanic bool
	}{
		{
			name:        "should panic if context absent",
			renderer:    NewHandlebarsRenderer(),
			args:        args{hbm: "<div>it's html</div>"},
			want:        "<div>it's html</div>",
			shouldPanic: true,
		}, {
			name:        "should panic if hbm is invalid",
			renderer:    NewHandlebarsRenderer(),
			args:        args{hbm: "<div>{{value</div>", model: &testComponent2{}},
			want:        "",
			shouldPanic: true,
		}, {
			name:     "should return html without variables if they are not in context",
			renderer: NewHandlebarsRenderer(),
			args:     args{hbm: "<div>{{someValue}}</div>", model: &testComponent2{}},
			want:     "<div></div>",
		}, {
			name:     "should return html with variables if it's string in context",
			renderer: NewHandlebarsRenderer(),
			args:     args{hbm: "<div>{{value}}</div>", model: &testComponent},
			want:     "<div>myValue</div>",
		}, {
			name:     "should return html with variables if it's function in context",
			renderer: NewHandlebarsRenderer(),
			args:     args{hbm: "<div>{{value}}</div>", model: &testComponent2{}},
			want:     "<div>myValueFromFunc</div>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			renderer := &HandlebarsRenderer{}
			doTest(renderer, tt, t)
		})
	}
}

func doTest(renderer *HandlebarsRenderer, tt struct {
	name        string;
	renderer    Renderer;
	args        args;
	want        string;
	shouldPanic bool
}, t *testing.T) {
	if tt.shouldPanic {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("HandlebarsRenderer.Render want panic but didnt happen")
			}
		}()
	}
	if got := renderer.Render(tt.args.hbm, tt.args.model); got != tt.want {
		t.Errorf("HandlebarsRenderer.Render() = %v, want %v", got, tt.want)
	}
}
