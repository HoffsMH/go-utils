package util

import (
	"bufio"
	"fmt"
	"os"
)

type OsInterface interface {
  WriteFile(name string, data []byte, perm os.FileMode) error
}

type OsWrapper struct{}

func (o *OsWrapper) WriteFile(name string, data []byte, perm os.FileMode) error {
    return os.WriteFile(name, data, perm)
}

// GetTerms ...
func GetTerms(terms []string) []string {
	if len(terms) > 0 {
		return terms
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		terms = append(terms, scanner.Text())
	}

	return terms
}

// PrintList ...
func PrintList(list []string) {
	for _, str := range list {
		fmt.Println(str)
	}
}

// Stdin ...
func Stdin() string {
	var content string

	fi, err := os.Stdin.Stat()
	if err == nil && (fi.Mode()&os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			content += scanner.Text()
		}
		return content
	}
	return ""
}

func StdinLines() []string {
    var lines []string
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines
}

func readFile(filename string) (string, error) {
	bytes, err := os.ReadFile(filename)
	return string(bytes), err
}

func truncateString(str string, length int) string {
	if len(str) > length {
		return str[:length-1]
	}
	return str
}

