package d14

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	data := parse(strings.Split(example, "\n"))
	t.Run("p1", func(t *testing.T) {
		assert.Equal(t, -1, p1(*data))
	})
}

type cell int
const (
	empty cell = iota
	wall
	sand
)

type pair struct {
	x X
	y Y
}

type X int
type Y int

type grid struct {
	data map[X]map[Y]cell

	minX *X
	maxX *X
	minY *Y
	maxY *Y
}

func sort[T ~int](v1,v2 T) (T,T) {
	if v1 < v2 {
		return v1,v2
	}
	return v2,v1
}

func (g *grid) addRocks(a, b pair) {
	addRock := func(x X, y Y) {
		v, ok := g.data[x]
		if !ok {
			v = map[Y]cell{}
		}
		v[y] = wall
		g.data[x] = v
	}

	startX, endX := sort(a.x, b.x)
	startY, endY := sort(a.y, b.y)
	
	for x := startX; x <= endX; x++ {
		x := x
		for y := startY; y <= endY; y++ {
			y := y
			addRock(x,y)
			if g.maxX == nil || x > *g.maxX {
				g.maxX = &x
			} 
			if g.minX == nil || x < *g.minX {
				g.minX = &x
			}
			
			if g.maxY == nil || y > *g.maxY {
				g.maxY = &y
			}
		}
	}
}

func parse(lines []string) *grid {
	g := &grid{ data: map[X]map[Y]cell{} }
	var minY Y = 0
	g.minY = &minY

	for _, line := range lines {
		coords := strings.Split(line, " -> ")
		pairs := []pair{}
		for _, p := range coords {
			xy := strings.Split(p,",")
			x, _ :=strconv.Atoi(xy[0])
			y, _ :=strconv.Atoi(xy[1])

			pairs = append(pairs, pair{x: X(x), y: Y(y)})
		}

		for i := 0; i < len(pairs)-1; i++ {
			this := pairs[i]
			next := pairs[i+1]
			g.addRocks(this, next)
		}
	}
	return g
}

func p1(g grid) int {
	const startX X = 500
	const startY Y = 0

	return 0
}

const example = `498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`