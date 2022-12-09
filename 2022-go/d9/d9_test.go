package d9

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
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

type coord struct {
	x int
	y int
}

func(c coord) samePlane(other coord) bool {
	return c.x == other.x || c.y == other.y
}

func(c coord) samePosition(other coord) bool {
	return c.x == other.x && c.y == other.y
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
	head coord
	tail coord
}

func(r *rope) move(rul rule) {
	if r.head.samePosition(r.tail) {
		
	} else if r.head.samePlane(r.tail){

	} else { // diagonal

	}

	// pixel grid indexing
	switch rul.d {
	case up:
		r.head.y -= rul.step
	case down:
		r.head.y += rul.step
	case left:
		r.head.x -= rul.step
	case right:
		r.head.x += rul.step
	}
}

func p1(rules []rule) int {
	tailPosition := map[coord]struct{}{}
	rop := &rope{}
	for _, r := range rules {
		rop.move(r)
		tailPosition[rop.tail]=struct{}{}
	}
	return len(tailPosition)
}
