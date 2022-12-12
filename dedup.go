package util

import (
	"regexp"
)

func DeDup(lines []string, before string, pattern string) []string {
  patternMatch := regexp.MustCompile(pattern)
  result := []string{}
  seen := make(map[string]int)

  // go through each line
  for i, line := range lines {
    match := patternMatch.FindStringSubmatch(line)
    // if there is a pattern input and if that pattern input matches this line
		if len(match) != 0 {
      seen[line] += 1
		}

    // if we have seen this line before
    if seen[line] > 1 {
      // trim prior line if specified
      result = trimBeforeLine(lines[i-1], result, before)
      continue
    }
		result = append(result, line)
  }

  // return the result
  return result
}

func trimBeforeLine(priorLine string, pendingResult []string, before string) []string {
  if before == "" { return pendingResult }
  beforeMatch := regexp.MustCompile(before)

  match := beforeMatch.FindStringSubmatch(priorLine)
  if len(match) != 0 {
    pendingResult = pendingResult[:len(pendingResult) - 1]
  }
  return pendingResult;
}
