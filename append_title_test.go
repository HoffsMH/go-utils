package util

import (
	"testing"

	"github.com/gocolly/colly/v2"
)

type MockCollector struct {
	onHTMLCallback colly.HTMLCallback
}

func (mc *MockCollector) OnHTML(goqueryMatcher string, f colly.HTMLCallback) {
	mc.onHTMLCallback = f
}

func (mc *MockCollector) Visit(URL string) error {
	mc.onHTMLCallback(&colly.HTMLElement{
		Text: "Google",
	})
	return nil
}

func (mc *MockCollector) OnError(f colly.ErrorCallback) {
}

func TestAppendTitle(t *testing.T) {
	content := []string{"https://google.com"}
	expected := "( Google )\nhttps://google.com"
	mockCollector := &MockCollector{}

	result := AppendTitle(mockCollector, content)
	if result != expected {
		t.Errorf("AppendTitle(%v) = %v; expected %v", content, result, expected)
	}
}
