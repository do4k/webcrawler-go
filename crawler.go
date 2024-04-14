package main

import (
	"fmt"
	"net/url"
	"time"

	"github.com/dandecrypted/webcrawler-go/data"
	"github.com/temoto/robotstxt"
	"golang.org/x/net/html"
)

type Crawler struct {
	startingUrl    string
	visited        data.Queue
	queue          data.Queue
	throttle       time.Duration
	robotsTxt      *robotstxt.RobotsData
	parseUrl       func(string) (*url.URL, error)
	getHttpContent func(string) (string, error)
	parseHtml      func(string) (*html.Node, error)
	getLinks       func(*html.Node, string) []string
	robotsAllowed  func(string, *robotstxt.RobotsData) bool
}

func NewCrawler(
	startingUrl string,
	throttle int,
	parseUrl func(string) (*url.URL, error),
	getHttpContent func(string) (string, error),
	parseHtml func(string) (*html.Node, error),
	getLinks func(*html.Node, string) []string,
	robotsAllowed func(string, *robotstxt.RobotsData) bool) *Crawler {

	parsedStartingUrl, err := parseUrl(startingUrl)
	if err != nil {
		fmt.Println("🚨 error parsing starting url " + err.Error())
		return nil
	}

	robotsTxtContent, err := getHttpContent(parsedStartingUrl.Scheme + "://" + parsedStartingUrl.Host + "/robots.txt")
	if err != nil {
		fmt.Println("🚨 error fetching robots.txt " + err.Error())
	}

	robots, err := robotstxt.FromBytes([]byte(robotsTxtContent))
	if err != nil {
		fmt.Println("🚨 error parsing robots.txt " + err.Error())
	}

	return &Crawler{
		startingUrl:    startingUrl,
		visited:        data.Queue{},
		queue:          data.Queue{startingUrl},
		throttle:       time.Duration(throttle) * time.Second,
		parseUrl:       parseUrl,
		robotsTxt:      robots,
		getHttpContent: getHttpContent,
		parseHtml:      parseHtml,
		getLinks:       getLinks,
		robotsAllowed:  robotsAllowed,
	}
}

func (c *Crawler) Crawl() {
	hasItems := true

	for hasItems {
		deqeuedUrl, hasItems := c.queue.Dequeue()
		if !hasItems {
			break
		}

		currentUrl, err := c.parseUrl(deqeuedUrl)
		if err != nil {
			errMsg := fmt.Errorf("🚨 error parsing url %s %s", deqeuedUrl, err)
			fmt.Println(errMsg)
			continue
		}

		fmt.Printf("-----\ncrawling %s (%d items in the queue)\n-----\n", currentUrl, c.queue.Count())

		if currentUrl.String() == "" {
			errMsg := fmt.Errorf("🚨 url cannot be empty")
			fmt.Println(errMsg)
			continue
		}

		if !c.robotsAllowed(deqeuedUrl, c.robotsTxt) {
			errMsg := fmt.Errorf("🚨 url %s is disallowed by robots.txt", currentUrl)
			fmt.Println(errMsg)
			continue
		}

		content, err := c.getHttpContent(deqeuedUrl)
		if err != nil {
			errMsg := fmt.Errorf("🚨 error fetching content for %s %s", deqeuedUrl, err)
			fmt.Println(errMsg)
			continue
		}

		doc, err := c.parseHtml(content)
		if err != nil {
			errMsg := fmt.Errorf("🚨 error parsing html for %s %s", deqeuedUrl, err)
			fmt.Println(errMsg)
			c.sleep()
			continue
		}

		if doc != nil {
			links := c.getLinks(doc, c.startingUrl)
			for _, link := range links {
				err := c.processLink(link, currentUrl)
				if err != nil {
					fmt.Println("🚨 " + err.Error())
					continue
				}
				fmt.Printf("✅ added %s to the queue\n", link)
				c.queue.Enqueue(link)
			}
		}

		c.visited.Enqueue(deqeuedUrl)
		c.sleep()
	}
}

// this is to be a good citizen and not spam the target server
func (c *Crawler) sleep() {
	if c.queue.Count() > 0 {
		fmt.Printf("sleeping for %d seconds %d items in the queue\n", int(c.throttle.Seconds()), c.queue.Count())
		time.Sleep(c.throttle)
	}
}

func (c *Crawler) processLink(link string, sourceUrl *url.URL) error {
	if link == "" {
		return fmt.Errorf("link cannot be empty")
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

	if c.queue.Contains(link) {
		return fmt.Errorf("link %s already in queue", link)
	}

	if c.visited.Contains(link) {
		err := fmt.Errorf("link %s already visited", link)
		return err
	}

	if !c.robotsAllowed(link, c.robotsTxt) {
		return fmt.Errorf("link %s is disallowed by robots.txt", link)
	}

	return nil
}
