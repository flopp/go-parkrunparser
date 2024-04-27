package parkrunparser

import (
	"testing"
	"time"
)

func TestParseDuration(t *testing.T) {
	for _, test := range []struct {
		input    string
		duration time.Duration
		err      bool
	}{
		{"1:23:45", 1*time.Hour + 23*time.Minute + 45*time.Second, false},
		{"12:34:56", 12*time.Hour + 34*time.Minute + 56*time.Second, false},
		{"12:34", 12*time.Minute + 34*time.Second, false},
		{"1:23", 1*time.Minute + 23*time.Second, false},
		{"BAD", 0, true},
		{"12", 0, true},
		{"12:34:56:78", 0, true},
	} {
		d, e := ParseDuration(test.input)
		if test.err {
			if e == nil {
				t.Errorf("duration '%s': expected error not emitted", test.input)
			}
		} else {
			if e != nil {
				t.Errorf("duration '%s': unexpected error: %v", test.input, e)
			}
			if d != test.duration {
				t.Errorf("duration '%s': unexpected duration: %s expected: %s", test.input, d, test.duration)
			}
		}
	}
}
