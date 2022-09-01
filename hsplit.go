package util

import (
  "path/filepath"
  "regexp"
  "path"
  "os"
)

type FileContent struct {
	Name    string
	Content string
	Dir     string
}

func Hsplit(lines []string, dir string) {
	var result []FileContent

	r := regexp.MustCompile(heading + " (.*)")

  for _, line := range lines {
		match := r.FindStringSubmatch(line)

    // if there is atleast one heading already and 
    // its not a match
		if len(result) > 0 && len(match) == 0 {
			result[0].Content += line + "\n"
		}

		// we found a header on the current line
		if len(match) > 1 {
			// start a new header
			result = append([]FileContent{newFileContent(match[1], dir)}, result...)
		}
  }
  result = pruneEmptyFileContents(result)

  WriteSplits(result)
}

func pruneEmptyFileContents(fcs []FileContent) []FileContent {
	var pruned []FileContent
	r := regexp.MustCompile(".")

	for _, fc := range fcs {
		if match := r.FindStringSubmatch(fc.Content); len(match) > 0 {
			pruned = append(pruned, fc)
		} else {
			os.Remove(filepath.Join(fc.Dir, fc.Name))
		}
	}

	return pruned
}

func newFileContent(name string, dir string) FileContent {
	dir, _ = filepath.Abs(dir)
	if _, err := parseDateFileName(name); err != nil {
		name = prependCurrentISODate(name)
	}

	return FileContent{
		Dir:     dir,
		Name:    path.Base(name),
		Content: "",
	}
}

func WriteSplits(fcs []FileContent) {
	for _, fc := range fcs {
		os.WriteFile(path.Join(fc.Dir, fc.Name), []byte(fc.Content), 0644)
	}
}
