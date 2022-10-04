package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTermsEmpty(t *testing.T) {
	noTerms := GetTerms([]string{})
	assert.Empty(t, noTerms)
}
