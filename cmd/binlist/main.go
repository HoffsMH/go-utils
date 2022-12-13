package main

import (
	"fmt"
	"os"
	"strings"
)

type void struct {}
var member void

func main() {

	pathstring := os.Getenv("PATH")
	paths := strings.Split(pathstring, ":")
	binset := make(map[string]void)

	for _, path := range paths {
		entries, _ := os.ReadDir(path)
		for _, entry := range entries {
			if !entry.IsDir() {
				info, _ := entry.Info()

				if IsExecAny(info.Mode()) {
					binset[entry.Name()] = member
				}
			}
		}
	}

	for k := range binset {
		fmt.Println(k)
	}
}

func IsExecAny(mode os.FileMode) bool {
	return mode&0111 != 0
}
