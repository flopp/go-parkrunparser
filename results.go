package parkrunparser

import (
	"fmt"
	"html"
	"regexp"
	"strconv"
	"time"
)

type Parkrunner struct {
	Id   string
	Name string
}

func (p Parkrunner) IsUnknown() bool {
	return len(p.Id) == 0
}

type Finisher struct {
	*Parkrunner
	AgeGroup              AgeGroup
	Time                  time.Duration
	Achievement           Achievement
	NumberOfRuns          int
	NumberOfVolunteerings int
}

func (f Finisher) IsUnknown() bool {
	return len(f.Id) == 0
}

type Results struct {
	Index              int
	Date               time.Time
	NumberOfFinishers  int
	NumberOfVolunteers int
	Finishers          []Finisher
	Volunteers         []Parkrunner
}

var reDateIndex = regexp.MustCompile(`<h3><span class="format-date">([^<]+)</span><span class="spacer">[^<]*</span><span>#([0-9]+)</span></h3>`)
var reVolunteerRow1 = regexp.MustCompile(`<a href='\./athletehistory/\?athleteNumber=(\d+)'>([^<]+)</a>`)
var reVolunteerRow2 = regexp.MustCompile(`<a href="/[^/]+/parkrunner/(\d+)">([^<]+)</a>`)
var reRunnerRow0 = regexp.MustCompile(`<tr class="Results-table-row" [^<]*><td class="Results-table-td Results-table-td--position">\d+</td><td class="Results-table-td Results-table-td--name"><div class="compact">(<a href="[^"]*/\d+")?.*?</tr>`)
var reRunnerRow = regexp.MustCompile(`^<tr class="Results-table-row" data-name="([^"]*)" data-agegroup="([^"]*)" data-club="[^"]*" (?:data-groups="[^"]*" )?data-gender="[^"]*" data-position="\d+" data-runs="(\d+)" data-vols="(\d+)" data-agegrade="[^"]*" data-achievement="([^"]*)"><td class="Results-table-td Results-table-td--position">\d+</td><td class="Results-table-td Results-table-td--name"><div class="compact"><a href="[^"]*/(\d+)"`)
var reRunnerRowUnknown = regexp.MustCompile(`^<tr class="Results-table-row" data-name="([^"]*)" data-agegroup="" data-club="" (?:data-groups="" )?data-position="\d+" data-runs="0" data-agegrade="0" data-achievement=""><td class="Results-table-td Results-table-td--position">\d+</td><td class="Results-table-td Results-table-td--name"><div class="compact">.*`)
var reTime = regexp.MustCompile(`Results-table-td--time[^"]*&#10;\s*"><div class="compact">(\d?:?\d\d:\d\d)</div>`)

func ParseResults(buf []byte) (Results, error) {
	reNewline := regexp.MustCompile(`\r?\n`)
	data := reNewline.ReplaceAllString(string(buf), " ")

	var event Results

	// date, index
	if match := reDateIndex.FindStringSubmatch(data); match != nil {
		if date, err := ParseDate(match[1]); err == nil {
			event.Date = date
		} else {
			return Results{}, fmt.Errorf("cannot parse event date '%s': %w", match[1], err)
		}
		if index, err := strconv.Atoi(match[2]); err == nil {
			event.Index = index
		} else {
			return Results{}, fmt.Errorf("cannot parse event index '%s': %w", match[2], err)
		}
	} else {
		return Results{}, fmt.Errorf("cannot find date/index header")
	}

	// runners
	for row, match0 := range reRunnerRow0.FindAllStringSubmatch(data, -1) {
		if match := reRunnerRow.FindStringSubmatch(match0[0]); match != nil {
			name := html.UnescapeString(match[1])
			ageGroup, err := ParseAgeGroup(match[2])
			if err != nil {
				return Results{}, fmt.Errorf("runner row %d - while parsing age group: %w", row, err)
			}
			achievement, err := ParseAchievement(match[5])
			if err != nil {
				return Results{}, fmt.Errorf("runner row %d - while parsing achievement: %w", row, err)
			}
			id := match[6]
			var runTime time.Duration = 0
			if matchTime := reTime.FindStringSubmatch(match0[0]); matchTime != nil {
				runTime, err = ParseDuration(matchTime[1])
				if err != nil {
					return Results{}, fmt.Errorf("runner row %d - while parsing time: %w", row, err)
				}
			}

			runs, err := strconv.Atoi(match[3])
			if err != nil {
				return Results{}, fmt.Errorf("runner row %d - while parsing #runs: %w", row, err)
			}

			vols, err := strconv.Atoi(match[4])
			if err != nil {
				return Results{}, fmt.Errorf("runner row %d - while parsing #vols: %w", row, err)
			}

			event.Finishers = append(event.Finishers, Finisher{&Parkrunner{id, name}, ageGroup, runTime, achievement, runs, vols})
		} else if match := reRunnerRowUnknown.FindStringSubmatch(match0[0]); match != nil {
			name := html.UnescapeString(match[1])
			event.Finishers = append(event.Finishers, Finisher{&Parkrunner{"", name}, AgeGroup{}, 0, AchievementNone, 0, 0})
		} else {
			return Results{}, fmt.Errorf("runner row %d - invalid format: %s", row, match0[0])
		}
	}
	event.NumberOfFinishers = len(event.Finishers)

	// volunteers
	for _, match := range reVolunteerRow1.FindAllStringSubmatch(data, -1) {
		id := match[1]
		name := html.UnescapeString(match[2])
		event.Volunteers = append(event.Volunteers, Parkrunner{id, name})
	}
	for _, match := range reVolunteerRow2.FindAllStringSubmatch(data, -1) {
		id := match[1]
		name := html.UnescapeString(match[2])
		event.Volunteers = append(event.Volunteers, Parkrunner{id, name})
	}
	event.NumberOfVolunteers = len(event.Volunteers)

	return event, nil
}
