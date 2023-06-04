package util

import (
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"sync"
)

var getAbs = filepath.Abs

type Result struct {
	Term      string
	FileCount int
}

// Search searches for the given search terms in the given directories.
func Search(searchTerms []string, dirs []string) []Result {
	var wg sync.WaitGroup
	results := make([]Result, 0, len(searchTerms))
	resultChan := make(chan Result) // what is being transmitted through the
	for _, term := range searchTerms {
		wg.Add(1)
		go func(term string) {
			defer wg.Done()

			result := Result{Term: term, FileCount: 0}
			for _, subTerm := range strings.Split(term, " ") {
				for _, dir := range dirs {
					if len(subTerm) > 0 {
						result.FileCount += SearchFor(subTerm, dir)
					}
				}
			}
			// send result to channel
			resultChan <- result
		}(term)
	}
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for result := range resultChan {
		results = append(results, result)
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].FileCount > results[j].FileCount
	})
	return results
}

func SearchFor(term string, dir string) int {
	out, _ := exec.Command("rg", "-IcF", term, dir).Output()
	return len(strings.Split(string(out), "\n")) - 1
}
