package d8

import (
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const example = `30373
25512
65332
33549
35390`
func TestExample(t *testing.T) {
	data := parse(strings.Split(example, "\n"))
	t.Run("p1", func(t *testing.T) {
		assert.Equal(t, 21, p1(data))
	})

	t.Run("p2", func(t *testing.T) {
		assert.Equal(t, 8, p2(data))
	})

	t.Run("p2 distance1", func(t *testing.T) {
		assert.Equal(t, 4, countDistance(data, 1, 2))
	})

	t.Run("p2 distance2", func(t *testing.T) {
		assert.Equal(t, 8, countDistance(data, 3, 2))
	})
}

func TestCountDistance(t *testing.T) {
	data := parse(strings.Split(example, "\n"))
	t.Run("1,2 up", func(t *testing.T) {
		v, _ := data.goUp(1,2)
		assert.Equal(t, 1, v)
	})

	t.Run("1,2 left", func(t *testing.T) {
		v, _ := data.goLeft(1,2)
		assert.Equal(t, 1, v)
	})

	t.Run("1,2 right", func(t *testing.T) {
		v, _ := data.goRight(1,2)
		assert.Equal(t, 2, v)
	})

	t.Run("1,2 down", func(t *testing.T) {
		v, _ := data.goDown(1,2)
		assert.Equal(t, 2, v)
	})

	t.Run("3,2 up", func(t *testing.T) {
		v, _ := data.goUp(3,2)
		assert.Equal(t, 2, v)
	})

	t.Run("3,2 left", func(t *testing.T) {
		v, _ := data.goLeft(3,2)
		assert.Equal(t, 2, v)
	})

	t.Run("3,2 right", func(t *testing.T) {
		v, _ := data.goRight(3,2)
		assert.Equal(t, 2, v)
	})

	t.Run("3,2 down", func(t *testing.T) {
		v, _ := data.goDown(3,2)
		assert.Equal(t, 1, v)
	})
}

func TestFile(t *testing.T) {
	raw,err := os.ReadFile("data.txt")
	require.NoError(t, err)

	data := parse(strings.Split(string(raw), "\r\n"))
	t.Run("p1", func(t *testing.T) {
		assert.Equal(t, 1820, p1(data))
	})

	t.Run("p2", func(t *testing.T) {
		assert.NotEqual(t, 195, p2(data))
		assert.Equal(t, 385112, p2(data))
	})
}

type grid [][]int

func (g grid) cols() int {
	return len(g[0])
}

func (g grid) rows() int {
	return len(g)
}

func (g grid) el(r,c int) int {
	return g[r][c]
}

func (g grid) goUp(r,c int) (int, bool) {
	tree := g.el(r,c)
	for row := r-1; row >= 0; row-- {
		if g.el(row,c) >= tree {
			return r-row, false
		}
	}
	return r, true
}

func (g grid) visibleFromTop(r,c int) bool {
	_, ok := g.goUp(r,c)
	return ok
}

func (g grid) goDown(r,c int) (int, bool) {
	tree := g.el(r,c)
	for row := r+1; row < g.rows(); row++ {
		if g.el(row,c) >= tree {
			return row - r, false
		}
	}
	return g.rows() -r -1, true
}

func (g grid) visibleFromBottom(r,c int) bool {
	_, ok := g.goDown(r,c)
	return ok
}

func (g grid) goRight(r,c int) (int, bool) {
	tree := g.el(r,c)
	for col := c+1; col < g.cols(); col++ {
		if g.el(r,col) >= tree {
			return col-c, false
		}
	}
	return g.cols()-c-1,true
}

func (g grid) visibleFromRight(r,c int) bool {
	_, ok := g.goRight(r,c)
	return ok
}

func (g grid) goLeft(r,c int) (int, bool) {
	tree := g.el(r,c)
	for col := c-1; col >= 0; col-- {
		if g.el(r,col) >= tree {
			return c-col, false
		}
	}
	return c, true
}

func (g grid) visibleFromLeft(r,c int) bool {
	_, ok := g.goLeft(r,c)
	return ok
}

func parse(lines []string) grid {
	out := [][]int{}
	for _, line := range lines {
		row := []int{}
		for _, c := range line {
			v, _ := strconv.Atoi(string(c))
			row = append(row, v)
		}
		out = append(out, row)
	}
	return out
}

func p1(g grid) int {
	out := 0
	for r := 1; r < g.rows()-1; r++ {
		for c := 1; c < g.cols()-1; c++ {
			if g.visibleFromLeft(r,c) || g.visibleFromRight(r,c) || g.visibleFromTop(r,c) || g.visibleFromBottom(r,c) {
				out++
			}
		}
	}
	return out + (g.cols()-2)*2 + (g.rows())*2 // result + grid
}

func p2(g grid) int {
	max := 0
	for r := 0; r < g.rows(); r++ {
		for c := 0; c < g.cols(); c++ {
			if v := countDistance(g, r,c); v > max{
				max = v
			}
		}
	}
	return max
}

func countDistance(g grid, row,col int) int {
	down, _ := g.goDown(row, col)
	up, _ := g.goUp(row, col)
	left, _ := g.goLeft(row, col)
	right, _ := g.goRight(row, col)
	return down * up * left * right
}