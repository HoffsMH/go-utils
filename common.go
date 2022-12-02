package util

import (
	"os"
)

func readFile(filename string) (string, error) {
	bytes, err := os.ReadFile(filename)
	return string(bytes), err
}
