package d8

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
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
}

type grid [][]int

func (g grid) col() int {
	return len(g[0])
}

func (g grid) rows() int {
	return len(g)
}

func (g grid) el(c,r int) int {
	return g[c][r]
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
	return 0
}
