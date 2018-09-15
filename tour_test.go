package tour

import (
	"image/color"
	"reflect"
	"strings"
	"testing"

	"fmt"

	"golang.org/x/tour/tree"
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
	type want struct {
		image Image
		model color.Model
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{"both are positive(1)", args{1, 2}, want{Image{2, 1}, color.RGBA64Model}},
		{"both are positive(2)", args{2, 1}, want{Image{1, 2}, color.RGBA64Model}},
		{"both are zero", args{0, 0}, want{Image{0, 0}, color.RGBA64Model}},
		{"x is zero", args{2, 0}, want{Image{0, 2}, color.RGBA64Model}},
		{"y is zero", args{0, 1}, want{Image{1, 0}, color.RGBA64Model}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Pic(tt.args.dx, tt.args.dy)
			if !reflect.DeepEqual(*got, tt.want.image) {
				t.Errorf("Pic() = %v, want %v", got, tt.want)
			}
			if got.ColorModel() != tt.want.model {
				t.Errorf("ColorModel() = %v, want %v", got.ColorModel(), tt.want.model)
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

func TestRot13Reader(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{"sample", "Lbh penpxrq gur pbqr!", "You cracked the code!"},
		{"japanese", "Lbh penpxrq gur 日本語1!", "You cracked the 日本語1!"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := strings.NewReader(tt.args)
			rot13 := &Rot13Reader{s}
			b := make([]byte, 1024)
			n, err := rot13.Read(b)
			if err != nil {
				t.Errorf("an error occurs: '%v'", err)
			}
			got := string(b[0:n])
			if got != tt.want {
				t.Errorf("Rot13Reader.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWalk(t *testing.T) {
	type args struct {
		t  *tree.Tree
		ch chan int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Tree1", args{tree.New(1), make(chan int, 1)}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{"Tree2", args{tree.New(2), make(chan int, 1)}, []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go Walk(tt.args.t, tt.args.ch)

			for v := range tt.args.ch {
				if v != tt.want[0] {
					t.Errorf("Walk() = %v, want %v", v, tt.want[0])
				}
				if len(tt.want) > 0 {
					tt.want = append(tt.want[:0], tt.want[1:]...)
				}
			}
		})
	}
}

func TestSame(t *testing.T) {
	type args struct {
		t1 *tree.Tree
		t2 *tree.Tree
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Tree1", args{tree.New(1), tree.New(1)}, true},
		{"Tree1", args{tree.New(3), tree.New(1)}, false},
		{"Tree1", args{tree.New(1), tree.New(3)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Same(tt.args.t1, tt.args.t2)
			if got != tt.want {
				t.Errorf("Same() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCrawl(t *testing.T) {
	sites := map[string]*fakeResult{
		"https://golang.org/": {
			"The Go Programming Language",
			[]string{
				"https://golang.org/pkg/",
				"https://golang.org/cmd/",
			},
		},
		"https://golang.org/pkg/": {
			"Packages",
			[]string{
				"https://golang.org/",
				"https://golang.org/cmd/",
				"https://golang.org/pkg/fmt/",
				"https://golang.org/pkg/os/",
			},
		},
		"https://golang.org/pkg/fmt/": {
			"Package fmt",
			[]string{
				"https://golang.org/",
				"https://golang.org/pkg/",
			},
		},
		"https://golang.org/pkg/os/": {
			"Package os",
			[]string{
				"https://golang.org/",
				"https://golang.org/pkg/",
			},
		},
	}
	fake := &fakeFetcher{sites, t, make(map[string]string)}
	Crawl("https://golang.org/", 4, fake)

	if 5 != len(fake.cache) {
		t.Errorf("numbers of urls are %v, want %v", len(fake.cache), 5)
	}
	for key := range sites {
		if _, ok := fake.cache[key]; !ok {
			t.Errorf("url %v is not fetched", key)
		}
	}
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher struct {
	result map[string]*fakeResult
	t      *testing.T
	cache  map[string]string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	f.cache[url] = url
	if res, ok := f.result[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

type fakeResult struct {
	body string
	urls []string
}
