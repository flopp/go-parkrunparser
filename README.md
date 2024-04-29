# go-parkrunparser
A Go module to parse various parkrun webpages

* `events.json` -> `func ParseEvents(buf []byte) (Events, error)`
* `eventshistory` -> `func ParseEventHistory(buf []byte) (EventHistory, error)`
* `results/123`, `latestresults` -> `func ParseResults(buf []byte) (Results, error)`

Note: the package will *not download* any files from the parkrun webpage!