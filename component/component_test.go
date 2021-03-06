package component

import (
	"testing"
)

func TestRawComponent_GetJavascript(t *testing.T) {
	type fields struct {
		html       string
		javascript string
		css        string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "should return raw javascript",
			fields: fields{javascript: "<script>console.log('test')</script>"},
			want:   "<script>console.log('test')</script>",
		}, {
			name:   "should return empty javascript",
			fields: fields{},
			want:   "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := &RawComponent{
				Html:       tt.fields.html,
				Javascript: tt.fields.javascript,
				Css:        tt.fields.css,
			}
			if got := component.GetJavascript(); got != tt.want {
				t.Errorf("RawComponent.GetJavascript() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRawComponent_GetCSS(t *testing.T) {
	type fields struct {
		html       string
		javascript string
		css        string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "should get raw css",
			fields: fields{css: ".mycss{}"},
			want:   ".mycss{}",
		}, {
			name:   "should get empty css",
			fields: fields{},
			want:   "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := &RawComponent{
				Html:       tt.fields.html,
				Javascript: tt.fields.javascript,
				Css:        tt.fields.css,
			}
			if got := component.GetCSS(); got != tt.want {
				t.Errorf("RawComponent.GetCSS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRawComponent_GetHTML(t *testing.T) {
	type fields struct {
		html       string
		javascript string
		css        string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "should get raw html",
			fields: fields{html: "<div>Html</div>"},
			want:   "<div>Html</div>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			component := &RawComponent{
				Html:       tt.fields.html,
				Javascript: tt.fields.javascript,
				Css:        tt.fields.css,
			}
			if got := component.GetHTML(); got != tt.want {
				t.Errorf("RawComponent.GetHTML() = %v, want %v", got, tt.want)
			}
		})
	}
}
