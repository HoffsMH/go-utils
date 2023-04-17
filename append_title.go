package util

import (
	"regexp"
	"strings"

	"github.com/gocolly/colly/v2"
)

// (?m) turns on multiline mode
var linkTest *regexp.Regexp = regexp.MustCompile("(?m)(^http.*$)")

// given a multiline string, append html title to previous lines that are http links
// given:
// foo
// http://google.com
// bar
//
// output:
// foo
// ( Google )
// http://google.com
// bar
func AppendTitle(c Collector, content []string) string {
	newContent := []string{}

	for _, line := range content {
		result := linkTest.FindStringSubmatch(line)

		if len(result) > 0 && !isAlreadyTitled(newContent) {

			// Find and visit all links
			c.OnHTML("title", func(e *colly.HTMLElement) {

				if isAlreadyTitled(newContent) {
					return
				}

				newContent = append(newContent, FormatTitle(e.Text))
			})

			c.OnError(func(_ *colly.Response, err error) {
				if isAlreadyTitled(newContent) {
					return
				}

				newContent = append(newContent, "( error )")
			})

			c.Visit(line)
		}

		newContent = append(newContent, line)
	}

	return strings.Join(newContent, "\n")
}

func isAlreadyTitled(newLines []string) bool {
	if len(newLines) == 0 {
		return false
	}
	testTitle := regexp.MustCompile("^\\(.*\\)")
	alreadyTitled := testTitle.FindStringSubmatch(newLines[len(newLines)-1])

	return len(alreadyTitled) > 0
}

func FormatTitle(text string) string {
	text = removeNewlinesAndTabs(text)
	text = truncateString(text, 75)
	text = "( " + text + " )"
	return removeDuplicateWhitespace(text)
}

func removeNewlinesAndTabs(text string) string {
	text = strings.Replace(text, "\n", "", -1)
	return strings.Replace(text, "\t", "", -1)
}

func truncateString(str string, length int) string {
	if len(str) > length {
		return str[:length-1]
	}
	return str
}

func removeDuplicateWhitespace(inputString string) string {
	newString := ""

	for i, char := range inputString {
		var prev byte

		if i != 0 {
			prev = inputString[i-1]
		}

		if byte(char) == prev && char == ' ' {
			continue
		}
		newString = newString + string(char)
	}
	return newString
}
