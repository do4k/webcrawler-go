package html

import (
	"github.com/dandecrypted/webcrawler-go/strings"
	"golang.org/x/net/html"
)

func GetLinks(n *html.Node, sourceAddress string) []string {
	var ignorePrefixes = []string{"sftp", "ssh", "ftp", "mailto", "tel", "javascript", "#"}

	var links []string
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" && !strings.StartsWithAny(ignorePrefixes, a.Val) {
				links = append(links, NormaliseLink(a.Val, sourceAddress))
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = append(links, GetLinks(c, sourceAddress)...)
	}

	return links
}
