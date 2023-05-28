package util

import (
	"os"
	"path"
	"strings"
	"testing"
	"time"

	"github.com/jmhodges/clock"
)

func TestPrefixNameISORelativeNameNoDir(t *testing.T) {
	content := `foo.md`
  wd, _ := os.Getwd()
	expected := wd + "/1970-01-01T00:00:00Z-" + "foo.md"

  fakeClock := clock.NewFake();
  prefixer := &Prefixer{
    Clock: fakeClock,
    Format: time.RFC3339,
  }

	result := prefixer.Name(content)
	if result != expected {
		t.Errorf("Name(%q) = %q; expected %q", content, result, expected)
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
  prefixer := &Prefixer{
    Clock: fakeClock,
    Format: time.RFC3339,
  }

	result := prefixer.Name(content)
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
  prefixer := &Prefixer{
    Clock: fakeClock,
    Format: time.RFC3339,
  }

	result := prefixer.Name(content)
	if result != expected {
		t.Errorf("Name(%q) = %q; expected %q", content, result, expected)
	}
}

func TestPrefixNameDateRelativeNameNoDir(t *testing.T) {
	content := `foo.md`
  wd, _ := os.Getwd()
	expected := wd + "/1970-01-01-" + "foo.md"

  fakeClock := clock.NewFake();
  prefixer := &Prefixer{
    Clock: fakeClock,
    Format: "2006-01-02",
  }

	result := prefixer.Name(content)
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
  prefixer := Prefixer{
    Clock: fakeClock,
    Format: "2006-01-02",
  }

	result := prefixer.Name(content)
	if result != expected {
		t.Errorf("PrefixNameISO(%q) = %q; expected %q", content, result, expected)
	}
}


