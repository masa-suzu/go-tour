package tour

import (
	"reflect"
	"testing"
)

func TestSqrt(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"integer", args{4}, 2},
		{"double", args{2}, 1.414213562373095},
		{"zero", args{0}, 0},
		{"negative value", args{-1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sqrt(tt.args.x); got != tt.want {
				t.Errorf("Sqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPic(t *testing.T) {
	type args struct {
		dx int
		dy int
	}
	tests := []struct {
		name string
		args args
		want [][]uint8
	}{
		{"both are positive(1)", args{1, 2}, [][]uint8{{0}, {0}}},
		{"both are positive(2)", args{2, 1}, [][]uint8{{0, 0}}},
		{"both are zero", args{0, 0}, [][]uint8{}},
		{"x is zero", args{0, 0}, [][]uint8{}},
		{"y is zero", args{0, 1}, [][]uint8{{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Pic(tt.args.dx, tt.args.dy); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pic() = %v, want %v", got, tt.want)
			}
		})
	}
}
