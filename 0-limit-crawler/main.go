//////////////////////////////////////////////////////////////////////
//
// Your task is to change the code to limit the crawler to at most one
// page per second, while maintaining concurrency (in other words,
// Crawl() must be called concurrently)
//
// @hint: you can achieve this by adding 3 lines
//

package main

import (
	"fmt"
	"sync"
	"time"
)

// Дан краулер (модифицированный из тура Go), который запрашивает страницы
// чрезмерно часто. Однако мы не хотим слишком нагружать веб-сервер
// сильно. Ваша задача - изменить код таким образом, чтобы он ограничивался не более чем
// одной страницы в секунду, сохраняя при этом параллелизм (другими словами,
// Crawl() должен вызываться параллельно)

// Crawl uses `fetcher` from the `mockfetcher.go` file to imitate a
// real crawler. It crawls until the maximum depth has reached.
func Crawl(url string, depth int, wg *sync.WaitGroup, token chan struct{}) {
	defer wg.Done()

	if depth <= 0 {
		return
	}

	token <- struct{}{}
	body, urls, err := fetcher.Fetch(url)
	time.Sleep(1 * time.Second)
	<-token
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)

	wg.Add(len(urls))
	for _, u := range urls {
		// Do not remove the `go` keyword, as Crawl() must be
		// called concurrently
		go Crawl(u, depth-1, wg, token)
	}
}

func main() {
	var wg sync.WaitGroup
	token := make(chan struct{}, 1)

	wg.Add(1)
	Crawl("http://golang.org/", 4, &wg, token)
	wg.Wait()
}
