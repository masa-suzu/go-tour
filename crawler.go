package tour

import "fmt"

type fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type message struct {
	message string
}

type uRL struct {
	url   string
	depth int
}
type quit struct{}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher fetcher) {
	ch := make(chan interface{}, 20)

	workers := 0
	visited := make(map[string]bool)

	ch <- uRL{url, depth}

	for {
		req := <-ch

		switch req := req.(type) {
		case message:
			fmt.Print(req.message)
		case uRL:
			if req.depth > 0 && !visited[req.url] {
				visited[req.url] = true
				workers++

				go fetch(fetcher, req.url, req.depth, ch)
			}
		case quit:
			workers--
		}

		if workers <= 0 {
			break
		}
	}
}

func fetch(f fetcher, url string, depth int, ch chan interface{}) {
	defer func() { ch <- quit{} }()

	body, urls, err := f.Fetch(url)

	if err != nil {
		ch <- message{fmt.Sprintf("%s\n", err)}
		return
	}

	ch <- message{fmt.Sprintf("found: %s %q\n", url, body)}

	for _, u := range urls {
		ch <- uRL{u, depth - 1}
	}
}
