package tour

import (
	"reflect"
	"testing"
)

func TestSqrt(t *testing.T) {
	type args struct {
		x float64
	}
	type want struct {
		x           float64
		expectError bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{"integer", args{4}, want{2, false}},
		{"double", args{2}, want{1.414213562373095, false}},
		{"zero", args{0}, want{0, false}},
		{"negative value", args{-1}, want{-1, true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Sqrt(tt.args.x)
			if got != tt.want.x {
				t.Errorf("Sqrt(%v) = %v, want %v", got, got, tt.want.x)
			}
			if (err != nil) != tt.want.expectError {
				t.Errorf("error is `%s`, but error is expected %v", err, tt.want.expectError)
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

func TestFibonacci(t *testing.T) {
	tests := []struct {
		name string
		in   int
		want []int
	}{
		{"10 times", 10, []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}},
	}
	for _, tt := range tests {
		f := Fibonacci()
		t.Run(tt.name, func(t *testing.T) {
			for i, v := range tt.want {
				if got := f(); got != v {
					t.Errorf("Fibonacci(%v) = %v, want %v", i, got, v)
				}
			}
		})
	}
}

func TestIPAddr(t *testing.T) {
	tests := []struct {
		name string
		args [4]byte
		want string
	}{
		{"googleDNS", [4]byte{8, 8, 8, 8}, "8.8.8.8"},
		{"loopBack", [4]byte{127, 0, 0, 1}, "localhost"},
		{"empty", [4]byte{}, "0.0.0.0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ip := IPAddr(tt.args)
			if got := ip.String(); got != tt.want {
				t.Errorf("IPAddr.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInfiniteA(t *testing.T) {
	type args struct {
		i int
	}
	type want struct {
		v byte
		i int
		e error
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{"100", args{100}, want{'A', 1, nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := InfiniteA{}
			for i := 0; i < tt.args.i; i++ {
				b := make([]byte, 1)
				got, err := r.Read(b)
				if b[0] != tt.want.v {
					t.Errorf("InfiniteA.Read() head = %v, want %v", got, tt.want.v)
				}

				if got != tt.want.i {
					t.Errorf("InfiniteA.Read() size = %v, want %v", got, tt.want.i)
				}

				if err != tt.want.e {
					t.Errorf("InfiniteA.Read() error = %v, wantErr %v", err, tt.want.e)
					return
				}
			}
		})
	}
}
