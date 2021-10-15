package simplefactory

import (
	"testing"
)

func TestRun(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				name: "product1",
			},
			want: "this is product1",
		},
		{
			name: "test2",
			args: args{
				name: "product2",
			},
			want: "this is product2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Run(tt.args.name); got != tt.want {
				t.Errorf("Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
