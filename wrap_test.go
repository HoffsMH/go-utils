package util

import (
	"testing"
)

// func TestWrap(t *testing.T) {
//   // collapses double newlines to one
//   // preserves indents
//   content := `asdf asdf asdf asdf asdf asdf asdf sadf sadf asdf asdf asdf asdf asdf asdf asdfas dfasf
//   asdfsadf asdf asdf asdf asdf asdf asdf sadf asdf asdf asdf asdf
//
//   asdfasd asdf asdf asdf asdf asdf asd fsdf
//
//       asdfasdfasdfas sadff sdf asdf asdf asdf asdf asdfas dfas dfsd asdf asdf sadf asdf asdf
//   asadfasdfasdf asdf asdf asdf asdf sadf sadf asdf sad fsadf asfd
// `
// 	expected := `asdf`
//
// 	result := Wrap(content, 50)
// 	if result != expected {
// 		t.Errorf("Wrap(%v) = %v; expected %v", content, result, expected)
// 	}
// }

func TestWrapok(t *testing.T) {
  // collapses double newlines to one
  // preserves indents
  content := `asdf asdf asdf asdf asdf asdf asdf sadf sadf asdf asdf asdf asdf asdf asdf asdfas dfasf
  asdfsadf asdf asdf asdf asdf asdf asdf sadf asdf asdf asdf asdf

  asdfasd asdf asdf asdf asdf asdf asd fsdf 

      asdfasdfasdfas sadff sdf asdf asdf asdf asdf asdfas dfas dfsd asdf asdf sadf asdf asdf 
  asadfasdfasdf asdf asdf asdf asdf sadf sadf asdf sad fsadf asfd
`
	expected := `asdf`

	result := WrapString(content, 50)
	if result != expected {
		t.Errorf("Wrap(%v) = %v; expected %v", content, result, expected)
	}
}
