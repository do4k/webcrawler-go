package http

import (
	"net/url"
	"strings"
)

func NormaliseLink(link, baseAddress string) string {
	link = strings.TrimSpace(link)
	if len(link) == 0 {
		return ""
	}

	if strings.HasPrefix(link, "//") {
		parsed, err := url.Parse(link)
		if err != nil {
			return parsed.Scheme + ":" + link
		}
		return "https:" + link
	}

	// I'm aware this doesn't account for ../ links and I was hoping that a nice URL library would handle that case
	if !strings.HasPrefix(link, "http") {
		trim := "./"
		link = strings.TrimRight(baseAddress, trim) + "/" + strings.TrimLeft(link, trim)
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
