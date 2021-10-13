package simple_factory

import (
	"reflect"
	"testing"
)

func Test_jsonRuleConfigParser_Parse(t *testing.T) {
	tests := []struct {
		name string
		j    *jsonRuleConfigParser
	}{
		{
			name: "json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &jsonRuleConfigParser{}
			j.Parse()
		})
	}
}

func Test_yamlRuleConfigParser_Parse(t *testing.T) {
	tests := []struct {
		name string
		y    *yamlRuleConfigParser
	}{
		{
			name: "yaml",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			y := &yamlRuleConfigParser{}
			y.Parse()
		})
	}
}

func TestNewRuleConfigParser(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want IRuleConfigParser
	}{
		{
			name: "json",
			args: args{t: "json"},
			want: &jsonRuleConfigParser{},
		},

		{
			name: "yaml",
			args: args{t: "yaml"},
			want: &yamlRuleConfigParser{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRuleConfigParser(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRuleConfigParser() = %v, want %v", got, tt.want)
			}
		})
	}
}
