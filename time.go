package parkrunparser

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func ParseDate(s string) (time.Time, error) {
	return time.Parse("02/01/2006", s)
}

func ParseDuration(s string) (time.Duration, error) {
	split := strings.Split(s, ":")
	splitLen := len(split)
	if splitLen != 2 && splitLen != 3 {
		return 0, fmt.Errorf("cannot parse duration '%s'", s)
	}

	t := time.Duration(0)

	// hh:mm::ss
	if splitLen == 3 {
		if value, err := strconv.Atoi(split[0]); err != nil {
			return 0, fmt.Errorf("cannot parse duration '%s': %w", s, err)
		} else {
			t += time.Duration(value) * time.Hour
		}
	}

	// mm::ss
	if value, err := strconv.Atoi(split[splitLen-2]); err != nil {
		return 0, fmt.Errorf("cannot parse duration '%s': %w", s, err)
	} else {
		t += time.Duration(value) * time.Minute
	}
	if value, err := strconv.Atoi(split[splitLen-1]); err != nil {
		return 0, fmt.Errorf("cannot parse duration '%s': %w", s, err)
	} else {
		t += time.Duration(value) * time.Second
	}

	return t, nil
}
