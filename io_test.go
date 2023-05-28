package util

import (
	"testing"
)

func TestGetTermsEmpty(t *testing.T) {
  content := []string{}
  expected := 0
	result := len(GetTerms(content))

	if result != expected {
		t.Errorf("len(GetTerms(%q)) = %q; expected %q", content, result, expected)
	}
}
