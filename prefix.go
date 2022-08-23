package util

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/araddon/dateparse"
)

var parseDate = dateparse.ParseAny

func nowISODate() string {
	return time.Now().Format(time.RFC3339)
}

func prependCurrentISODate(str string) string {
	return nowISODate() + "-" + str
}

// given a file -- prefix it by moving it
func PrefixFiles(filepaths []string) {
	for _, name := range filepaths {
		oldabs, _ := filepath.Abs(name)
		oldbasename := path.Base(oldabs)
		dir := filepath.Dir(oldabs)

		_, err := parseDateFileName(oldbasename)

		if err != nil {
			newbasename := prependCurrentISODate(oldbasename)
			newabs := filepath.Join(dir, newbasename)
			fmt.Println(newabs)
      os.Rename(oldabs, newabs)
		} 
  }
}

// given a string -- outputs a filepath prefixed with current date
func PrefixNames(filepaths []string) {
	for _, name := range filepaths {
		oldabs, _ := filepath.Abs(name)
		oldbasename := path.Base(oldabs)
		dir := filepath.Dir(oldabs)

		_, err := parseDateFileName(oldbasename)

		if err != nil {
			newbasename := prependCurrentISODate(oldbasename)
			newabs := filepath.Join(dir, newbasename)
			fmt.Println(newabs)
		} else {
			fmt.Println(filepath.Join(dir, oldbasename))
		}
	}
}

func parseDateFileName(fn string) (time.Time, error) {
	if len(fn) < 10 {
		return time.Now(), errors.New("not long enough to contain a date")
	}

  datelengths := []int{25,10,7,4}

  for _,dl := range datelengths {
    if len(fn) >= dl {
      datePortion := fn[:dl]
      dateOutput, err := parseDate(datePortion)
      if err == nil {
        return dateOutput, nil
      }
    }
  }

	return time.Now(), errors.New("No date detected")
}
