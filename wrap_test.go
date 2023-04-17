package util

import (
	"testing"
)

func TestWrapGetsRidOfPrematureNewlines1(t *testing.T) {
	// t.Skip("trying to switch to my wrap func")
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
	// t.Skip("trying to switch to my wrap func")
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
	// t.Skip("trying to switch to my wrap func")
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
	// t.Skip("trying to switch to my wrap func")
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
	// t.Skip("skip")
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
	// t.Skip("trying to switch to my wrap func")
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
	// t.Skip("skip")
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
	// t.Skip("trying to switch to my wrap func")
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
	// t.Skip("skip")
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
	// t.Skip("skip")
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
		t.Errorf("Wrap(%q) = %q; expected %q", content, result, expected)
	}
}

func TestWrapPreservesIndents3(t *testing.T) {
	// t.Skip("skip")
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
		t.Errorf("Wrap(%q) = %q; expected %q", content, result, expected)
	}
}
