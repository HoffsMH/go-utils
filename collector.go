package util

import "github.com/gocolly/colly/v2"

type Collector interface {
	OnHTML(goqueryMatcher string, f colly.HTMLCallback)
	Visit(URL string) error
	OnError(f colly.ErrorCallback)
}

type RealCollector struct {
	*colly.Collector
}

func (rc *RealCollector) OnHTML(goqueryMatcher string, f colly.HTMLCallback) {
	rc.Collector.OnHTML(goqueryMatcher, f)
}

func (rc *RealCollector) OnError(f colly.ErrorCallback) {
	rc.Collector.OnError(f)
}

func (rc *RealCollector) Visit(URL string) error {
	return rc.Collector.Visit(URL)
}
