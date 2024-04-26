package parkrunparser

import (
	"fmt"
	"regexp"
	"strings"
)

type Sex int

const (
	SEX_UNKNOWN = iota
	SEX_FEMALE
	SEX_MALE
)

func (s Sex) String() string {
	switch s {
	case SEX_MALE:
		return "male"
	case SEX_FEMALE:
		return "female"
	}
	return "unknown"
}

type AgeGroup struct {
	Name string
	Age  string
	Sex  Sex
}

var reAgeGroup = regexp.MustCompile(`^[A-Z]?([fFnKkNvVwWhHmM])((?:\d+-\d+)|(?:\d+)|(?:---))$`)
var reAgeGroupWC = regexp.MustCompile(`^([fFnKkNvVwWhHmM])WC$`)

func getSex(s string) Sex {
	if strings.Contains("hHmM", s) {
		return SEX_MALE
	}
	if strings.Contains("fFKknNvVwW", s) {
		return SEX_FEMALE
	}
	return SEX_UNKNOWN
}

func ParseAgeGroup(s string) (AgeGroup, error) {
	if s == "" {
		return AgeGroup{s, "??", SEX_UNKNOWN}, nil
	}
	if match := reAgeGroup.FindStringSubmatch(s); match != nil {
		return AgeGroup{s, match[2], getSex(match[1])}, nil
	}
	if match := reAgeGroupWC.FindStringSubmatch(s); match != nil {
		return AgeGroup{s, "WC", getSex(match[1])}, nil
	}

	return AgeGroup{s, "??", SEX_UNKNOWN}, fmt.Errorf("unknown age group: %s", s)
}
