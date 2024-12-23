package web_crawler_multithreaded

import (
	"context"
	"net/url"
	"slices"
	"sync"
	"time"
)

// https://leetcode.com/problems/web-crawler-multithreaded/

type HtmlParser interface {
	GetUrls(url string) []string
}

type htmlParser struct {
	urls  []string
	edges [][]int
}

func (h *htmlParser) GetUrls(startUrl string) []string {
	// Simulate network delay
	time.Sleep(100 * time.Millisecond)

	startIndex := slices.Index(h.urls, startUrl)
	if startIndex == -1 {
		return nil
	}

	var result []string
	for _, edge := range h.edges {
		if edge[0] != startIndex {
			continue
		}

		result = append(result, h.urls[edge[1]])
	}

	return result
}

func crawl(startUrl string, htmlParser HtmlParser) []string {
	const maxWorkers = 10

	ctx, cancelFunc := context.WithCancel(context.Background())
	visited := make(map[string]struct{}, 1)
	input := make(chan string)
	output := make(chan []string)
	lock := sync.RWMutex{}

	var wg sync.WaitGroup
	wg.Add(1)
	startParsedUrl, _ := url.Parse(startUrl)

	for i := 0; i < maxWorkers; i++ {
		go func() {
			for {
				select {
				case getUrl, chOk := <-input:
					if !chOk {
						return
					}

					lock.Lock()
					visited[getUrl] = struct{}{}
					lock.Unlock()

					var urls []string
					for _, u := range htmlParser.GetUrls(getUrl) {
						parsedUrl, _ := url.Parse(u)
						// Check if the host is the same as the start host
						if parsedUrl.Host == startParsedUrl.Host {
							urls = append(urls, u)
						}
					}
					if len(urls) != 0 {
						wg.Add(len(urls))
						output <- urls
					}
					wg.Done()
				case <-ctx.Done():
					return
				}
			}
		}()
	}

	go func() {
		defer close(input)
		input <- startUrl
		for {
			select {
			case urls, chOk := <-output:
				if !chOk {
					return
				}

				for _, u := range urls {
					lock.RLock()
					if _, ok := visited[u]; ok {
						lock.RUnlock()
						continue
					}
					lock.RUnlock()

					input <- u
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	wg.Wait()
	cancelFunc()
	close(output)

	result := make([]string, 0, len(visited))
	for k := range visited {
		result = append(result, k)
	}
	return result
}
