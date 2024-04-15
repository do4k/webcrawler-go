# Golang Webscraper

To use this app call `go run . "<your url (string)>" <optional: throttle (int seconds) (default 20)>` from your terminal which will then crawl the link provided and output any links found and crawl them if they are in the same domain. The app will sleep for `throttle` seconds between making calls to be a good citizen and not completely spam the target server.

# Tests

run `go test` from terminal

# Features 
- [X] Scans a given starting url for any more Urls
- [X] All found urls are added to a queue
- [X] Queue has throttling to avoid getting denylisted for too many requests
  - [X] Enable this to be managed by configuration / command line argument
- [X] Robots.txt integration

# Known limitations
- parent directory relative paths are funky (..../../../) etc
- no real use of concurrency as web crawlers generally should be slow to avoid being rate limited and to be good citizens. Concurrency could be considered if a new channel is made for every new domain that is found to be crawled.
- state is not persisted
- runs cannot be exited early or resumed
- output is currently quite verbose, could use a debug flag or something to reduce some of the noise around why links aren't considered to continue to be crawled.