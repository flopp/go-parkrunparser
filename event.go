package parkrunparser

import (
	"fmt"
	"html"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Parkrunner struct {
	Id   string
	Name string
}

type Runner struct {
	*Parkrunner
	AgeGroup    AgeGroup
	Time        time.Duration
	Achievement Achievement
}

type Event struct {
	Index      int
	Date       time.Time
	Runners    []Runner
	Volunteers []Parkrunner
}

func parseDate(s string) (time.Time, error) {
	return time.Parse("02/01/2006", s)
}

func parseDuration(s string) (time.Duration, error) {
	split := strings.Split(s, ":")

	// hh:mm::ss
	if len(split) == 3 {
		return time.ParseDuration(fmt.Sprintf("%sh%sm%ss", split[0], split[1], split[2]))
	}

	// mm::ss
	if len(split) == 2 {
		return time.ParseDuration(fmt.Sprintf("%sm%ss", split[0], split[1]))
	}

	return 0, fmt.Errorf("cannot parse duration: %s", s)
}

var reDateIndex = regexp.MustCompile(`<h3><span class="format-date">([^<]+)</span><span class="spacer">[^<]*</span><span>#([0-9]+)</span></h3>`)
var reVolunteerRow = regexp.MustCompile(`<a href='\./athletehistory/\?athleteNumber=(\d+)'>([^<]+)</a>`)
var reRunnerRow0 = regexp.MustCompile(`<tr class="Results-table-row" [^<]*><td class="Results-table-td Results-table-td--position">\d+</td><td class="Results-table-td Results-table-td--name"><div class="compact">(<a href="[^"]*/\d+")?.*?</tr>`)
var reRunnerRow = regexp.MustCompile(`^<tr class="Results-table-row" data-name="([^"]*)" data-agegroup="([^"]*)" data-club="[^"]*" data-gender="[^"]*" data-position="\d+" data-runs="(\d+)" data-vols="(\d+)" data-agegrade="[^"]*" data-achievement="([^"]*)"><td class="Results-table-td Results-table-td--position">\d+</td><td class="Results-table-td Results-table-td--name"><div class="compact"><a href="[^"]*/(\d+)"`)
var reRunnerRowUnknown = regexp.MustCompile(`^<tr class="Results-table-row" data-name="([^"]*)" data-agegroup="" data-club="" data-position="\d+" data-runs="0" data-agegrade="0" data-achievement=""><td class="Results-table-td Results-table-td--position">\d+</td><td class="Results-table-td Results-table-td--name"><div class="compact">.*`)
var reTime = regexp.MustCompile(`Results-table-td--time[^"]*&#10;                      "><div class="compact">(\d?:?\d\d:\d\d)</div>`)

func ParseEvent(data string) (Event, error) {
	reNewline := regexp.MustCompile(`\r?\n`)
	data = reNewline.ReplaceAllString(data, " ")

	var event Event

	// date, index
	if match := reDateIndex.FindStringSubmatch(data); match != nil {
		if date, err := parseDate(match[1]); err == nil {
			event.Date = date
		} else {
			return Event{}, fmt.Errorf("cannot parse event date '%s': %w", match[1], err)
		}
		if index, err := strconv.Atoi(match[2]); err == nil {
			event.Index = index
		} else {
			return Event{}, fmt.Errorf("cannot parse event index '%s': %w", match[2], err)
		}
	} else {
		return Event{}, fmt.Errorf("cannot find date/index header")
	}

	// runners
	for row, match0 := range reRunnerRow0.FindAllStringSubmatch(data, -1) {
		if match := reRunnerRow.FindStringSubmatch(match0[0]); match != nil {
			name := html.UnescapeString(match[1])
			ageGroup, err := ParseAgeGroup(match[2])
			if err != nil {
				return Event{}, fmt.Errorf("runner row %d - while parsing age group: %w", row, err)
			}
			achievement, err := ParseAchievement(match[5])
			if err != nil {
				return Event{}, fmt.Errorf("runner row %d - while parsing achievement: %w", row, err)
			}
			id := match[6]
			var runTime time.Duration = 0
			if matchTime := reTime.FindStringSubmatch(match0[0]); matchTime != nil {
				runTime, err = parseDuration(matchTime[1])
				if err != nil {
					return Event{}, fmt.Errorf("runner row %d - while parsing time: %w", row, err)
				}
			}

			event.Runners = append(event.Runners, Runner{&Parkrunner{id, name}, ageGroup, runTime, achievement})
		} else if match := reRunnerRowUnknown.FindStringSubmatch(match0[0]); match != nil {
			name := html.UnescapeString(match[1])
			event.Runners = append(event.Runners, Runner{&Parkrunner{"", name}, AgeGroup{}, 0, AchievementNone})
		} else {
			return Event{}, fmt.Errorf("runner row %d - invalid format: %s", row, match0[0])
		}
	}

	// volunteers
	for _, match := range reVolunteerRow.FindAllStringSubmatch(data, -1) {
		id := match[1]
		name := html.UnescapeString(match[2])
		event.Volunteers = append(event.Volunteers, Parkrunner{id, name})
	}

	return event, nil
}
