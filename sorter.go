package util

import (
  "fmt"
  "io/ioutil"
  "os"
  "io/fs"
)

var kb int64 = 1000
var mb int64 = 1000 * kb
var limit int64 = 500 * mb

func HeyLook() {
  var dirs []fs.FileInfo

  if len(os.Args) > 0 {
    dirs, _ = ioutil.ReadDir(os.Args[1])
  } else {
    dirs, _ = ioutil.ReadDir(".")
  }

	for _, file := range dirs {
    if file.Size() > limit {
      fmt.Println("big boi")
      fmt.Println(file.Size())
      fmt.Println(file.Name())
    }
	}
}
