package util

import (
	"fmt"
	"testing"
)

func TestWrapGetsRidOfPrematureNewlines1(t *testing.T) {
	content := `
asdf`
	expected := `
asdf`

	result := Wrap(content, 20)
	if result != expected {
		t.Errorf("Wrap(%q) = %q; expected %q", content, result, expected)
	}
}

func TestWrapGetsRidOfPrematureNewlines2(t *testing.T) {
	content := `
asdf
`
	expected := `
asdf
`

	result := Wrap(content, 20)
	if result != expected {
		t.Errorf("Wrap(%q) = %q; expected %q", content, result, expected)
	}
}

func TestWrapGetsRidOfPrematureNewlines3(t *testing.T) {
	content := `
asdf
asdf
`
	expected := `
asdf asdf
`

	result := Wrap(content, 20)
	if result != expected {
		t.Errorf("Wrap(%q) = %q; expected %q", content, result, expected)
	}
}

func TestWrap2(t *testing.T) {
	content := `
asdf asdf
asdf
`
	expected := `
asdf asdf asdf
`

	result := Wrap(content, 20)
	if result != expected {
		t.Errorf("Wrap(%q) = %q; expected %q", content, result, expected)
	}
}

func TestWrap3(t *testing.T) {
	content := `
asdf asdf asdf asdf
asdf
`
	expected := `
asdf asdf asdf asdf
asdf
`

	result := Wrap(content, 20)
	if result != expected {
		t.Errorf("Wrap(%q) = %q; expected %q", content, result, expected)
	}
}

func TestWrap4(t *testing.T) {
	content := `
asdf asdf asdf asdf
asdf`
	expected := `
asdf asdf asdf asdf
asdf`

	result := Wrap(content, 20)
	if result != expected {
		t.Errorf("Wrap(%q) = %q; expected %q", content, result, expected)
	}
}

func TestWrapPreservesParagraphs1(t *testing.T) {
	content := `
asdf asdf asdf asdf
asdf

asdf`
	expected := `
asdf asdf asdf asdf
asdf

asdf`

	result := Wrap(content, 20)
	if result != expected {
		t.Errorf("Wrap(%q) = %q; expected %q", content, result, expected)
	}
}

func TestWrapGetsRidOfPrematureNewlines(t *testing.T) {
	content := `
asdf
asdf
asdf
asdf
asdf
asdf
asd
fsdf
`
	expected := `
asdf asdf asdf asdf
asdf asdf asd fsdf
`

	result := Wrap(content, 20)
	if result != expected {
		t.Errorf("Wrap(%q) = %q; expected %q", content, result, expected)
	}
}

func TestWrapPreservesIndents1(t *testing.T) {
	content := `
asdf asdf
asdf

    asdf
`
	expected := `
asdf asdf asdf

    asdf
`

	result := Wrap(content, 20)
	if result != expected {
		t.Errorf("Wrap(%q) = %q; expected %q", content, result, expected)
	}
}

func TestWrapPreservesIndents2(t *testing.T) {
	content := `
asdf asdf
asdf

    asdf
asdf
asdf
asdf
asdf
asdf
`
	expected := `
asdf asdf asdf

    asdf asdf asdf
asdf asdf asdf
`

	result := Wrap(content, 20)
	if result != expected {
		fmt.Print("expected:", expected)
		fmt.Print("result:", result)
		t.Errorf("Wrap(%q) = %q; expected %q", content, result, expected)
	}
}

func TestWrapPreservesIndents3(t *testing.T) {
	content := `
asdf asdf asdf asdf asdf asdf asdf asdf asdf

    asdf
asdf
asdf
asdf
asdf
asdf

`
	expected := `
asdf asdf asdf asdf
asdf asdf asdf asdf
asdf

    asdf asdf asdf
asdf asdf asdf

`

	result := Wrap(content, 20)
	if result != expected {
    fmt.Printf("result %s", result)
    fmt.Printf("expected %s", expected)
		t.Errorf("Wrap(%q) = %q; expected %q", content, result, expected)
	}
}

func TestWrapPreservesSpaces(t *testing.T) {
	content := `asdf asdf asdf asdf
asdfghjkl`

	expected := `asdf asdf asdf
asdf asdfghjkl`

	result := Wrap(content, 18)
	if result != expected {
		t.Errorf("Wrap(%q) = %q; expected %q", content, result, expected)
	}
}

func TestWrapStripsTrailingSpaces(t *testing.T) {
  // isDev := true
  // Init(&isDev)
	content := `asdf asdf asdf asdf 
asdfghjkl`

	expected := `asdf asdf asdf
asdf asdfghjkl`

	result := Wrap(content, 18)
	if result != expected {
		t.Errorf("Wrap(%q) = %q; expected %q", content, result, expected)
	}
}

func TestWrapStripsTrailingSpaces2(t *testing.T) {
  // t.Skip()
  // isDev := true
  // Init(&isDev)

	content := `asdf asdf asdf asdf iasdf 
asdfghjkl`

	expected := `asdf asdf asdf
asdf iasdf
asdfghjkl`

	result := Wrap(content, 18)
	if result != expected {
		t.Errorf("Wrap(%q) = %q; expected %q", content, result, expected)
	}
}

func TestWrapShouldNotStripNeededSpaces(t *testing.T) {
  // t.Skip()
  // isDev := true
  // Init(&isDev)

	content := `asdf assdf asdf asdf asdf asdf asdf asdf asdf
asjsdf asdf asdf asdf asdf asdf asdf asdf asdf
asldf asdf asdf asdf asdf asdf asdf asdf asdf
asdfghjkl`

	expected := `asdf asdf asdf
asdf iasdf
asdfghjkl`

	result := Wrap(content, 18)
  fmt.Println("============================")
  fmt.Printf("result %s", result)
  fmt.Println("============================")
  fmt.Printf("expected %s", expected)
	if result != expected {
		t.Errorf("Wrap(%q) = %q; expected %q", content, result, expected)
	}
}
