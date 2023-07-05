package util

import (
	"log"
	"path"
	"strings"
)

var plainTextHeading = "##"
var heading = "^" + plainTextHeading

func ensureNewline(s string) string {
	if !strings.HasSuffix(s, "\n") {
		s += "\n"
	}

	return s
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

// # ls $somedir | sort | hcat > $somefile

// Hcat ...
func Hcat(filenames []string, dir string) string {
	var text string
	for _, filename := range filenames {

		content, err := readFile(path.Join(dir, filename))

		if err != nil {
			log.Fatalf("failed to Hcat: %s", err)
		}

		text += ensureNewline(plainTextHeading + " " +  filename)
		text += ensureNewline(content)
	}

	return text
}
