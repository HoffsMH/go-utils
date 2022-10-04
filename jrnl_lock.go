package util

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

// JrnlLock dir
// goes through every file
// arg is the directory with split mds
func JrnlLock(relPath string) string {
  //
  fmt.Println("locking")
  mdPattern := "*.md"
  tarPattern := "*.tar"
  abs ,err := filepath.Abs(relPath)
  if err != nil {
    log.Fatalf("failed to jrnlLock: %s", err)
  }
  os.Chdir(abs)

  finalMdPath := filepath.Join("./", mdPattern)
  matches, err := filepath.Glob(finalMdPath)
  totalContent := Hcat(matches)

  datePrefix := time.Now().Format(time.RFC3339)

  shaText := sha(totalContent);
  tarName := datePrefix + "-" + shaText + ".tar"

  xargs := []string{"-cf", tarName}
  xargs = append(xargs, matches...)
  xargs = append(xargs, "--force-local")

  _, err = exec.Command("tar", xargs...).Output()
  if err != nil {
    log.Fatalf("failed to jrnlLock: tar command %s", err)
  }

  xargs = []string{"-u"}
  xargs = append(xargs, matches...)
  _, err = exec.Command("shred", xargs...).Output()

  finalTarPath := filepath.Join(abs, tarPattern)
  tarmatches, err := filepath.Glob(finalTarPath)

  for _, match := range tarmatches {
    _, err = exec.Command("gpg", "--encrypt", "--recipient", "matthecker@pm.me", match).Output()
    if err != nil {
      log.Fatalf("failed to jrnlLock: gpg encrypt: %s", err)
    }
    _, err = exec.Command("shred", "-u", match).Output()
  }

  return shaText
}


func sha(text string) string {
  sum := sha256.Sum256([]byte(text))
  return hex.EncodeToString(sum[:10])
}

// #!/bin/bash
//
// jrnldir=~/personal/jrnl/
// jrnlfile=~/personal/jrnl/jrnl.md
//
// pushd $jrnldir
//
//
//   # https://unix.stackexchange.com/questions/100871/in-a-bash-if-condition-how-to-check-whether-any-files-matching-a-simple-wildcar
//   for file in $(fd -g '*.md'); do
//     if [[ -f "$file" ]]; then
//       # get sum of all md files in dir
//
//     fi
//   done
//
//   y=*.md
//   x=$(cat *.md | sha256sum)
//   sum=${x::10}
//   dateprefix=$(parsedate "")
//   tarname="$dateprefix-$sum.tar"
//
//   # form tarname
//   echo "form tarname: $tarname"
//   tar -cf $tarname --force-local $(fd -g '*.md')
//
//   if [[ -f $tarname ]]
//   then
//     shred -u $(fd -g '*.md')
//   fi
//
//   for tarfile in $(fd -g '*.tar'); do
//     if [[ -f "$tarfile" ]]; then
//
//       echo "looking at tarfile: $tarfile"
//       gpg --encrypt --recipient "matthecker@pm.me" "$tarfile"
//       shred -u "$tarfile"
//     fi
//   done
// popd
//
