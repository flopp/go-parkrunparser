package parkrunparser

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

type EventHistory struct {
	Events []*Event
}

var patternNumberOfRuns = regexp.MustCompile("<td class=\"Results-table-td Results-table-td--position\"><a href=\"\\.\\./(\\d+)\">(\\d+)</a></td>")
var patternRunRow = regexp.MustCompile(`<tr class="Results-table-row" data-parkrun="(\d+)" data-date="(\d+/\d+/\d+)" data-finishers="(\d+)" data-volunteers="(\d+)" data-male="([^"]*)" data-female="([^"]*)" data-maletime="(\d*)" data-femaletime="(\d*)">`)

func ParseEventHistory(data string) (EventHistory, error) {
	reNewline := regexp.MustCompile(`\r?\n`)
	data = reNewline.ReplaceAllString(data, " ")

	var history EventHistory

	match := patternNumberOfRuns.FindStringSubmatch(data)
	if match == nil {
		return history, fmt.Errorf("cannot find number of events")
	}

	count, err := strconv.Atoi(match[1])
	if err != nil || count < 0 {
		return history, fmt.Errorf("cannot parse number of events from '%s': %v", match[1], err)
	}

	history.Events = make([]*Event, count)
	matches := patternRunRow.FindAllStringSubmatch(data, -1)
	for row, match := range matches {
		index, err := strconv.Atoi(match[1])
		if err != nil {
			return history, fmt.Errorf("row %d: cannot parse event index from '%s': %v", row, match[1], err)
		}
		if index <= 0 || index > count {
			return history, fmt.Errorf("row %d: invalid index '%d'; must be >= 1 and <= %d", row, index, count)
		}

		date, err := time.Parse("02/01/2006", match[2])
		if err != nil {
			return history, fmt.Errorf("row %d: cannot parse event date from '%s': %v", row, match[2], err)
		}

		finishers, err := strconv.Atoi(match[3])
		if err != nil {
			return history, fmt.Errorf("row %d: cannot parse numner of finishers from '%s': %v", row, match[3], err)
		}

		volunteers, err := strconv.Atoi(match[4])
		if err != nil {
			return history, fmt.Errorf("row %d: cannot parse numner of volunteers from '%s': %v", row, match[4], err)
		}

		if history.Events[index-1] != nil {
			return history, fmt.Errorf("row %d: duplicate event index %d", row, index)
		}
		history.Events[index-1] = &Event{index, date, finishers, volunteers, nil, nil}
	}

	for index, event := range history.Events {
		if event == nil {
			return history, fmt.Errorf("missing event %d", index+1)
		}
	}

	return history, nil
}
