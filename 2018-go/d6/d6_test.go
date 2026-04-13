package d6

import (
	"fmt"
	"iter"
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

func TestP2(t *testing.T) {
	t.Run("ex", func(t *testing.T) {
		d, err := parse(example)
		require.NoError(t, err)

		assert.Equal(t, 16, p2(d, 32))
	})

	t.Run("file", func(t *testing.T) {
		f, err := os.ReadFile("d6.txt")
		require.NoError(t, err)
		d, err := parse(string(f))
		require.NoError(t, err)

		got := p2(d, 10000)
		assert.Equal(t, -1, got)
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

type coordIdx int
type distance struct {
	len        int
	coordIdx   coordIdx
	hadConfict bool
}

func p1(c []coord) int {
	start, end := findMapSize(c)

	areas := map[coordIdx]int{}
	infinite := map[coordIdx]bool{}

	for x := start.x; x <= end.x; x++ {
		for y := start.y; y <= end.y; y++ {

			this := coord{x, y}
			e := findMinimapDistanceForPoint(this, c)

			if e != nil && !e.hadConfict {
				areas[e.coordIdx]++

				// if current point was on edge - we're definitely in infinite categories, we need to skip it
				onEdge := x <= start.x || x >= end.x || y <= start.y || y >= end.y
				if onEdge {
					infinite[e.coordIdx] = true
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

func xyiter(start, end coord) iter.Seq[coord] {
	return func(yield func(coord) bool) {
		for x := start.x; x <= end.x; x++ {
			for y := start.y; y <= end.y; y++ {
				if !yield(coord{x, y}) {
					return
				}
			}
		}
	}
}

func p2(c []coord, limit int) int {
	start, end := findMapSize(c)

	distances := map[coord]bool{}
	for this := range xyiter(start, end) {
		sumOfDistances := calculatedAllDistances(this, c)
		if sumOfDistances < limit {
			distances[this] = true
		}
	}

	// todo: group together

	return 0
}

func calculatedAllDistances(this coord, allCoords []coord) int {
	sum := 0
	for _, c := range allCoords {
		sum += mahnattanDistance(this, c)
	}
	return sum
}

func findMinimapDistanceForPoint(this coord, c []coord) *distance {
	var e *distance

	for idx, v := range c {
		i := coordIdx(idx)
		thisDistance := mahnattanDistance(this, v)

		if e == nil {
			e = &distance{thisDistance, i, false}
		} else if thisDistance < e.len {
			e = &distance{thisDistance, i, false}
		} else if e.len == thisDistance {
			e = &distance{thisDistance, i, true}
		}
	}

	return e
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
