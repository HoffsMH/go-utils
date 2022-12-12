package util

// #!/bin/bash
//
// capdir=$(~/personal/00-capture/time)
// capfile=$(t $(echo "$HOME/personal/00-capture/time/cap.md" | prefix name))
// selection=$(select-input $@)
//
//
// echo "$selection" | append_title >> $capfile
//
// notify-send "cap: ${selection}"
//
// echo $capfile
import (
	"log"
	"os"
	"path/filepath"
)

func Cap(capdir string) string {
  newCapFile := PrefixNames([]string{capdir})

	abs, err := filepath.Abs(capdir)
  if err != nil {
    log.Fatalf("Failed to get Abs of %s", capdir)
  }
  newCapFilePath := filepath.Join(abs, newCapFile[0])

	os.WriteFile(newCapFilePath, getInput(), 0644)
  // return the file name
  return newCapFilePath
}

func getInput() string {
  // Get it from clipboard or args or stdin??
  return "hi"
}
