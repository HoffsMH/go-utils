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
func AppendTitle(content []string) string {
	newContent := []string{}

	for _, line := range content {
		result := linkTest.FindStringSubmatch(line)

		if len(result) > 0 && !AlreadyTitled(newContent) {
			c := colly.NewCollector()

			// Find and visit all links
			c.OnHTML("title", func(e *colly.HTMLElement) {
				if AlreadyTitled(newContent) {
					return
				}

				newContent = append(newContent, FormatTitle(e.Text))
			})

			c.OnError(func(_ *colly.Response, err error) {
				if AlreadyTitled(newContent) {
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

func AlreadyTitled(newLines []string) bool {
	if len(newLines) == 0 {
		return false
	}
	testTitle := regexp.MustCompile("^\\(.*\\)")
	alreadyTitled := testTitle.FindStringSubmatch(newLines[len(newLines)-1])

	return len(alreadyTitled) > 0
}

func FormatTitle(text string) string {
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "\t", "", -1)

	// Some HTML titles can be arbitrarily long
	if len(text) > 75 {
		text = "( " + text[:74] + " )"
	} else {
		text = "( " + text + " )"
	}

	return removeDuplicateWhitespace(text)
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
