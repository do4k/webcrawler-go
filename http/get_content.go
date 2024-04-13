package http

import (
	"net/http"

	"golang.org/x/net/html"
)

const UserAgent = "DanOakBot/1.0"

func GetContent(url string) (*html.Node, error) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", UserAgent)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}
