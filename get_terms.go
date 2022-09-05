package util

import (
	"bufio"
	"os"
)

func GetTerms(args []string) []string {
	terms := []string{}

	if len(args) > 0 {
		terms = args
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			terms = append(terms, scanner.Text())
		}
	}

	return terms
}
