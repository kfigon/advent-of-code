package d3

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_example(t *testing.T) {
	raw := `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`
	bags := strings.Split(raw, "\n")

	t.Run("p1", func(t *testing.T) {
		assert.Equal(t, 157, p1(bags))
	})
	t.Run("p2", func(t *testing.T) {
		assert.Equal(t, 70, p2(bags))
	})
}

func Test_real(t *testing.T) {
	raw, err := os.ReadFile("data.txt")
	require.NoError(t, err)
	bags := strings.Split(string(raw), "\r\n")

	t.Run("p1", func(t *testing.T) {
		assert.Equal(t, 8493, p1(bags))
	})
	t.Run("p2", func(t *testing.T) {
		assert.Equal(t, 2552, p2(bags))
	})
}

func TestPrio(t *testing.T) {
	assert.Equal(t, 1, priority('a'))
	assert.Equal(t, 26, priority('z'))
	assert.Equal(t, 27, priority('A'))
	assert.Equal(t, 52, priority('Z'))
}

func priority(char rune) int {
	if char >= 'a' && char <= 'z' {
		return int(char) - int('a') + 1
	} else if char >= 'A' && char <= 'Z' {
		return int(char) - int('A') + 27
	}
	return 0
}

func buildSet(s string) map[rune]bool{
	occurences := map[rune]bool{}
	for _, v := range s {
		occurences[v] = true
	}
	return occurences
}

func p1(bags []string) int {
	repeating := func(b string) rune {
		first := b[:len(b)/2]
		second := b[len(b)/2:]

		occurences := buildSet(first)
		for _, v := range second {
			if ok := occurences[v]; ok {
				return v
			}
		}
		return 0
	}

	out := 0
	for _, b := range bags {
		out += priority(repeating(b))
	}
	return out
}

func p2(bags []string) int {
	const limit = 3

	findCommonItem := func(bags []string) rune{
		occurencesInAllBags := map[rune]int{}

		for _, bag := range bags {
			bagSet := buildSet(bag)
			for c := range bagSet {
				v := occurencesInAllBags[c]
				occurencesInAllBags[c] = v+1
			}
		}

		for item,count := range occurencesInAllBags {
			if count == limit {
				return item
			}
		}
		return 0
	}

	out := 0
	for i := 0; i < len(bags); i+=limit {
		out += priority(findCommonItem(bags[i:i+limit]))
	}
	return out
}