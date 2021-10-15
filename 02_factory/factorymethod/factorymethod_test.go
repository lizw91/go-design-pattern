package factorymethod

import (
	"testing"
)

func TestRun(t *testing.T) {
	type args struct {
		bunShop string
		bunType string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				bunShop: "QS",
				bunType: "pig",
			},
			want: "this is QSPigMeatBuns",
		},
		{
			name: "test2",
			args: args{
				bunShop: "QS",
				bunType: "3s",
			},
			want: "this is QSSamSunStuffingBuns",
		},
		{
			name: "test3",
			args: args{
				bunShop: "GD",
				bunType: "pig",
			},
			want: "this is GDPigMeatBuns",
		},
		{
			name: "test4",
			args: args{
				bunShop: "GD",
				bunType: "3s",
			},
			want: "this is GDSamSunStuffingBuns",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Run(tt.args.bunShop, tt.args.bunType); got != tt.want {
				t.Errorf("Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
