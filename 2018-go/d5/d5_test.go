package d5

import (
	"os"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestP1(t *testing.T) {
	got, err := os.ReadFile("d5.txt")
	require.NoError(t, err)
	assert.Equal(t, 9704, p1(string(got)))
}

func resolve(in string) string {
	out := in
	prevIdx := 0
	i := 1
	for len(out) != 0 && i < len(out) {
		this := out[i]
		prev := out[prevIdx]
		thisDbg := string(this)
		prevDbg := string(prev)
		_ = thisDbg
		_ = prevDbg
		if arePolar(prev, this) {
			out = out[:prevIdx] + out[i+1:]
			prevIdx = max(prevIdx-1, 0)
			i = prevIdx + 1
		} else {
			i++
			prevIdx++
		}
	}
	return out
}

func arePolar(a, b byte) bool {
	ar := rune(a)
	br := rune(b)
	return (unicode.IsUpper(ar) && unicode.IsLower(br) && unicode.ToLower(ar) == br) ||
		(unicode.IsUpper(br) && unicode.IsLower(ar) && unicode.ToLower(br) == ar)
}

func p1(input string) int {
	return len(resolve(input))
}
