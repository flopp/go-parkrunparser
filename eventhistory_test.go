package parkrunparser

import (
	"testing"
)

func loadEventHistory(t *testing.T, fileName string, expectedResults int) {
	data, err := readFile(fileName)
	if err != nil {
		t.Errorf("failed to read file '%s': %v", fileName, err)
	}

	history, err := ParseEventHistory(data)
	if err != nil {
		t.Errorf("%s: failed to parse eventhistory: %v", fileName, err)
	}

	if len(history.Results) != expectedResults {
		t.Errorf("%s: unexpected number of results: %d; expected: %d", fileName, len(history.Results), expectedResults)
	}
}

func TestParseEventHistory(t *testing.T) {
	loadEventHistory(t, "test-data/de-dietenbach-history.gz", 129)
	loadEventHistory(t, "test-data/uk-simmons-park-history.gz", 7)
	loadEventHistory(t, "test-data/uk-eastville-history.gz", 289)
}
