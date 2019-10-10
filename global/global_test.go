package global

import (
	"context"
	"goasmf/component"
	"reflect"
	"runtime"
	"testing"
)

type testComponent struct {
	component.BaseComponent
}

func (t testComponent) GetName() string {
	panic("implement me")
}

func componentFactory(context.Context) component.Component {
	return &testComponent{}
}

func TestGetComponentFactoryByName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want func(ctx context.Context) component.Component
		before func()
	}{
		{
			name: "should return nil if component name does not exists",
			args: args{name:"randomComponentName"},
			want: nil,
		}, {
			name:"should return component factory if name exists",
			args:args{name:"componentName"},
			want: componentFactory,
			before: func() {
				RegisterComponent("componentName", componentFactory)
			},
		},
	}
	for _, tt := range tests {
		if tt.before != nil {
			tt.before()
		}
		t.Run(tt.name, func(t *testing.T) {
			if got := GetComponentFactoryByName(tt.args.name); getFunctionName(got) != getFunctionName(tt.want) {
				t.Errorf("GetComponentFactoryByName() = %v, want %v", getFunctionName(got), getFunctionName(tt.want))
			}
		})
	}
}

func getFunctionName(i interface{}) string {
	if i == nil {
		return "nil"
	}
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func TestRegisterComponent(t *testing.T) {
	type args struct {
		name            string
		factoryFunction func(context.Context) component.Component
	}
	tests := []struct {
		name string
		args args
		verification func(t *testing.T)
	}{

		{
			name:         "should register component",
			args:         args{
				name:            "myComponent",
				factoryFunction: componentFactory,
			},
			verification: func(t *testing.T) {
				if comp := GetComponentFactoryByName("myComponent"); comp != nil {
					if getFunctionName(comp) != getFunctionName(componentFactory) {
						t.Errorf("TestRegisterComponent() = got %v, want %v", getFunctionName(comp), getFunctionName(componentFactory))
					}
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RegisterComponent(tt.args.name, tt.args.factoryFunction)
			tt.verification(t)
		})
	}
}