package util

import (
	"testing"
)

type FakeTitleGetter struct {
}

func (gt *FakeTitleGetter)GetTitle(s string) string {
  return "Google"
}

func TestAppendTitle(t *testing.T) {
	content := []string{"https://google.com"}
	expected := "( Google )\nhttps://google.com"
	appender := &HtmlTitleAppender{
    TitleGetter: &FakeTitleGetter{},
  }

	result := appender.Call(content)
	if result != expected {
		t.Errorf("AppendTitle(%v) = %v; expected %v", content, result, expected)
	}
}
