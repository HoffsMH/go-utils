package util

import (
	"io/fs"
	"io/ioutil"
	"os"
)

var kb int64 = 1000
var mb int64 = 1000 * kb
var limit int64 = 500 * mb

func SFilter(limit int64) []string {
	var dirs []fs.FileInfo
  var results = []string{}

	if len(os.Args) > 0 {
		dirs, _ = ioutil.ReadDir(os.Args[1])
	} else {
		dirs, _ = ioutil.ReadDir(".")
	}

	for _, file := range dirs {
		if file.Size() > limit {
      results = append(results, file.Name())
		}
	}
  return results
}
