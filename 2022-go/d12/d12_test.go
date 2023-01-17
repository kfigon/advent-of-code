package d12

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const example = `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`

func TestExample(t *testing.T) {
	lines := strings.Split(example, "\n")
	data := parse(lines)
	t.Run("p1", func(t *testing.T) {
		assert.Equal(t, 31, p1(data, lines))
	})

	t.Run("p2", func(t *testing.T) {
		assert.Equal(t, 29, p2(data, lines))
	})
}

func TestFile(t *testing.T) {
	raw, err := os.ReadFile("data.txt")
	require.NoError(t, err)

	lines := strings.Split(string(raw), "\r\n")
	data := parse(lines)
	t.Run("p1", func(t *testing.T) {
		assert.Equal(t, 497, p1(data, lines))
	})

	t.Run("p2", func(t *testing.T) {
		assert.Equal(t, 492, p2(data, lines))
	})
}

func TestQueue(t *testing.T) {
	q := &queue{}
	q.enqueue(coord{1,2})
	q.enqueue(coord{2,3})
	q.enqueue(coord{3,4})

	assert.Equal(t, 3, len(*q))
	
	v1, _ := q.dequeue()
	assert.Equal(t, coord{1,2}, v1)
	
	v2, _ := q.dequeue()
	assert.Equal(t, coord{2,3}, v2)

	v3, _ := q.dequeue()
	assert.Equal(t, coord{3,4}, v3)
}

func TestConnected(t *testing.T) {
	assert.True(t, areConnected('a', 'b'))
	assert.True(t, areConnected('a', 'a'))
	
	assert.True(t, areConnected('d', 'e'))
	assert.True(t, areConnected('d', 'a'))

	assert.False(t, areConnected('d', 'f'))
	assert.False(t, areConnected('d', 'g'))
	assert.False(t, areConnected('d', 'z'))
}

type void struct{}
type set map[coord]void
type graph map[coord]set
type coord struct {
	x int
	y int
}

func (c coord) String() string {
	return fmt.Sprintf("%v,%v", c.x,c.y)
}

func findCoord(lines []string, v byte) []coord {
	out := []coord{}
	for y, line := range lines {
		for x, c := range line {
			if byte(c) == v {
				out = append(out, coord{x: x, y: y})
			}
		}
	}
	return out
}

func (g graph) connect(a, b coord) {
	v, ok := g[a]
	if !ok {
		v = set{}
	}
	v[b] = void{}
	g[a] = v
}

func parse(lines []string) graph {
	g := graph{}
	for y, line := range lines {
		for x, current := range line {
			candidates := []coord{
						{x: x,y: y - 1},
				{x: x - 1, y:y}, {x: x + 1, y: y},
						{x: x, y: y + 1},
			}
			for _, nei := range candidates {
				if !(nei.x >= 0 && nei.x < len(line) && nei.y >= 0 && nei.y < len(lines)) {
					continue
				}
				candidate := lines[nei.y][nei.x]
				if !areConnected(byte(current), candidate) {
					continue
				}
				g.connect(
					coord{x: x, y: y},
					coord{x: nei.x, y: nei.y},
				)
			}
		}
	}
	return g
}
func convert(v byte) byte {
	if v == 'S' {
		return 'a'
	} else if v == 'E' {
		return 'z'
	}
	return v
}

func areConnected(src, dst byte) bool {
	s := convert(src)
	d := convert(dst)
	return s+1 >= d
}

func solve(g graph, src, dst coord) int {
	distances := map[coord]int{} // also visited set
	q := &queue{}

	q.enqueue(src)
	distances[src] = 0

	for len(*q) > 0 {
		current, _ := q.dequeue()
		if current == dst {
			break
		}

		for child := range g[current] {
			if _, ok := distances[child]; !ok {
				distances[child] = distances[current]+1
				q.enqueue(child)
			}
		}
	}

	return distances[dst]
}

func p1(g graph, lines []string) int {
	src := findCoord(lines, 'S')[0]
	dst := findCoord(lines, 'E')[0]
	return solve(g, src, dst)
}

func p2(g graph, lines []string) int {
	// ideally search the shortest path from E to any 'a', but need to rework the
	// areConnected method and move it to bfs
	// but i'm lazy :D

	src := findCoord(lines, 'a')
	dst := findCoord(lines, 'E')[0]
	src = append(src, findCoord(lines, 'S')[0])

	min := 99999999999
	for _, s := range src {
		r := solve(g, s,dst)
		if r < min && r != 0 {
			min = r
		}
	}
	return min
}

type queue []coord

func (q *queue) enqueue(c coord) {
	*q = append(*q, c)
}

func (q *queue) dequeue() (coord, bool) {
	if len(*q) == 0 {
		return coord{}, false
	}
	out := (*q)[0]
	*q = (*q)[1:]
	return out, true
}
