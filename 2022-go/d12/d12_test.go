package d12

import (
	"fmt"
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
	lines := strings.Split(example,"\n")
	data := parse(lines)
	t.Run("p1", func(t *testing.T) {
		assert.Equal(t, 31, p1(data, findCoord(lines, 'S')[0],findCoord(lines, 'E')[0]))
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

func findCoord(lines []string, v byte) []coord {
	out := []coord{}
	for y, line := range lines {
		for x,c := range line {
			if byte(c) == v {
				out = append(out, coord{x:x, y:y})
			}
		}
	}
	return out
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
				{y,x-1},		  {y,x+1},
				{y+1,x-1},{y+1,x},{y+1,x+1},
			}
			for _, nei := range candidates {
				if !(nei.x >= 0 && nei.x < len(line) && nei.y >= 0 && nei.y < len(lines)) {
					continue
				}
				candidate := lines[nei.y][nei.x]
				fmt.Println(string(current), string(candidate))		
				if !areConnected(byte(current), candidate) {
					continue
				}
				g.connect(
					coord{ 
						x:int(convert(byte(x))),
						y:int(convert(byte(y))),
					},
					coord{ 
						x:int(convert(byte(nei.x))),
						y:int(convert(byte(nei.y))),
					})
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
	return s <= d+1
}

func p1(g graph, src, dst coord) int {
	dstE := dst.toEntry()

	pathTo := map[entry]entry{}
	visited := set{}
	q := queue{}

	var dfs func(entry)
	dfs = func(e entry) {
		if _, ok := visited[e]; ok {
			return
		}
		visited[e] = void{}
		for child := range g[e] {
			q = append(q, child)
		}
		for len(q) > 0 {
			next, _ := q.dequeue()
			dfs(next)
		}
	}

	dfs(src.toEntry())
	return 0
}

type queue []entry
func (q queue) enqueue(e entry) {
	q = append(q, e)
}

func (q queue) dequeue() (entry, bool) {
	if len(q) == 0 {
		var e entry
		return e, false
	}
	out := q[len(q)-1]
	q = q[:len(q)-1]
	return out, true
}