package parkrunparser

import (
	"fmt"
	"regexp"
	"strconv"
)

type EventHistory struct {
	Results []*Results
}

var patternNumberOfResults = regexp.MustCompile("<td class=\"Results-table-td Results-table-td--position\"><a href=\"\\.\\./(\\d+)\">(\\d+)</a></td>")
var patternResultsRow = regexp.MustCompile(`<tr class="Results-table-row" data-parkrun="(\d+)" data-date="([^"]+)" data-finishers="(\d+)" data-volunteers="(\d+)" data-male="([^"]*)" data-female="([^"]*)" data-maletime="(\d*)" data-femaletime="(\d*)">`)

func ParseEventHistory(buf []byte) (EventHistory, error) {
	reNewline := regexp.MustCompile(`\r?\n`)
	data := reNewline.ReplaceAllString(string(buf), " ")

	var history EventHistory

	match := patternNumberOfResults.FindStringSubmatch(data)
	if match == nil {
		return history, fmt.Errorf("cannot find number of results")
	}

	count, err := strconv.Atoi(match[1])
	if err != nil || count < 0 {
		return history, fmt.Errorf("cannot parse number of results from '%s': %v", match[1], err)
	}

	history.Results = make([]*Results, count)
	matches := patternResultsRow.FindAllStringSubmatch(data, -1)
	for row, match := range matches {
		index, err := strconv.Atoi(match[1])
		if err != nil {
			return history, fmt.Errorf("row %d: cannot parse results index from '%s': %v", row, match[1], err)
		}
		if index <= 0 || index > count {
			return history, fmt.Errorf("row %d: invalid index '%d'; must be >= 1 and <= %d", row, index, count)
		}

		date, err := ParseDate(match[2])
		if err != nil {
			return history, fmt.Errorf("row %d: cannot parse results date from '%s': %v", row, match[2], err)
		}

		finishers, err := strconv.Atoi(match[3])
		if err != nil {
			return history, fmt.Errorf("row %d: cannot parse number of finishers from '%s': %v", row, match[3], err)
		}

		volunteers, err := strconv.Atoi(match[4])
		if err != nil {
			return history, fmt.Errorf("row %d: cannot parse number of volunteers from '%s': %v", row, match[4], err)
		}

		if history.Results[index-1] != nil {
			return history, fmt.Errorf("row %d: duplicate results index %d", row, index)
		}
		history.Results[index-1] = &Results{index, date, finishers, volunteers, nil, nil}
	}

	for index, results := range history.Results {
		if results == nil {
			return history, fmt.Errorf("missing results %d", index+1)
		}
	}

	return history, nil
}
