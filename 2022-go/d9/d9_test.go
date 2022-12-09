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

func TestExample(t *testing.T) {
	rules := parse(strings.Split(example, "\n"))
	t.Run("p1", func(t *testing.T) {
		assert.Equal(t, 13, p1(rules))
	})

	t.Run("p2", func(t *testing.T) {
		t.Fail()
	})
}

func TestFile(t *testing.T) {
	raw, err := os.ReadFile("data.txt")
	require.NoError(t, err)

	rules := parse(strings.Split(string(raw), "\r\n"))
	t.Run("p1", func(t *testing.T) {
		// too high
		assert.Equal(t, 5930, p1(rules))
	})

	t.Run("p2", func(t *testing.T) {
		t.Fail()
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

func(d direction) String() string {
	return []string{"up", "down", "left", "right"}[d]
}

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
	head coord
	tail coord
}

func p1(rules []rule) int {
	// pixel grid indexing
	moveOneField := func(c *coord, d direction) {
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

	tailPosition := map[coord]struct{}{}
	rop := &rope{}
	tailPosition[rop.tail]= struct{}{}
	for _, r := range rules {
		for i := 0; i < r.step; i++ {

			oldHead := rop.head
			moveOneField(&rop.head, r.d)
			if !rop.tail.neighbours(rop.head){
				rop.tail = oldHead
				tailPosition[rop.tail]=struct{}{}
			}
		}
	}
	return len(tailPosition)
}
