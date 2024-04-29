package parkrunparser

import (
	"compress/gzip"
	"io"
	"os"
	"strings"
	"testing"
	"time"
)

func readFile(fileName string) ([]byte, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	if strings.HasSuffix(fileName, ".gz") {
		fz, err := gzip.NewReader(f)
		if err != nil {
			return nil, err
		}
		defer fz.Close()

		return io.ReadAll(fz)
	} else {
		return io.ReadAll(f)
	}
}

func loadParkrunSimple(t *testing.T, fileName string) *Results {
	data, err := readFile(fileName)
	if err != nil {
		t.Errorf("failed to read file '%s': %v", fileName, err)
		return nil
	}

	event, err := ParseResults(data)
	if err != nil {
		t.Errorf("%s: failed to parse results: %v", fileName, err)
		return nil
	}

	return &event
}

func loadParkrun(t *testing.T, fileName string, index int, dateStr string, finishers int, volunteers int) *Results {
	event := loadParkrunSimple(t, fileName)
	if event == nil {
		return nil
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		t.Errorf("%s: cannot parse date '%s': %v", fileName, dateStr, err)
		return nil
	}

	if event.Index != index {
		t.Errorf("%s: unexpected results index: %d, expected: %d", fileName, event.Index, index)
		return nil
	}
	if event.Date != date {
		t.Errorf("%s: unexpected results date: %v, expected: %s", fileName, event.Date, dateStr)
		return nil
	}

	if len(event.Finishers) != finishers {
		t.Errorf("%s: unexpected number of finishers: %d, expected: %d", fileName, len(event.Finishers), finishers)
		return nil
	}

	if len(event.Volunteers) != volunteers {
		t.Errorf("%s: unexpected number of volunteers: %d, expected: %d", fileName, len(event.Volunteers), volunteers)
		return nil
	}

	return event
}

func TestParse(t *testing.T) {
	event := loadParkrun(t, "test-data/de-dietenbach-53.gz", 53, "2022-11-05", 20, 9)
	if event.Finishers[0].Name != "Daniel THOMA" {
		t.Errorf("unexpected name of first finisher: %s, expected: %s", event.Finishers[0].Name, "Daniel THOMA")
	}

	loadParkrun(t, "test-data/dk-amagerstrandpark-560.gz", 560, "2024-03-16", 45, 4)
	loadParkrun(t, "test-data/fi-tokoinranta-165.gz", 165, "2024-04-13", 79, 12)
	loadParkrun(t, "test-data/fr-rouen-189.gz", 189, "2022-07-02", 38, 5)
	loadParkrun(t, "test-data/it-etna-196.gz", 196, "2024-04-20", 25, 7)
	loadParkrun(t, "test-data/jp-chuokoen-89.gz", 89, "2024-02-24", 24, 5)
	loadParkrun(t, "test-data/nl-depotten-52.gz", 52, "2024-04-20", 20, 7)
	loadParkrun(t, "test-data/no-nansenparken-108.gz", 108, "2024-04-06", 45, 7)
	loadParkrun(t, "test-data/pl-rumia-419.gz", 419, "2024-04-13", 39, 9)
	loadParkrun(t, "test-data/se-broparken-84.gz", 84, "2024-04-20", 24, 7)
	loadParkrun(t, "test-data/uk-hamsterleyforest-62.gz", 62, "2023-03-18", 24, 14)
}
