package main

import (
	"fmt"
	"sync"
	"time"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type SafeMap struct {
	m   map[string]bool
	mux sync.Mutex
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, 
				urlMap *SafeMap, done chan bool) {

	if depth <= 0 {
		return
	}
	

	urlMap.mux.Lock()
	if urlMap.m[url] {
		//fmt.Printf("skipping, already crawled %s\n", url)
		urlMap.mux.Unlock()		
		return
	}
	urlMap.mux.Unlock()
	
	body, urls, err := fetcher.Fetch(url)
	
	urlMap.mux.Lock()
	urlMap.m[url] = true
	urlMap.mux.Unlock()
	
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	doneChildren := make([]chan bool, len(urls))	
	for	i, u := range urls {
		doneChild := make(chan bool)
		doneChildren[i] = doneChild
		go Crawl(u, depth-1, fetcher, urlMap, doneChild)
	}
	for _, doneChild := range doneChildren {
		select {
		case <-doneChild:
			break
		default:
			time.Sleep(50 * time.Millisecond)
		}
	}
	done <- true
}

func main() {
	urlMap := SafeMap{m: make(map[string]bool)}
	done := make(chan bool)
	go Crawl("http://golang.org/", 4, fetcher, &urlMap, done)
	<-done
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
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
