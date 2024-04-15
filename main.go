package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/dandecrypted/webcrawler-go/crawler"
	"github.com/dandecrypted/webcrawler-go/html"
	"github.com/dandecrypted/webcrawler-go/http"
	"github.com/temoto/robotstxt"
	xhtml "golang.org/x/net/html"
)

const UserAgent = "DanOakBot/1.0"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: go run main.go \"<your url (string)>\" <optional: throttle (int)>")
		os.Exit(1)
	}
	startUrl := os.Args[1]

	throttle := 20
	if len(os.Args) == 3 {
		throttleStr := os.Args[2]
		parsedThrottle, err := strconv.Atoi(throttleStr)
		if err != nil {
			fmt.Println("🚨 error parsing throttle value: " + err.Error())
			os.Exit(1)
		}
		throttle = parsedThrottle
	}

	c := crawler.NewCrawler(
		startUrl,
		throttle,
		url.Parse,
		func(currentUrl string) (string, error) {
			return http.GetContent(currentUrl, UserAgent)
		},
		func(content string) (*xhtml.Node, error) {
			return xhtml.Parse(strings.NewReader(content))
		},
		html.GetLinks,
		func(url string, robots *robotstxt.RobotsData) bool {
			if robots == nil {
				return true
			}

			return robots.TestAgent(url, UserAgent)
		},
	)

	c.Crawl()

	fmt.Println("📃 visited links for", startUrl)
	for _, link := range c.Visited {
		fmt.Printf("✈️  %s\n", link)
	}
}
