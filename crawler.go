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

type Cache struct {
    ca  map[string]bool
    mux sync.Mutex
}

func Crawler(url string, depth int, fetcher Fetcher, cache Cache) {
    cache.mux.Lock()
    _, ok := cache.ca[url]
    if !ok {
        cache.ca[url] = true
    }
    cache.mux.Unlock()

    if ok {
        return
    }

    if depth <= 0 {
        return
    }
    body, urls, err := fetcher.Fetch(url)
    if err != nil {
        return
    }

    fmt.Printf("found: %s %q\n", url, body)
    for _, u := range urls {
        Crawler(u, depth-1, fetcher, cache)
    }
    return
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
    // TODO: Fetch URLs in parallel.
    // TODO: Don't fetch the same URL twice.
    // This implementation doesn't do either:
    cache := Cache{ca: make(map[string]bool)}
    _, urls, err := fetcher.Fetch(url)
    if err != nil {
        return
    }
    for _, u := range urls {
        go Crawler(u, depth-1, fetcher, cache)
    }
}

func main() {
    Crawl("http://golang.org/", 4, fetcher)
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
