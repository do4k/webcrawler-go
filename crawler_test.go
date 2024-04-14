package main

import (
	"errors"
	"net/url"
	"testing"

	"github.com/temoto/robotstxt"
	"golang.org/x/net/html"
)

func defaultParseUrl(s string) (*url.URL, error) {
	return url.Parse(s)
}

func defaultGetHttpContent(s string) (string, error) {
	return "", nil
}

func defaultParseHtml(content string) (*html.Node, error) {
	return &html.Node{}, nil
}

func defaultGetLinks(n *html.Node, s string) []string {
	return []string{"http://test.com/link1", "http://test.com/link2"}
}

func defaultRobotsTxtAllowed(link string, robotsTxt *robotstxt.RobotsData) bool {
	return true
}

func TestCrawl(t *testing.T) {
	c := NewCrawler("http://test.com", 1, defaultParseUrl, defaultGetHttpContent, defaultParseHtml, defaultGetLinks, defaultRobotsTxtAllowed)
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

	c := NewCrawler("http://test.com", 1, parseUrl, defaultGetHttpContent, defaultParseHtml, defaultGetLinks, defaultRobotsTxtAllowed)
	c.Crawl()

	if len(c.visited) != 2 {
		t.Errorf("Expected visited to have 2 items, got %d", len(c.visited))
	}

	if (c.visited[0] != "http://test.com") || (c.visited[1] != "http://test.com/link2") {
		t.Errorf("Expected visited to contain 'http://test.com' and 'http://test.com/link2'")
	}
}

func TestCrawl_RobotsTxtDisallowed(t *testing.T) {
	c := NewCrawler("http://test.com", 1, defaultParseUrl, defaultGetHttpContent, defaultParseHtml, defaultGetLinks, func(url string, robotsTxt *robotstxt.RobotsData) bool {
		return url != "http://test.com/link1"
	})

	c.Crawl()

	if len(c.visited) != 2 {
		t.Errorf("Expected visited to have 2 items, got %d", len(c.visited))
	}

	if (c.visited[0] != "http://test.com") || (c.visited[1] != "http://test.com/link2") {
		t.Errorf("Expected visited to contain 'http://test.com' and 'http://test.com/link2'")
	}
}

func TestCrawl_DuplicateLinksOnlyVisitedOnce(t *testing.T) {
	getLinks := func(n *html.Node, s string) []string {
		return []string{"http://test.com/link1", "http://test.com/link1"}
	}

	c := NewCrawler("http://test.com", 1, defaultParseUrl, defaultGetHttpContent, defaultParseHtml, getLinks, defaultRobotsTxtAllowed)
	c.Crawl()

	if len(c.visited) != 2 {
		t.Errorf("Expected visited to have 2 items, got %d", len(c.visited))
	}
}

func TestCrawl_LinkIsSameAsSource(t *testing.T) {
	getLinks := func(n *html.Node, s string) []string {
		return []string{"http://test.com"}
	}

	c := NewCrawler("http://test.com", 1, defaultParseUrl, defaultGetHttpContent, defaultParseHtml, getLinks, defaultRobotsTxtAllowed)
	c.Crawl()

	if len(c.visited) != 1 {
		t.Errorf("Expected visited to have 1 item, got %d", len(c.visited))
	}
}
