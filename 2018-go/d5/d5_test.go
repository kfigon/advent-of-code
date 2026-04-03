package d5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestP1Examples(t *testing.T) {
	testCases := []struct {
		input string
		exp   string
	}{
		{"aA", ""},
		{"abBA", ""},
		{"abAB", "abAB"},
		{"aabAAB", "aabAAB"},
		{"dabAcCaCBAcCcaDA", "dabCBAcaDA"},
	}
	for _, tC := range testCases {
		t.Run(tC.input, func(t *testing.T) {
			assert.Equal(t, tC.exp, resolve(tC.input))
		})
	}
}

func resolve(in string) string {
	return ""
}
