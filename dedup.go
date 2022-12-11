package util

import "regexp"

func DeDup(lines []string, pattern string) []string {
  x := regexp.MustCompile(pattern)

  // go through each line
  // for _, line := range lines {
    // if there is a pattern input and if that pattern input matches this line
    // if x.match(line) {}
  // }
      // if there is an entry in the map for this line refrain from adding to the result
      // otherwise add this line to the map

  // return the result
  return []string{"hi"}
}
