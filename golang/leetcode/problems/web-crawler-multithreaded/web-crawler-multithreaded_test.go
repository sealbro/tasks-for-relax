package web_crawler_multithreaded

import (
	"fmt"
	"golang/assert"
	"slices"
	"testing"
)

func TestCase(t *testing.T) {
	testCases := []struct {
		startUrl string
		urls     []string
		edges    [][]int
		expected []string
	}{
		{
			startUrl: "http://news.yahoo.com/news/topics/",
			urls:     []string{"http://news.yahoo.com", "http://news.yahoo.com/news", "http://news.yahoo.com/news/topics/", "http://news.google.com", "http://news.yahoo.com/us"},
			edges:    [][]int{{2, 0}, {2, 1}, {3, 2}, {3, 1}, {0, 4}},
			expected: []string{"http://news.yahoo.com", "http://news.yahoo.com/news", "http://news.yahoo.com/news/topics/", "http://news.yahoo.com/us"},
		},
		{
			startUrl: "http://news.google.com",
			urls:     []string{"http://news.yahoo.com", "http://news.yahoo.com/news", "http://news.yahoo.com/news/topics/", "http://news.google.com"},
			edges:    [][]int{{0, 2}, {2, 1}, {3, 2}, {3, 1}, {0, 3}},
			expected: []string{"http://news.google.com"},
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc), func(t *testing.T) {
			h := &htmlParser{
				urls:  tc.urls,
				edges: tc.edges,
			}

			actualUrls := crawl(tc.startUrl, h)

			slices.Sort(actualUrls)
			slices.Sort(tc.expected)

			assert.EqualMany(t, tc.expected, actualUrls)
		})
	}
}
