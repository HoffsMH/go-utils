package util

import (
	"errors"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/jmhodges/clock"

	"github.com/araddon/dateparse"
)

type PrefixOptions struct {
  Clock clock.Clock
  Format string
}

func nowFormat(...opts PrePrefixOptions) {
  return opts.clock.Now().Format(opts.Format)
}
// current time in iso format
func nowISODate(clock clock.Clock) string {
  return clock.Now().Format(time.RFC3339)
}

// appends current iso time to any string
func prependCurrentISODate(str string, clock clock.Clock) string {
  return nowISODate(clock) + "-" + str
}

func nowDate(clock clock.Clock) string {
  return clock.Now().Format("2006-01-02")
}

func prependCurrentDate(str string, clock clock.Clock) string {
  return nowDate(clock) + "-" + str
}


// given a string -- outputs a filepath prefixed with current date
func PrefixNamesISO(filepaths []string, opts ...Options) []string {
	results := []string{}
	for _, name := range filepaths {
    PrefixNameISO(name, opts...)
	}
	return results
}

func PrefixNameISO(str string, opts ...Options) string {
    // Default values
    clock := clock.New();

    // Override if passed
    if len(opts) > 0 {
        clock = opts[0].Clock
    }

    oldabs, _ := filepath.Abs(str)
    oldbasename := path.Base(oldabs)
    dir := filepath.Dir(oldabs)

    _, err := parseDateFileName(oldbasename)

    if err != nil {
        newbasename := prependCurrentISODate(oldbasename, clock)
        newabs := filepath.Join(dir, newbasename)

        return newabs
    }
    return filepath.Join(dir, oldbasename)
}

// given a string -- outputs a filepath prefixed with current date
func PrefixNamesDate(filepaths []string, opts ...Options) []string {
	results := []string{}
	for _, name := range filepaths {
    PrefixName(name, opts...)
	}
	return results
}

func PrefixNameDate(str string, opts ...Options) string {
    // Default values
    clock := clock.New();

    // Override if passed
    if len(opts) > 0 {
        clock = opts[0].Clock
    }

    oldabs, _ := filepath.Abs(str)
    oldbasename := path.Base(oldabs)
    dir := filepath.Dir(oldabs)

    _, err := parseDateFileName(oldbasename)

    if err != nil {
        newbasename := prependCurrentDate(oldbasename, clock)
        newabs := filepath.Join(dir, newbasename)

        return newabs
    }
    return filepath.Join(dir, oldbasename)
}

// if a given filename begins with a parsable date extract that date otherwise
// error
func parseDateFileName(fn string) (time.Time, error) {
	if len(fn) < 10 {
		return time.Time{}, errors.New("not long enough to contain a date")
	}
	base := filepath.Base(fn)
	datelengths := []int{25, 10, 7, 4}

	for _, dl := range datelengths {
		if len(base) >= dl {
			datePortion := base[:dl]
			dateOutput, err := dateparse.ParseAny(datePortion)
			if err == nil {
				return dateOutput, nil
			}
		}
	}

	return time.Time{}, errors.New("No date detected")
}
