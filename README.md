[![PkgGoDev](https://pkg.go.dev/badge/github.com/flopp/go-parkrunparser)](https://pkg.go.dev/github.com/flopp/go-parkrunparser)
[![Go Report Card](https://goreportcard.com/badge/github.com/flopp/go-parkrunparser)](https://goreportcard.com/report/flopp/go-parkrunparser)
![golang/static](https://github.com/flopp/go-parkrunparser/workflows/go/static/badge.svg)
[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/flopp/go-parkrunparser/)

# go-parkrunparser
A Go module to parse various parkrun webpages

* `events.json` -> `func ParseEvents(buf []byte) (Events, error)`
* `eventshistory` -> `func ParseEventHistory(buf []byte) (EventHistory, error)`
* `results/123`, `latestresults` -> `func ParseResults(buf []byte) (Results, error)`

Note: the package will *not download* any files from the parkrun webpage!