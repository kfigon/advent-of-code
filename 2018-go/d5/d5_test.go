package d5

import (
	"os"
	"strings"
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

func TestP2(t *testing.T) {
	t.Run("ex", func(t *testing.T) {
		assert.Equal(t, 4, p2("dabAcCaCBAcCcaDA"))
	})
	t.Run("task", func(t *testing.T) {
		got, err := os.ReadFile("d5.txt")
		require.NoError(t, err)
		assert.Equal(t, 6942, p2(string(got)))
	})
}

func resolve(in string) string {
	// this is a stack, if last interracts with current - pop
	// reassinging strings out = out[:prevIdx] + out[i+1:] is slow, because it allocates
	out := make([]byte, 0, len(in))

	for i := 0; i < len(in); i++ {
		this := in[i]

		if len(out) > 0 && arePolar(out[len(out)-1], this) {
			out = out[:len(out)-1] // pop
		} else {
			out = append(out, this) // push
		}
	}
	return string(out)
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

func p2(input string) int {
	allChars := map[rune]bool{}
	for _, v := range input {
		allChars[unicode.ToLower(v)] = true
	}
	minPol := 999999999999999
	for c := range allChars {
		newStr := strings.ReplaceAll(input, string(unicode.ToLower(c)), "")
		newStr = strings.ReplaceAll(newStr, string(unicode.ToUpper(c)), "")
		got := len(resolve(newStr))
		minPol = min(minPol, got)
	}
	return minPol
}
