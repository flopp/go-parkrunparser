package parkrunparser

import (
	"testing"
)

func TestParseEventHistory(t *testing.T) {
	fileName := "test-data/de-dietenbach-history.gz"
	data, err := readFile(fileName)
	if err != nil {
		t.Errorf("failed to read file '%s': %v", fileName, err)
	}

	history, err := ParseEventHistory(string(data))
	if err != nil {
		t.Errorf("%s: failed to parse eventhistory: %v", fileName, err)
	}

	if len(history.Results) != 129 {
		t.Errorf("%s: unexpected number of results: %d; expected: %d", fileName, len(history.Results), 129)
	}
}
