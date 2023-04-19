package util

import (
	"bytes"
	"strings"
	"unicode"
)

const nbsp = 0xA0

type NewContent struct {
	WordBuf          *bytes.Buffer
	SpaceBuf         *bytes.Buffer
	Value            *bytes.Buffer
	CurrentLineIndex int
	Lim              int
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

func handleNonWhitespace(i int, char rune, content string, newContent *NewContent) {
	wordBuf := newContent.WordBuf
	spaceBuf := newContent.SpaceBuf
	value := newContent.Value
	lim := newContent.Lim

	wordBuf.WriteRune(char)

	if i == len(content)-1 {
		if newContent.CurrentLineIndex > lim {
			value.WriteRune('\n')
			wordBuf.WriteTo(value)
			return
		}
		spaceBuf.WriteTo(value)
		wordBuf.WriteTo(value)
		return
	}
}
func isPastLim(newContent *NewContent) bool {
	currentLineIndex := newContent.CurrentLineIndex
	lim := newContent.Lim

  return currentLineIndex >= lim;
}

func handleSpace(i int, char rune, newContent *NewContent) {
	wordBuf := newContent.WordBuf
	spaceBuf := newContent.SpaceBuf
	value := newContent.Value
	lim := newContent.Lim
	currentLineIndex := newContent.CurrentLineIndex

	newContent.CurrentLineIndex++

	if currentLineIndex >= lim {
		value.WriteRune('\n')
		newContent.CurrentLineIndex = wordBuf.Len()
	} else {
		spaceBuf.WriteTo(value)
	}
	wordBuf.WriteTo(value)
	wordBuf.Reset()
	spaceBuf.Reset()
	spaceBuf.WriteRune(char)
}

func handleNewline(i int, char rune, content string, newContent *NewContent) {
	wordBuf := newContent.WordBuf
	spaceBuf := newContent.SpaceBuf
	value := newContent.Value
	lim := newContent.Lim
	currentLineIndex := newContent.CurrentLineIndex

  // If we have a double newline that represents a paragraph
  // so I want to immeadialy append the 2 newlines
  // and get ready for the new paragraph's indent level
	if i > 0 && string(content[i-1]) == "\n" {
		value.WriteRune('\n')
		value.WriteRune('\n')
		spaceBuf.Reset()
		newContent.CurrentLineIndex = 0
		return
	}

	// we need to wrap and current char is a newline
	if currentLineIndex > lim {
		value.WriteRune(char)
		newContent.CurrentLineIndex = 0
		wordBuf.WriteTo(value)
		wordBuf.Reset()
		if i == len(content)-1 {
			value.WriteRune('\n')
		}

		return
	}

	// the entire input ends in newline
	if i == len(content)-1 {
		spaceBuf.WriteTo(value)
		wordBuf.WriteTo(value)
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
		spaceBuf.WriteTo(value)
		wordBuf.WriteTo(value)
		spaceBuf.Reset()
		wordBuf.Reset()
		spaceBuf.WriteRune(' ')
		newContent.CurrentLineIndex++
		return
	}
}

// WrapString wraps the given string within lim width in characters.
//
// Wrapping is currently naive and only happens at white-space. A future
// version of the library will implement smarter wrapping. This means that
// pathological cases can dramatically reach past the limit, such as a very
// long word.
func WrapString(s string, lim uint) string {
	// Initialize a buffer with a slightly larger size to account for breaks
	init := make([]byte, 0, len(s))
	buf := bytes.NewBuffer(init)

	var current uint
	var wordBuf, spaceBuf bytes.Buffer
	var wordBufLen, spaceBufLen uint

	for _, char := range s {
		if char == '\n' {
			// does the text begin with a newline?
			if wordBuf.Len() == 0 {
				if current+spaceBufLen > lim {
					current = 0
				} else {
					current += spaceBufLen
					spaceBuf.WriteTo(buf)
				}
				spaceBuf.Reset()
				spaceBufLen = 0
			} else {
				current += spaceBufLen + wordBufLen
				spaceBuf.WriteTo(buf)
				spaceBuf.Reset()
				spaceBufLen = 0
				wordBuf.WriteTo(buf)
				wordBuf.Reset()
				wordBufLen = 0
			}
			buf.WriteRune(char)
			current = 0
		} else if unicode.IsSpace(char) && char != nbsp {
			if spaceBuf.Len() == 0 || wordBuf.Len() > 0 {
				current += spaceBufLen + wordBufLen
				spaceBuf.WriteTo(buf)
				spaceBuf.Reset()
				spaceBufLen = 0
				wordBuf.WriteTo(buf)
				wordBuf.Reset()
				wordBufLen = 0
			}

			spaceBuf.WriteRune(char)
			spaceBufLen++
		} else {
			wordBuf.WriteRune(char)
			wordBufLen++

			if current+wordBufLen+spaceBufLen > lim && wordBufLen < lim {
				buf.WriteRune('\n')
				current = 0
				spaceBuf.Reset()
				spaceBufLen = 0
			}
		}
	}

	if wordBuf.Len() == 0 {
		if current+spaceBufLen <= lim {
			spaceBuf.WriteTo(buf)
		}
	} else {
		spaceBuf.WriteTo(buf)
		wordBuf.WriteTo(buf)
	}

	return strings.Trim(string(buf.String()), " ")
}
