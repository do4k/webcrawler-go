package main

import (
	"errors"
	"net/url"
	"testing"

	"golang.org/x/net/html"
)

func TestCrawl(t *testing.T) {
	parseUrl := func(s string) (*url.URL, error) {
		return url.Parse(s)
	}
	getHttpContent := func(s string) (*html.Node, error) {
		return &html.Node{}, nil
	}
	getLinks := func(n *html.Node, s string) []string {
		return []string{"http://test.com/link1", "http://test.com/link2"}
	}

	c := NewCrawler("http://test.com", 1, parseUrl, getHttpContent, getLinks)
	c.Crawl()

	if len(c.visited) != 3 {
		t.Errorf("Expected visited to have 3 items, got %d", len(c.visited))
	}
}

func TestCrawl_FailsToParseSecondUrl(t *testing.T) {
	parseUrl := func(s string) (*url.URL, error) {
		if s == "http://test.com/link1" {
			// returns nil and new error
			return nil, errors.New("Could not parse https://test.com/link1")
		}
		return url.Parse(s)
	}

	getHttpContent := func(s string) (*html.Node, error) {
		return &html.Node{}, nil
	}
	getLinks := func(n *html.Node, s string) []string {
		return []string{"http://test.com/link1", "http://test.com/link2"}
	}

	c := NewCrawler("http://test.com", 1, parseUrl, getHttpContent, getLinks)
	c.Crawl()

	if len(c.visited) != 2 {
		t.Errorf("Expected visited to have 2 items, got %d", len(c.visited))
	}

	if (c.visited[0] != "http://test.com") || (c.visited[1] != "http://test.com/link2") {
		t.Errorf("Expected visited to contain 'http://test.com' and 'http://test.com/link2'")
	}
}
