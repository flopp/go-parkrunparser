package parkrunparser

import "fmt"

type Achievement int

const (
	AchievementNone Achievement = iota
	AchievementFirst
	AchievementPB
)

func ParseAchievement(s string) (Achievement, error) {
	if s == "" {
		return AchievementNone, nil
	}

	var first = [...]string{
		"First Timer!",          // UK, SA, CA, US, NZ, IE, MY, AUS
		"Erstläufer!",           // Germany
		"Erstteilnahme!",        // Germany
		"Première perf' !",      // France
		"Première perf&#039; !", // France
		"Prima volta!",          // Italy
		"Debut!",                // Sweden
		"Debiutant",             // Poland
		"Nieuwe loper!",         // Netherlands
		"Første gang!",          // Denmark
		"初参加!",                  // Japan
		"Ensikertalainen!",      // Finland
	}
	var pb = [...]string{
		"New PB!",                // UK, SA, CA, US, NZ, IE, MY, AUS
		"Neue PB!",               // Germany
		"Meilleure perf' !",      // France
		"Meilleure perf&#039; !", // France
		"Nuovo PB!",              // Italy
		"Nytt PB!",               // Sweden
		"Nowy PB!",               // Poland
		"Nieuw PR!",              // Netherlands
		"Ny PB!",                 // Denmark
		"自己ベスト!",                 // Japan
		"Oma ennätys!",           // Finland
	}

	for _, pattern := range first {
		if pattern == s {
			return AchievementFirst, nil
		}
		if fmt.Sprintf("[parkrun_translate phrase='%s']", pattern) == s {
			return AchievementFirst, nil
		}
	}
	for _, pattern := range pb {
		if pattern == s {
			return AchievementPB, nil
		}
		if fmt.Sprintf("[parkrun_translate phrase='%s']", pattern) == s {
			return AchievementPB, nil
		}
	}

	return AchievementNone, fmt.Errorf("cannot parse achievement: %s", s)
}
