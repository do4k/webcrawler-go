package main

import (
	"net/url"
	"testing"

	"golang.org/x/net/html"
)

func TestNewCrawler(t *testing.T) {
	parseUrl := func(s string) (*url.URL, error) {
		return url.Parse(s)
	}
	getHttpContent := func(s string) (*html.Node, error) {
		return &html.Node{}, nil
	}
	getLinks := func(n *html.Node, s string) []string {
		return []string{}
	}

	c := NewCrawler("http://test.com", 1, parseUrl, getHttpContent, getLinks)

	if c.startingUrl != "http://test.com" {
		t.Errorf("Expected startingUrl to be 'http://test.com', got '%s'", c.startingUrl)
	}
}

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

	if len(c.visited) != 1 {
		t.Errorf("Expected visited to have 1 item, got %d", len(c.visited))
	}
}
