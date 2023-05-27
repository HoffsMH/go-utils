package util

import (
	"os"
	"path"
	"strings"
	"testing"

	"github.com/jmhodges/clock"
)

func TestPrefixNameISORelativeNameNoDir(t *testing.T) {
	content := `foo.md`
  wd, _ := os.Getwd()
	expected := wd + "/1970-01-01T00:00:00Z-" + "foo.md"

  fakeClock := clock.NewFake();
  opts := PrefixOptions{
    Clock: fakeClock,
  }

	result := PrefixNameISO(content, opts)
	if result != expected {
		t.Errorf("PrefixNameISO(%q) = %q; expected %q", content, result, expected)
	}
}

func TestPrefixNameISORelativeName(t *testing.T) {
	content := `../foo.md`
  wd, _ := os.Getwd()
  dirs := strings.Split(wd, "/")
  dirs = dirs[:len(dirs)-1]
  wd = strings.Join(dirs, "/")
	expected := wd + "/1970-01-01T00:00:00Z-" + "foo.md"

  fakeClock := clock.NewFake();
  opts := PrefixOptions{
    Clock: fakeClock,
  }

	result := PrefixNameISO(content, opts)
	if result != expected {
		t.Errorf("PrefixNameISO(%q) = %q; expected %q", content, result, expected)
	}
}

func TestPrefixNameISOAbs(t *testing.T) {
  wd, _ := os.Getwd()
	content := path.Join(wd, "foo.md")
  dirs := strings.Split(wd, "/")
  wd = strings.Join(dirs, "/")
	expected := wd + "/1970-01-01T00:00:00Z-" + "foo.md"

  fakeClock := clock.NewFake();
  opts := PrefixOptions{
    Clock: fakeClock,
  }

	result := PrefixNameISO(content, opts)
	if result != expected {
		t.Errorf("PrefixNameISO(%q) = %q; expected %q", content, result, expected)
	}
}

func TestPrefixNameDateRelativeNameNoDir(t *testing.T) {
	content := `foo.md`
  wd, _ := os.Getwd()
	expected := wd + "/1970-01-01-" + "foo.md"

  fakeClock := clock.NewFake();
  opts := PrefixOptions{
    Clock: fakeClock,
  }

	result := PrefixNameDate(content, opts)
	if result != expected {
		t.Errorf("PrefixNameISO(%q) = %q; expected %q", content, result, expected)
	}
}

func TestPrefixNameDateRelativeName(t *testing.T) {
	content := `../foo.md`
  wd, _ := os.Getwd()
  dirs := strings.Split(wd, "/")
  dirs = dirs[:len(dirs)-1]
  wd = strings.Join(dirs, "/")
	expected := wd + "/1970-01-01-" + "foo.md"
  fakeClock := clock.NewFake();
  opts := PrefixOptions{
    Clock: fakeClock,
  }

	result := PrefixNameDate(content, opts)
	if result != expected {
		t.Errorf("PrefixNameISO(%q) = %q; expected %q", content, result, expected)
	}
}


