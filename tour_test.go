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

func TestWordCount(t *testing.T) {
	tests := []struct {
		in   string
		want map[string]int
	}{
		{"I am learning Go!", map[string]int{
			"I": 1, "am": 1, "learning": 1, "Go!": 1,
		}},
		{"The quick brown fox jumped over the lazy dog.", map[string]int{
			"The": 1, "quick": 1, "brown": 1, "fox": 1, "jumped": 1,
			"over": 1, "the": 1, "lazy": 1, "dog.": 1,
		}},
		{"I ate a donut. Then I ate another donut.", map[string]int{
			"I": 2, "ate": 2, "a": 1, "donut.": 2, "Then": 1, "another": 1,
		}},
		{"A man a plan a canal panama.", map[string]int{
			"A": 1, "man": 1, "a": 2, "plan": 1, "canal": 1, "panama.": 1,
		}},
	}

	for _, c := range tests {
		ok := true
		got := WordCount(c.in)
		if len(c.want) != len(got) {
			ok = false
		} else {
			for k := range c.want {
				if c.want[k] != got[k] {
					ok = false
				}
			}
		}
		if !ok {
			t.Errorf("WordCount(%q) = %#v, want:\n  %#v",
				c.in, got, c.want)
			break
		}
	}
}
