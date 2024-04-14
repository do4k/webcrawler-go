package main

import (
	"fmt"
	"net/url"
	"time"

	"github.com/dandecrypted/webcrawler-go/data"
	"golang.org/x/net/html"
)

type Crawler struct {
	startingUrl    string
	visited        data.Queue
	queue          data.Queue
	throttle       int
	parseUrl       func(string) (*url.URL, error)
	getHttpContent func(string) (*html.Node, error)
	getLinks       func(*html.Node, string) []string
}

func NewCrawler(
	startingUrl string,
	throttle int,
	parseUrl func(string) (*url.URL, error),
	getHttpContent func(string) (*html.Node, error),
	getLinks func(*html.Node, string) []string) *Crawler {

	return &Crawler{
		startingUrl:    startingUrl,
		visited:        data.Queue{},
		queue:          data.Queue{startingUrl},
		throttle:       throttle,
		parseUrl:       parseUrl,
		getHttpContent: getHttpContent,
		getLinks:       getLinks,
	}
}

func (c *Crawler) Crawl() {
	hasItems := true

	for hasItems {
		u, hasItems := c.queue.Dequeue()
		if !hasItems {
			break
		}

		sourceUrl, err := c.parseUrl(u)
		if err != nil {
			errMsg := fmt.Errorf("error parsing url %s: %s", u, err)
			fmt.Println(errMsg)
			continue
		}

		fmt.Printf("-----\ncrawling %s\n-----\n", sourceUrl)

		if sourceUrl.String() == "" {
			errMsg := fmt.Errorf("url cannot be empty")
			fmt.Println(errMsg)
			continue
		}

		doc, err := c.getHttpContent(u)
		if err != nil {
			errMsg := fmt.Errorf("error crawling %s: %s", u, err)
			fmt.Println(errMsg)
		}

		if doc != nil {
			links := c.getLinks(doc, c.startingUrl)
			for _, link := range links {
				err := c.processLink(link, sourceUrl)
				if err != nil {
					fmt.Println("🚨 " + err.Error())
					continue
				}
				fmt.Printf("✅ added %s to the queue\n", link)
				c.queue.Enqueue(link)
			}
		}

		c.visited.Enqueue(u)
		if c.queue.Count() > 0 {
			fmt.Printf("sleeping for %d seconds\n", c.throttle)
			time.Sleep(time.Duration(c.throttle) * time.Second)
		}
	}
}

func (c *Crawler) processLink(link string, sourceUrl *url.URL) error {
	if link == "" {
		return fmt.Errorf("link cannot be empty")
	}

	if c.queue.Contains(link) {
		return fmt.Errorf("link %s already in queue", link)
	}

	if c.visited.Contains(link) {
		err := fmt.Errorf("link %s already visited", link)
		return err
	}

	parsedUrlForLink, parseErr := c.parseUrl(link)
	if parseErr != nil {
		return fmt.Errorf("error parsing url %s: %s", link, parseErr)
	}

	if (sourceUrl.Host == parsedUrlForLink.Host) && (sourceUrl.Path == parsedUrlForLink.Path) {
		return fmt.Errorf("link %s is the same as the source url", link)
	}

	if parsedUrlForLink.Host != sourceUrl.Host {
		return fmt.Errorf("link %s is not in the same domain", link)
	}

	return nil
}
