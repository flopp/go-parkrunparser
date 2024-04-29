package parkrunparser

import (
	"testing"
)

func TestParseEvents(t *testing.T) {
	fileName := "test-data/events.json.gz"
	data, err := readFile(fileName)
	if err != nil {
		t.Errorf("failed to read file '%s': %v", fileName, err)
	}

	events, err := ParseEvents(data)
	if err != nil {
		t.Errorf("%s: failed to parse events: %v", fileName, err)
	}

	if len(events.Countries) != 20 {
		t.Errorf("%s: number of countries = %d; expected = 20", fileName, len(events.Countries))
	}
	if len(events.Events) != 2533 {
		t.Errorf("%s: number of events = %d; expected = 2533", fileName, len(events.Events))
	}
	for _, country := range events.Countries {
		name := country.Name()
		if name == "UNKNOWN" {
			t.Errorf("%s: country name = 'UNKNOWN' (url = %s)", fileName, country.Url)
		}
	}
}
