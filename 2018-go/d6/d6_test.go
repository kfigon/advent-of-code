package d6

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var example = `1, 1
1, 6
8, 3
3, 4
5, 5
8, 9`

func TestP1(t *testing.T) {
	t.Run("ex", func(t *testing.T) {
		d, err := parse(example)
		require.NoError(t, err)

		assert.Equal(t, 17, p1(d))
	})

	t.Run("file", func(t *testing.T) {
		f, err := os.ReadFile("d6.txt")
		require.NoError(t, err)
		d, err := parse(string(f))
		require.NoError(t, err)

		got := p1(d)
		assert.Equal(t, 4284, got)
	})
}

type coord struct {
	x int
	y int
}

func parse(data string) ([]coord, error) {
	out := []coord{}
	for line := range strings.Lines(data) {
		line = strings.TrimSpace(line)
		got := strings.Split(line, ", ")
		if len(got) != 2 {
			return nil, fmt.Errorf("separator not found for line: %q", line)
		}
		a, err := strconv.Atoi(got[0])
		if err != nil {
			return nil, fmt.Errorf("parsing error for %q: %w", line, err)
		}
		b, err := strconv.Atoi(got[1])
		if err != nil {
			return nil, fmt.Errorf("parsing error for %q: %w", line, err)
		}
		out = append(out, coord{a, b})
	}
	return out, nil
}

func p1(c []coord) int {
	start, end := findMapSize(c)
	type distance struct {
		len        int
		coordIdx   int
		hadConfict bool
	}

	areas := map[coord]int{}
	infinite := map[coord]bool{}

	for x := start.x; x <= end.x; x++ {
		for y := start.y; y <= end.y; y++ {

			var e *distance
			this := coord{x, y}

			for i, v := range c {

				thisDistance := mahnattanDistance(this, v)

				if e == nil {
					e = &distance{thisDistance, i, false}
				} else if thisDistance < e.len {
					e = &distance{thisDistance, i, false}
				} else if e.len == thisDistance {
					e = &distance{thisDistance, i, true}
				}
			}

			if e != nil && !e.hadConfict {
				cell := c[e.coordIdx]
				areas[cell]++

				onEdge := x <= start.x || x >= end.x || y <= start.y || y >= end.y
				if onEdge {
					infinite[cell] = true
				}
			}
		}
	}

	maxLen := 0
	for candidate, v := range areas {
		if infinite[candidate] {
			continue
		}
		maxLen = max(maxLen, v)
	}

	return maxLen
}

func findMapSize(c []coord) (start coord, end coord) {
	var zero coord
	if len(c) < 1 {
		return zero, zero
	}

	var minX *int
	var minY *int
	var maxX *int
	var maxY *int
	for _, v := range c {
		if minX == nil || v.x < *minX {
			minX = &v.x
		}
		if minY == nil || v.y < *minY {
			minY = &v.y
		}
		if maxX == nil || v.x > *maxX {
			maxX = &v.x
		}
		if maxY == nil || v.y > *maxX {
			maxY = &v.y
		}
	}
	return coord{x: *minX, y: *minY}, coord{x: *maxX, y: *maxY}
}

func mahnattanDistance(a, b coord) int {
	return int(math.Abs(float64(a.x-b.x)) + math.Abs(float64(a.y-b.y)))
}
