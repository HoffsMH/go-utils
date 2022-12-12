package util

import (
	"bufio"
	"fmt"
	"os"
)

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

func PrinList(list []string) {
  for _, str := range list {
    fmt.Println(str)
  }
}

// Stdin ...
func Stdin() []string {
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
