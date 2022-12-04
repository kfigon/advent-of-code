package d4

import (
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExample(t *testing.T) {
	input := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`
	data := parse(strings.Split(input, "\n"))
	t.Run("p1", func(t *testing.T) {
		assert.Equal(t, 2, p1(data))
	})

	t.Run("p2", func(t *testing.T) {
		assert.Equal(t, 4, p2(data))
	})
}

func TestActual(t *testing.T) {
	raw, err := os.ReadFile("data.txt")
	require.NoError(t, err)

	data := parse(strings.Split(string(raw), "\r\n"))
	t.Run("p1", func(t *testing.T) {
		assert.Equal(t, 562, p1(data))
	})

	t.Run("p2", func(t *testing.T) {
		assert.Equal(t, 924, p2(data))
	})
}

type ranges struct {
	min int
	max int
}

func parseRange(in string) ranges {
	v := strings.Split(in, "-")
	min, _ := strconv.Atoi(v[0])
	max, _ := strconv.Atoi(v[1])
	return ranges{min, max}
}

func (r ranges) contains(other ranges) bool {
	return r.min >= other.min && r.max <= other.max
}

func (r ranges) doesntOverlap(other ranges) bool {
	return r.max < other.min || r.min > other.max
}

type pair struct {
	a ranges
	b ranges
}

func (p pair) overlapped() bool {
	first := p.a
	second := p.b
	return first.contains(second) || second.contains(first)
}

func (p pair) overlapPartially() bool {
	first := p.a
	second := p.b
	return !first.doesntOverlap(second) && !second.doesntOverlap(first)
}

func parse(lines []string) []pair {
	out := []pair{}
	for _, line := range lines {
		parts := strings.Split(line,",")
		p := pair {
			a: parseRange(parts[0]),
			b: parseRange(parts[1]),
		}
		out = append(out, p)
	}
	return out
}

func p1(pairs []pair) int {
	return count(pairs, func(p pair) bool {
		return p.overlapped()
	})
}

func p2(pairs []pair) int {
	return count(pairs, func(p pair) bool { 
		return p.overlapPartially()
	})
}

func count(pairs []pair, fn func(pair) bool) int {
	out := 0
	for _, p := range pairs {
		if fn(p) {
			out += 1
		}
	}
	return out
}