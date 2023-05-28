package util

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

// (?m) turns on multiline mode
var linkTest *regexp.Regexp = regexp.MustCompile("(?m)(^http.*$)")

type HtmlTitleAppender struct {
	TitleGetter
}

type TitleGetter interface {
	GetTitle(line string) string
}

type GetTitleWrapper struct {}

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
func (a *HtmlTitleAppender) Call(content []string) string {
	newContent := []string{}

	for _, line := range content {
		result := linkTest.FindStringSubmatch(line)

		if len(result) > 0 && !isAlreadyTitled(newContent) {
      title := a.GetTitle(line)
      newContent = append(newContent, FormatTitle(title))
		}

		newContent = append(newContent, line)
	}

	return strings.Join(newContent, "\n")
}

func (gt *GetTitleWrapper) GetTitle(url string) string {
  html, _ := getHtml(url)
  return getTitle(html)
}

func getHtml(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("HTTP error:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Read error:", err)
		return nil, err
	}

	doc, err := html.Parse(strings.NewReader(string(body)))
	if err != nil {
		log.Println("Parse error:", err)
		return nil, err
	}

	return doc, nil
}

func getTitle(n *html.Node) string {
	if n.Type == html.ElementNode && n.Data == "title" {
		return n.FirstChild.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
    title := getTitle(c)
		if title != "" {
			return title
		}
	}
	return "No title Found"
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
