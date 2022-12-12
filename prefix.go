package util

import (
	"errors"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/araddon/dateparse"
)

// current time in iso format
func nowISODate() string {
	return time.Now().Format(time.RFC3339)
}

// appends current iso time to any string
func prependCurrentISODate(str string) string {
	return nowISODate() + "-" + str
}

// given a file -- prefix it by moving it
func PrefixFiles(filepaths []string) []string {
  results := []string{}
	for _, name := range filepaths {
		oldabs, _ := filepath.Abs(name)
		oldbasename := path.Base(oldabs)
		dir := filepath.Dir(oldabs)

		_, err := parseDateFileName(oldbasename)

		if err != nil {
			newbasename := prependCurrentISODate(oldbasename)
			newabs := filepath.Join(dir, newbasename)
      results = append(results, newabs)
			os.Rename(oldabs, newabs)
		}
	}
  return results
}

// given a string -- outputs a filepath prefixed with current date
func PrefixNames(filepaths []string) []string {
  results := []string{}
	for _, name := range filepaths {
		oldabs, _ := filepath.Abs(name)
		oldbasename := path.Base(oldabs)
		dir := filepath.Dir(oldabs)

		_, err := parseDateFileName(oldbasename)

		if err != nil {
			newbasename := prependCurrentISODate(oldbasename)
			newabs := filepath.Join(dir, newbasename)
      results = append(results, newabs)
		} else {
      results = append(results, filepath.Join(dir, oldbasename))
		}
	}
  return results
}

// if a given filename begins with a parsable date extract that date otherwise
// error
func parseDateFileName(fn string) (time.Time, error) {
	if len(fn) < 10 {
		return time.Now(), errors.New("not long enough to contain a date")
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

	return time.Now(), errors.New("No date detected")
}
