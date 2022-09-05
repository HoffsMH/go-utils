package util

import (
	"path/filepath"
	"time"
)

type TimeRange struct {
	Months int
	Weeks  int
	Days   int
}

func TFilter(filnames []string, tr *TimeRange, ignore bool) []string {
	var results []string
	timeHoriz := time.Now().AddDate(-tr.Months, -tr.Weeks, -tr.Days-2)
	for _, filename := range filnames {
		abs, _ := filepath.Abs(filename)
		t, err := parseDateFileName(filename)

		if ignore && err != nil {
			continue
		}

		if t.After(timeHoriz) {
			results = append(results, abs)
		}
	}
	return results
}
