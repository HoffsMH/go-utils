package util

import (
	"time"
)

func ParseDateISO(str string) string {
	t, _ := parseDateFileName(str)

	return t.Format(time.RFC3339)
}
