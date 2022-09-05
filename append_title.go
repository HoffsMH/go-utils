package util

import (
	"regexp"
	"strings"

	"github.com/gocolly/colly/v2"
)

func AppendTitle(filepaths []string) string {
	content := Hcat(filepaths)
	lines := strings.Split(content, "\n")

	// (?m) turns on multiline mode
	r := regexp.MustCompile("(?m)(^http.*$)")

	newLines := []string{}

	for _, line := range lines {
		result := r.FindStringSubmatch(line)

		if len(result) > 0 && !AlreadyTitled(newLines) {
			c := colly.NewCollector()

			// Find and visit all links
			c.OnHTML("title", func(e *colly.HTMLElement) {
				if AlreadyTitled(newLines) {
					return
				}

				newLines = append(newLines, FormatTitle(e.Text))

			})

			c.OnError(func(_ *colly.Response, err error) {
				if AlreadyTitled(newLines) {
					return
				}

				newLines = append(newLines, "( error )")
			})

			c.Visit(line)
		}

		newLines = append(newLines, line)
	}

	return strings.Join(newLines, "\n")
}

func AlreadyTitled(newLines []string) bool {
	testTitle := regexp.MustCompile("^\\(.*\\)")
	alreadyTitled := testTitle.FindStringSubmatch(newLines[len(newLines)-1])

	return len(alreadyTitled) > 0
}

func FormatTitle(text string) string {
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "\t", "", -1)

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
