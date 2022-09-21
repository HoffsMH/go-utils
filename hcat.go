package util

import (
	"os"
	"path"
	"path/filepath"
	"regexp"
)

var plainTextHeading = "##"
var heading = "^" + plainTextHeading

func readFile(filename string) (string, error) {
	bytes, err := os.ReadFile(filename)
	return string(bytes), err
}

func ensureNewline(s string) string {
	match, _ := regexp.MatchString("\n$", s)
	if match == true {
		return s
	}
	return s + "\n"
}

// given a list of text files cats them together into a single stream of text
// separated by headings
// given file1.md that contains "hello"
// and given file2.md that contains "ok"
// output:
// "## file1.md
//  hello
//  ## file2.md
//  ok"

//  # ls $somedir | tfilter | sort -h | hcat > $somefile
func Hcat(filenames []string) string {
	var text string
	for _, filename := range filenames {
		abs, _ := filepath.Abs(filename)
		basename := path.Base(abs)
		content, _ := readFile(abs)

		text += plainTextHeading + " " + basename + "\n"
		text += ensureNewline(string(content))

	}
	return text
}
