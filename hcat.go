package util

import (
	"log"
	"path"
	"path/filepath"
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

// # ls $somedir | tfilter | sort | hcat > $somefile

// Hcat ...
func Hcat(filenames []string) string {
	var text string
	for _, filename := range filenames {
		abs, err := filepath.Abs(filename)
		if err != nil {
			log.Fatalf("failed to Hcat: %s", err)
		}

		basename := path.Base(abs)
		content, err := readFile(abs)
		if err != nil {
			log.Fatalf("failed to Hcat: %s", err)
		}

		text += ensureNewline(plainTextHeading + " " + basename)
		text += ensureNewline(content)
	}

	return text
}
