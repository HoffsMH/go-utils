package util

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"

	"github.com/jmhodges/clock"
)

type FileContent struct {
	Name    string
	Content string
	Dir     string
}

type Hsplitter struct {
	Clock clock.Clock
	Os    OsInterface
}

func NewHsplitter(opts ...interface{}) *Hsplitter {
	newClock := clock.New()
	if len(opts) > 0 {
		newClock = opts[0].(clock.Clock)
	}

	return &Hsplitter{
		Clock: newClock,
		Os:    &OsWrapper{},
	}
}

// given text that contains some lines that begin with ## (h2 heading in md)
// split every heading into its own file in a given directory
//
// # cat $somefile | hsplit $somdir
func (h *Hsplitter) Call(lines []string, dir string) {
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
			result = append([]FileContent{h.newFileContent(match[1], dir)}, result...)
		}
	}
	result = compressDuplicateFileContents(result)
	result = pruneEmptyFileContents(result)

	h.WriteSplits(result)
}

func compressDuplicateFileContents(fcs []FileContent) []FileContent {
  var compressed []FileContent

  // proceed in reverse order
  for i := len(fcs) - 1; i >= 0; i-- {
    found := false
    fc := fcs[i]
    for i, c := range compressed {
      if c.Name == fc.Name {
        found = true
        compressed[i].Content += fc.Content
      }
    }

    if !found {
      compressed = append(compressed, fc)
    }
  }

  return compressed
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

func (h *Hsplitter) newFileContent(name string, dir string) FileContent {
	dir, _ = filepath.Abs(dir)

	return FileContent{
		Dir:     dir,
		Name:    name,
		Content: "",
	}
}

func (h *Hsplitter) WriteSplits(fcs []FileContent) {
	for _, fc := range fcs {
    filePath := path.Join(fc.Dir, fc.Name)

    dir := filepath.Dir(filePath)
    if err := h.Os.MkdirAll(dir, 0755); err != nil {
      fmt.Println("Error creating directory:", err)
    }
		h.Os.WriteFile(filePath, []byte(fc.Content), 0644)
    fmt.Println(filePath)
	}
}
