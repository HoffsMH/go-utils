package util

import (
	"log"
	"os"
)

func SFilter(fileNames []string, lowerLimit int, upperLimit int) []string {
	var results = []string{}

	for _, fileName := range fileNames {
    file, err := os.Stat(fileName)
    if err != nil {
      log.Fatalf("no such file %q", fileName)
    }

    if lowerLimit > 0 && upperLimit > 0 {
      if file.Size() > int64(lowerLimit) && file.Size() < int64(upperLimit) {
        results = append(results, fileName)
        continue
      }
    }
    if lowerLimit > 0 {
      if file.Size() > int64(lowerLimit) {
        results = append(results, fileName)
        continue
      }
    }

    if upperLimit > 0 {
      if file.Size() < int64(upperLimit) {
        results = append(results, fileName)
        continue
      }
    }
	}
	return results
}
