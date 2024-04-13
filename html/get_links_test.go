package html

import (
	"testing"

	"golang.org/x/net/html"
)

func TestGetLinks(t *testing.T) {
	n := &html.Node{
		Type: html.ElementNode,
		Data: "a",
		Attr: []html.Attribute{
			{Key: "href", Val: "https://example.com"},
		},
	}

	links := GetLinks(n, "")

	if len(links) != 1 {
		t.Errorf("Expected 1 link, got %d", len(links))
	}

	if links[0] != "https://example.com" {
		t.Errorf("Expected 'https://example.com', got '%s'", links[0])
	}
}

func TestGetLinks_NoLinks(t *testing.T) {
	n := &html.Node{
		Type: html.ElementNode,
		Data: "div",
		Attr: []html.Attribute{
			{Key: "a", Val: "https://i-know-this-isnt-relistic.com"},
		},
	}

	sourceAddress := "https://base.com"

	links := GetLinks(n, sourceAddress)

	if len(links) != 0 {
		t.Errorf("Expected 0 links, got %d", len(links))
	}
}
