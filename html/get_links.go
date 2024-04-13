package html

import "golang.org/x/net/html"

func GetLinks(n *html.Node, sourceAddress string) []string {
	var links []string
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, NormaliseLink(a.Val, sourceAddress))
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = append(links, GetLinks(c, sourceAddress)...)
	}

	return links
}
