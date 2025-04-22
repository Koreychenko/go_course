package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

var urlList *UrlList

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	if depth <= 0 {
		return
	}

	if urlList.Crawled(url) {
		fmt.Println("already crawled, exiting", url)

		return
	}

	body, urls, err := fetcher.Fetch(url)

	urlList.Add(url, body)

	ch <- url

	if err != nil {
		urlList.Add(url, err.Error())

		return
	}

	for _, u := range urls {
		wg.Add(1)
		go Crawl(u, depth-1, fetcher, ch, wg)
	}
	return
}

func main() {
	urlList = NewUrlList()
	ch := make(chan string)
	var wg sync.WaitGroup

	wg.Add(1)
	go Crawl("https://golang.org/", 4, fetcher, ch, &wg)

	go func() {
		wg.Wait()
		close(ch)
	}()

	for url := range ch {
		fmt.Println(url)
	}
}

type UrlList struct {
	mutex sync.Mutex
	urls  map[string]string
}

func (l *UrlList) Add(url string, body string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if _, ok := l.urls[url]; !ok {
		l.urls[url] = body
	}
}

func (l *UrlList) Crawled(url string) bool {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	_, ok := l.urls[url]

	return ok
}

func NewUrlList() *UrlList {
	return &UrlList{urls: make(map[string]string)}
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
