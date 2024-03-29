package util

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"time"
)

var mdPattern = "*.md"
var tarPattern = "*.tar"

// JrnlLock dir
// goes through every md file in rel path
// tars them
// encrypts them
func JrnlLock(relPath string) string {
	log.Println("locking...")
	abs, err := filepath.Abs(relPath)
	if err != nil {
		log.Fatalf("failed to jrnlLock abs path: %s", err)
	}

	os.Chdir(abs)

	matches, err := filepath.Glob(mdPattern)

	tar(matches, relPath)

	encrypt(abs)

	return "Done"
}
func shredFiles(matches []string) {
	var wg sync.WaitGroup
	wg.Add(len(matches))

	for _, match := range matches {
		go func(match string) {
			defer wg.Done()
			_, err := exec.Command("shred", "-u", "-n 1", match).Output()
			if err != nil {
				log.Fatalf("failed to shred file: %s", err)
			}
			log.Printf("finished shredding %s...", match)
		}(match)
	}
	wg.Wait()
}

func tar(matches []string, relPath string) {
	tarName := genTarName(matches, relPath)

	xargs := []string{"-cf", tarName}
	xargs = append(xargs, matches...)
	xargs = append(xargs, "--force-local")

	_, err := exec.Command("tar", xargs...).Output()
	if err != nil {
		log.Fatalf("failed to jrnlLock: tar command %s", err)
	}
	log.Println("finished tar...")
	log.Println("shreding md files ...")

	shredFiles(matches)
}

func genTarName(matches []string, relPath string) string {
	totalContent := Hcat(matches, relPath)
	shaText := sha(totalContent)
	datePrefix := time.Now().Format(time.RFC3339)
	return datePrefix + "-" + shaText + ".tar"
}

func encrypt(abs string) {
	finalTarPath := filepath.Join(abs, tarPattern)
	tarmatches, err := filepath.Glob(finalTarPath)

	for _, match := range tarmatches {
		log.Println("encrypting" + match + " ...")
		_, err = exec.Command("gpg", "--encrypt", "--recipient", "matthecker@pm.me", match).Output()
		if err != nil {
			log.Fatalf("failed to jrnlLock: gpg encrypt: %s", err)
		}
		log.Println("finished encrypt...")
		log.Println("shreding tar file " + match + " ...")
		_, err = exec.Command("shred", "-u", match).Output()
	}
}

func sha(text string) string {
	sum := sha256.Sum256([]byte(text))
	return hex.EncodeToString(sum[:10])
}
