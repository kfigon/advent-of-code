package d12

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const example = `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`

func TestExample(t *testing.T) {
	data := parse(strings.Split(example,"\n"))
	t.Run("p1", func(t *testing.T) {
		assert.Equal(t, 31, p1(data))
	})
}

type entry string
type void struct{}
type set map[entry]void
type graph map[entry]set
type coord struct {
	y int
	x int
}

func (g graph) connect(a,b coord) {
	aC := a.toEntry()
	bC := b.toEntry()

	v, ok := g[aC]
	if !ok {
		v = set{}
	}
	v[bC] = void{}
	g[aC] = v
}

func (c coord) toEntry() entry {
	return entry(strconv.Itoa(c.x)+","+strconv.Itoa(c.y))
}

func (e entry) toCoord() coord {
	pair := strings.Split(string(e), ",")
	x,_ := strconv.Atoi(pair[0])
	y,_ := strconv.Atoi(pair[1])
	return coord{x:x,y:y}
}

func parse(lines []string) graph {
	g := graph{}
	for y, line := range lines {
		for x, current := range line {
			candidates := []coord{
				{y-1,x-1},{y-1,x},{y-1,x+1},
				{y,x-1},{y,x},{y,x+1},
				{y+1,x-1},{y+1,x},{y+1,x+1},
			}
			for _, nei := range candidates {
				if !(nei.x >= 0 && nei.x < len(line) && nei.y >= 0 && nei.y < len(lines)) {
					continue
				}
				candidate := lines[nei.y][nei.x]
				if !areConnected(byte(current), candidate) {
					continue
				}
				g.connect(coord{x:x, y:y},nei)
			}
		}
	}
	return g
}

func areConnected(src, dst byte) bool {
	convert := func(v byte) byte {
		if v == 'S' {
			return 'a'
		} else if v == 'E' {
			return 'z'
		}
		return v
	}
	return convert(src) + 1 <= convert(dst)
}

func p1(g graph) int {
	return 0
}