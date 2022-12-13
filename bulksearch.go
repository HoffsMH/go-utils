package util

import (
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

var getAbs = filepath.Abs

type Result struct {
	Term      string
	FileCount int
}

func Search(searchTerms []string, dirs []string) []Result {
	results := []Result{}
	for _, term := range searchTerms {
		result := Result{Term: term, FileCount: 0}
		for _, subTerm := range strings.Split(term, " ") {
			for _, dir := range dirs {
				if len(subTerm) > 0 {
					result.FileCount += SearchFor(subTerm, dir)
				}
			}
		}
		results = append(results, result)
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].FileCount > results[j].FileCount
	})
	return results
}

func SearchFor(term string, dir string) int {
	modTerm := strings.Replace(term, "-", "\\-", -1)

	out, _ := exec.Command("rg", "-IcF", modTerm, dir).Output()
	return len(strings.Split(string(out), "\n")) - 1
}
