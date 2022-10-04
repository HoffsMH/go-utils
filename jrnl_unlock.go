package util

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// JrnlUnlock
// decrypt gpg
// extract archive
func JrnlUnlock(relPath string) {
  decryptGpgs(relPath)
  unpackTars(relPath)
}

func decryptGpgs(relPath string) {
  abs ,err := filepath.Abs(relPath)
  os.Chdir(abs)
  if err != nil {
    log.Fatalf("failed to JrnlUnlock: %s", err)
  }
  gpgPattern := "*.gpg"
  finalPath := filepath.Join(abs, gpgPattern)
  matches, err := filepath.Glob(finalPath)
  if err != nil {
    log.Fatalf("failed to JrnlUnlock: %s", err)
  }
  for _, match := range matches {
    fmt.Println(match)
    _, err := exec.Command("gpg", "--decrypt", "--use-embedded-filename", match).Output()
    if err != nil {
      log.Fatalf("failed to JrnlUnlock gpg: %s", err)
    }

    _, err = exec.Command("shred", "-u", match).Output()
    if err != nil {
      log.Fatalf("failed to JrnlUnlock shred: %s", err)
    }
  }
}

func unpackTars(relPath string) {
  abs ,err := filepath.Abs(relPath)
  os.Chdir(abs)
  if err != nil {
    log.Fatalf("failed to JrnlUnlock: %s", err)
  }
  tarPattern := "*.tar"
  finalPath := filepath.Join(abs, tarPattern)
  matches, err := filepath.Glob(finalPath)
  if err != nil {
    log.Fatalf("failed to JrnlUnlock: %s", err)
  }
  for _, match := range matches {
    fmt.Println(match)
    _, err := exec.Command("tar", "-xvf", match, "--force-local").Output()
    if err != nil {
      log.Fatalf("failed to JrnlUnlock tar: %s", err)
    }

    _, err = exec.Command("shred", "-u", match).Output()
    if err != nil {
      log.Fatalf("failed to JrnlUnlock shred: %s", err)
    }
  }
}

// for i in $gpgs
// do
//   echo "looking at gpgfile: $i"
//   gpg --decrypt --use-embedded-filename "$i" && shred -u "$i"
//   tar -xkvf "${i/.gpg}" --force-local
//   shred -u "${i/.gpg}"
// done
//
// fd -g '*.md' | tfilter -w 1 | hcat > tmp
// mv tmp $jrnlfile
