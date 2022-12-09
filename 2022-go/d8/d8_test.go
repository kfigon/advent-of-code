package d8

import (
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExample(t *testing.T) {
	raw := `30373
25512
65332
33549
35390`
	data := parse(strings.Split(raw, "\n"))
	t.Run("p1", func(t *testing.T) {
		assert.Equal(t, 21, p1(data))
	})

	t.Run("p2", func(t *testing.T) {
		t.Fail()
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
		t.Fail()
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

func (g grid) visibleFromTop(r,c int) bool {
	tree := g.el(r,c)
	for row := r-1; row >= 0; row-- {
		if g.el(row,c) >= tree {
			return false
		}
	}
	return true
}

func (g grid) visibleFromBottom(r,c int) bool {
	tree := g.el(r,c)
	for row := r+1; row < g.rows(); row++ {
		if g.el(row,c) >= tree {
			return false
		}
	}
	return true
}

func (g grid) visibleFromRight(r,c int) bool {
	tree := g.el(r,c)
	for col := c+1; col < g.cols(); col++ {
		if g.el(r,col) >= tree {
			return false
		}
	}
	return true
}

func (g grid) visibleFromLeft(r,c int) bool {
	tree := g.el(r,c)
	for col := c-1; col >= 0; col-- {
		if g.el(r,col) >= tree {
			return false
		}
	}
	return true
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
