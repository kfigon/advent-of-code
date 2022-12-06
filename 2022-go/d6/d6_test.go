package d6

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExampleP1(t *testing.T) {
	testCases := []struct {
		in	string
		exp int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
	}
	for _, tC := range testCases {
		t.Run(tC.in, func(t *testing.T) {
			assert.Equal(t, tC.exp, p1(tC.in))
		})
	}
}

func TestExampleP2(t *testing.T) {
	testCases := []struct {
		in	string
		exp int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 23},
		{"nppdvjthqldpwncqszvftbrmjlhg", 23},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26},
	}
	for _, tC := range testCases {
		t.Run(tC.in, func(t *testing.T) {
			assert.Equal(t, tC.exp, p2(tC.in))
		})
	}
}

func TestActual(t *testing.T) {
	raw, err := os.ReadFile("data.txt")
	require.NoError(t, err)

	t.Run("p1", func(t *testing.T) {
		assert.Equal(t, 1235, p1(string(raw)))
	})

	t.Run("p2", func(t *testing.T) {
		assert.Equal(t, 3051, p2(string(raw)))
	})
}

func findUniqueSubsequent(input string, width int) int {
	for i := width; i < len(input); i++ {
		set := map[rune]bool{}
		for _, c := range input[i-width:i] {
			set[c] = true
		}

		if len(set) == width {
			return i
		}
	}

	return -1
}

func p1(input string) int{
	return findUniqueSubsequent(input, 4)
}

func p2(input string) int{
	return findUniqueSubsequent(input, 14)
}