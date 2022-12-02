package util

import (
	"bufio"
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

// Stdin ...
func Stdin() []string {
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
