package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"

	"github.com/dandecrypted/webcrawler-go/html"
	"github.com/dandecrypted/webcrawler-go/http"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: go run main.go <url> <optional: throttle>")
		os.Exit(1)
	}

	baseAddress := html.NormaliseLink(os.Args[1], "")
	_, err := url.Parse(baseAddress)
	if err != nil {
		fmt.Println("error parsing url: " + err.Error())
		os.Exit(1)
	}

	throttle := 20
	if len(os.Args) == 3 {
		throttleStr := os.Args[2]
		throttle, err = strconv.Atoi(throttleStr)
		if err != nil {
			fmt.Println("error parsing throttle value: " + err.Error())
			os.Exit(1)
		}
	}

	var wg sync.WaitGroup
	wg.Add(1)

	c := NewCrawler(
		baseAddress,
		throttle,
		url.Parse,
		http.GetContent,
		html.GetLinks,
	)

	go func() {
		c.Crawl()
		wg.Done()
	}()

	wg.Wait()
}
