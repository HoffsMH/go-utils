package util

import (
	"testing"
)

var colly = "ok"

func TestAppendTitleError(t *testing.T) {
	AppendTitle([]string{"https://google.com"})
}
