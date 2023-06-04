package util

import (
	"bytes"
)

type NewContent struct {
	WordBuf          *bytes.Buffer
	SpaceBuf         *bytes.Buffer
	Value            *bytes.Buffer
	CurrentLineIndex int
	Lim              int
}

func (nc *NewContent) resetWordBuf()  {
  nc.WordBuf.Reset();
}

func (nc *NewContent) writeWordBuf() {
  nc.WordBuf.WriteTo(nc.Value)
}

func (nc *NewContent) resetSpaceBuf()  {
  nc.SpaceBuf.Reset();
}

func (nc *NewContent) writeSpaceBuf() {
  nc.SpaceBuf.WriteTo(nc.Value)
}

func (nc *NewContent) dumpWordBuf()  {
  nc.writeWordBuf()
  nc.resetWordBuf()
}

func (nc *NewContent) dumpSpaceBuf()  {
  nc.writeSpaceBuf()
  nc.resetSpaceBuf()
}

func (nc *NewContent) incSpaceBuf()  {
  nc.SpaceBuf.WriteRune(' ')
}

func (nc *NewContent) incWordBuf(char rune)  {
  nc.WordBuf.WriteRune(char)
}

func (nc *NewContent) addNewline()  {
  nc.Value.WriteRune('\n')
}

func Wrap(content string, lim int) string {
	Debug.Println("content: ", content)
	newContent := &NewContent{
		Lim:      lim,
		WordBuf:  bytes.NewBuffer([]byte{}),
		SpaceBuf: bytes.NewBuffer([]byte{}),
		Value:    bytes.NewBuffer([]byte{}),
	}

	for i, char := range content {
		if char != ' ' && char != '\n' {
			newContent.CurrentLineIndex++
			handleNonWhitespace(i, char, content, newContent)
			continue
		}

		if char == ' ' {
			handleSpace(i, char, newContent)
			continue
		}

		if char == '\n' {
			handleNewline(i, char, content, newContent)
			continue
		}
	}

	return string(newContent.Value.String())
}

func handleNonWhitespace(i int, char rune, content string, nc *NewContent) {
	wordBuf := nc.WordBuf
	value := nc.Value
	lim := nc.Lim

	wordBuf.WriteRune(char)

	if i == len(content)-1 {
		if nc.CurrentLineIndex > lim {
			value.WriteRune('\n')
      nc.dumpWordBuf()
			return
		}
    nc.dumpSpaceBuf()
    nc.dumpWordBuf()
		return
	}
}

func (nc *NewContent) isPastOrAtLim() bool {
	currentLineIndex := nc.CurrentLineIndex
	lim := nc.Lim

	return currentLineIndex >= lim
}

func (nc *NewContent) isPastLim() bool {
  currentLineIndex := nc.CurrentLineIndex
  lim := nc.Lim

  return currentLineIndex > lim
}


func handleSpace(i int, char rune, nc *NewContent) {
	wordBuf := nc.WordBuf

	nc.CurrentLineIndex++
	Debug.Println("in handleSpace:")
	Debug.Println("newContent", nc)
  t := true
	if nc.isPastOrAtLim() {
		nc.addNewline()
    nc.resetSpaceBuf()
		nc.CurrentLineIndex = wordBuf.Len()
	} else {
    t = false
    nc.writeSpaceBuf()
	}
  nc.dumpWordBuf()
  nc.resetSpaceBuf()
  if (t == false) {
    nc.incSpaceBuf()
  }
}

func handleNewline(i int, char rune, content string, nc *NewContent) {
	spaceBuf := nc.SpaceBuf
	value := nc.Value
	lim := nc.Lim
	currentLineIndex := nc.CurrentLineIndex

	// If we have a double newline that represents a paragraph
	// so I want to immeadialy append the 2 newlines
	// and get ready for the new paragraph's indent level
	if i > 0 && string(content[i-1]) == "\n" {
		value.WriteRune('\n')
		value.WriteRune('\n')
		spaceBuf.Reset()
		nc.CurrentLineIndex = 0
		return
	}

	// we need to wrap and current char is a newline
	if currentLineIndex > lim {
		value.WriteRune(char)
		nc.CurrentLineIndex = 0
    nc.dumpWordBuf()
		if i == len(content)-1 {
			value.WriteRune('\n')
		}

		return
	}

	// the entire input ends in newline
	if i == len(content)-1 {
    nc.dumpSpaceBuf()
    nc.dumpWordBuf()
		value.WriteRune(char)
		return
	}

	// the current line begins with a newline
	if currentLineIndex == 0 {
		value.WriteRune(char)
		spaceBuf.Reset()

		return
	}

	// if we are before the limit treat the newline as a space and write
	if currentLineIndex < lim {
    nc.dumpSpaceBuf()
    nc.dumpWordBuf()
		spaceBuf.WriteRune(' ')
		nc.CurrentLineIndex++
		return
	}
}

