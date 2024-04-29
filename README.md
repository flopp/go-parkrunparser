# go-parkrunparser
A Go module to parse various parkrun webpages


* `eventshistory.html` -> `func ParseEventHistory(data string) (EventHistory, error)`
* results page -> `func ParseResults(data string) (Results, error)`