package html

import "strings"

func NormaliseLink(link, baseAddress string) string {
	link = strings.TrimSpace(link)
	if len(link) == 0 {
		return ""
	}

	if link[0] == '/' {
		link = baseAddress + link
	}

	if i := strings.Index(link, "?"); i != -1 {
		link = link[:i]
	}

	if i := strings.Index(link, "#"); i != -1 {
		link = link[:i]
	}

	if len(link) > 0 && link[len(link)-1] == '/' {
		link = link[:len(link)-1]
	}

	return link
}
