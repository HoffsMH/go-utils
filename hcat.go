package util

import (
	"fmt"
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

func Hcat(filenames []string) {
  var text string
  for _, filename := range filenames {
    abs, _ := filepath.Abs(filename)
		basename := path.Base(abs)
    content, _ := readFile(abs)

    text += plainTextHeading + " " + basename + "\n"
    text += ensureNewline(string(content))

  }
  fmt.Println(text)
}
