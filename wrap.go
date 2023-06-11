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

func (nc *NewContent) isAtLim() bool {
  currentLineIndex := nc.CurrentLineIndex
  lim := nc.Lim

  return currentLineIndex == lim
}

func Wrap(content string, lim int) string {
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
  value := nc.Value

  nc.incWordBuf(char)
  Debug.Println("======================")
  Debug.Println("in non whitespace")
  Debug.Println("i: ", i)
  Debug.Println("char: ", string(char))
  Debug.Println("content: ", content)
  Debug.Println("nc: ", nc)
  Debug.Println("nc.SpaceBuf.Len(): ", nc.SpaceBuf.Len())
  Debug.Println("nc.WordBuf.Len(): ", string(nc.WordBuf.Bytes()))

  if i == len(content)-1 {
    if nc.isPastLim() {
      value.WriteRune('\n')
      nc.dumpWordBuf()
      return
    }
    nc.dumpSpaceBuf()
    nc.dumpWordBuf()
    return
  }
}

func handleSpace(i int, char rune, nc *NewContent) {

  nc.CurrentLineIndex++

  Debug.Println("======================")
  Debug.Println("in Space")
  Debug.Println("i: ", i)
  Debug.Println("char: ", string(char))
  Debug.Println("nc: ", nc)
  Debug.Println("nc.SpaceBuf.Len(): ", nc.SpaceBuf.Len())
  Debug.Println("nc.WordBuf.Len(): ", string(nc.WordBuf.Bytes()))
  if nc.isAtLim() {
    Debug.Println("at lim")
    nc.dumpSpaceBuf()
    nc.dumpWordBuf()
    nc.addNewline()
    nc.CurrentLineIndex = 0
    return
  }

  if nc.isPastLim() {
    nc.addNewline()
    nc.CurrentLineIndex = nc.WordBuf.Len()
    nc.dumpWordBuf()
    nc.resetSpaceBuf()
    nc.incSpaceBuf()
    nc.CurrentLineIndex += nc.SpaceBuf.Len()
    return
  }

  nc.dumpSpaceBuf()
  nc.dumpWordBuf()
  nc.incSpaceBuf()
}

func handleNewline(i int, char rune, content string, nc *NewContent) {
  currentLineIndex := nc.CurrentLineIndex

  Debug.Println("======================")
  Debug.Println("in Handle Newline")
  Debug.Println("i: ", i)
  Debug.Println("char: ", char)
  Debug.Println("content: ", content)
  Debug.Println("nc: ", nc)
  Debug.Println("nc.SpaceBuf.Len(): ", nc.SpaceBuf.Len())
  Debug.Println("nc.WordBuf.Len(): ", string(nc.WordBuf.Bytes()))

  // If we have a double newline that represents a paragraph
  // so I want to immeadialy append the 2 newlines
  // and get ready for the new paragraph's indent level
  if i > 0 && string(content[i-1]) == "\n" {
    nc.addNewline()
    nc.addNewline()
    nc.resetSpaceBuf()
    nc.CurrentLineIndex = 0
    return
  }

  // we need to wrap and current char is a newline
  if nc.isPastLim() {
    nc.addNewline()
    nc.CurrentLineIndex = 0
    nc.dumpWordBuf()
    if i == len(content)-1 {
      nc.addNewline()
    }

    return
  }

  // the entire input ends in newline
  if i == len(content)-1 {
    nc.dumpSpaceBuf()
    nc.dumpWordBuf()
    nc.addNewline()
    return
  }

  // the current line begins with a newline
  if currentLineIndex == 0 {
    nc.addNewline()
    nc.resetSpaceBuf()

    return
  }

  // if we are before the limit treat the newline as a space and write
  if !nc.isPastOrAtLim() && nc.WordBuf.Len() > 0 {
    nc.dumpSpaceBuf()
    nc.dumpWordBuf()
    nc.incSpaceBuf()
		nc.CurrentLineIndex++
		return
	}
}

