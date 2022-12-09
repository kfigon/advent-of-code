package d9

import (
	"math"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const example = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`


const example2 = `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`

func TestExample(t *testing.T) {
	rules := parse(strings.Split(example, "\n"))
	t.Run("p1", func(t *testing.T) {
		assert.Equal(t, 13, p1(rules))
	})

	t.Run("p2", func(t *testing.T) {
		assert.Equal(t, 1, p2(parse(strings.Split(example, "\n"))))
	})

	t.Run("p2 - 2", func(t *testing.T) {
		assert.Equal(t, 36, p2(parse(strings.Split(example2, "\n"))))
	})
}

func TestFile(t *testing.T) {
	raw, err := os.ReadFile("data.txt")
	require.NoError(t, err)

	rules := parse(strings.Split(string(raw), "\r\n"))
	t.Run("p1", func(t *testing.T) {
		assert.Equal(t, 5930, p1(rules))
	})

	t.Run("p2", func(t *testing.T) {
		assert.Less(t, p2(rules), 4581)
		assert.Equal(t, -1, p2(rules))
	})
}

type coord struct {
	x int
	y int
}

func (c coord) neighbours(other coord) bool {
	xDif := math.Abs(float64(c.x) - float64(other.x))
	yDif := math.Abs(float64(c.y) - float64(other.y))

	return xDif <= 1 && yDif <= 1
}

type direction int
const (
	up direction = iota
	down
	left
	right
)

type rule struct {
	d direction
	step int
}

func parse(lines []string) []rule {
	rules := []rule{}
	for _, line := range lines {
		parts := strings.Split(line, " ")
		step,_ := strconv.Atoi(parts[1])
		var dir direction
		switch parts[0] {
		case "R": dir = right
		case "U": dir = up
		case "L": dir = left
		case "D": dir = down
		}

		r := rule {
			d: dir,
			step: step,
		}
		rules = append(rules, r)
	}
	return rules
}

type rope struct {
	parts []coord
}

func newRope(length int) *rope {
	return &rope{
		make([]coord, length),
	}
}

func (r *rope) tail() coord {
	return r.parts[len(r.parts)-1]
}

func (r *rope) move(d direction) {
	// pixel grid indexing
	moveOneField := func(c *coord) {
		switch d {
		case up:
			c.y--
		case down:
			c.y++
		case left:
			c.x--
		case right:
			c.x++
		}
	}
	prev := r.parts[0]
	moveOneField(&r.parts[0])
	for i := 1; i < len(r.parts); i++ {
		if !r.parts[i-1].neighbours(r.parts[i]){
			tmp := r.parts[i]
			r.parts[i] = prev
			prev = tmp
		}
	}
}

func processRules(rules []rule, rop *rope) int {
	tailPosition := map[coord]struct{}{}
	tailPosition[rop.tail()]= struct{}{}

	for _, r := range rules {
		for i := 0; i < r.step; i++ {
			rop.move(r.d)
			tailPosition[rop.tail()]=struct{}{}
		}
	}
	return len(tailPosition)
}

func p1(rules []rule) int {
	return processRules(rules, newRope(2))
}

func p2(rules []rule) int {
	return processRules(rules, newRope(10))
}