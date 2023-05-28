package util

import (
	"errors"
	"path"
	"path/filepath"
	"time"

	"github.com/jmhodges/clock"

	"github.com/araddon/dateparse"
)

type Prefixer struct {
  Clock clock.Clock
  Format string
}

func NewPrefixer(opts ...interface{}) (*Prefixer) {
  newClock := clock.New();
  if len(opts) > 0 {
    newClock = opts[0].(clock.Clock)
  }

  // "2006-01-02"
  format := time.RFC3339
  if len(opts) > 1 {
    format = opts[1].(string)
  }

  return &Prefixer {
    Clock: newClock,
    Format: format,
  }
}

func (p *Prefixer) nowFormat() string {
  return p.Clock.Now().Format(p.Format)
}

// appends current iso time to any string
func (p *Prefixer) prependCurrentDate(str string) string {
  return p.nowFormat() + "-" + str
}

// given a string -- outputs a filepath prefixed with current date
func (p *Prefixer) Names(filepaths []string) []string {
	results := []string{}
	for _, name := range filepaths {
    results = append(results, p.Name(name))
	}
	return results
}

func (p *Prefixer) Name(str string) string {
    oldabs, _ := filepath.Abs(str)
    oldbasename := path.Base(oldabs)
    dir := filepath.Dir(oldabs)

    _, err := parseDateFileName(oldbasename)

    if err != nil {
        newbasename := p.prependCurrentDate(oldbasename)
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
