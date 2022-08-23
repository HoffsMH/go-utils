package util

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

// given a file -- prefix it by moving it
func SuffixFiles(filepaths []string) {
	for _, name := range filepaths {
		oldabs, _ := filepath.Abs(name)
		oldbasename := path.Base(oldabs)
		dir := filepath.Dir(oldabs)

		_, err := parseDateFileName(oldbasename)

		if err != nil {
			newbasename := prependCurrentISODate(oldbasename)
			newabs := filepath.Join(dir, newbasename)
			fmt.Println(newabs)
      os.Rename(oldabs, newabs)
		} 
  }
}

// given a string -- outputs a filepath prefixed with current date
func SuffixNames(filepaths []string) {
	for _, name := range filepaths {
		oldabs, _ := filepath.Abs(name)
		oldbasename := path.Base(oldabs)
		dir := filepath.Dir(oldabs)

		_, err := parseDateFileName(oldbasename)

		if err != nil {
			newbasename := prependCurrentISODate(oldbasename)
			newabs := filepath.Join(dir, newbasename)
			fmt.Println(newabs)
		} else {
			fmt.Println(filepath.Join(dir, oldbasename))
		}
	}
}
