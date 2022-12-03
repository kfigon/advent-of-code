package d1

import (
	"os"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestD1(t *testing.T) {
	file := readFile("data.txt")
	t.Run("p1", func(t *testing.T) {
		assert.Equal(t, 66719, solveP1(file))
	})

	t.Run("p2", func(t *testing.T) {
		assert.Equal(t, 198551, solveP2(file))
	})
}

func readFile(filename string) []int {
	d, err := os.ReadFile(filename)
	if err != nil {
		return nil
	}

	lines := strings.Split(string(d), "\r\n")
	out := []int{}
	currentElf := 0
	for _, line := range lines {
		if line == "" {
			out = append(out, currentElf)
			currentElf = 0
			continue
		}
		v, err := strconv.Atoi(line)
		if err != nil {
			return nil
		}
		currentElf += v
	}
	return out
}

func solveP1(data []int) int {
	max := data[0]
	for i := 1; i < len(data); i++ {
		d := data[i]
		if d > max {
			max = d
		}
	}
	return max
}

// priority queue, yeah I know...
func solveP2(data []int) int {
	sort.Ints(data)
	last := len(data)-1
	return data[last] + data[last-1] + data[last-2]
}