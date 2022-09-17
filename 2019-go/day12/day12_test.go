package day12

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var input string =`<x=-8, y=-18, z=6>
<x=-11, y=-14, z=4>
<x=8, y=-3, z=-10>
<x=-2, y=-16, z=1>`

// https://adventofcode.com/2019/day/12

func TestParsing(t *testing.T) {
	in := parseInput()
	assert.Equal(t, position{-8,-18,6}, in[0])
	assert.Equal(t, position{-11,-14,4}, in[1])
	assert.Equal(t, position{8,-3,-10}, in[2])
	assert.Equal(t, position{-2,-16,1}, in[3])
}

func TestApplyGravity(t *testing.T) {
	ganymede := newMoon(position{3,0,0})
	callisto := newMoon(position{6,0,0})

	changePos1, changePos2 := applyGravity(ganymede, callisto)

	assert.Equal(t, position{1,0,0}, changePos1)
	assert.Equal(t, position{-1,0,0}, changePos2)
}

func TestApplyVelocity(t *testing.T) {
	m := newMoon(position{1,2,3})
	m.velocity = position{3,-1,2}
	
	applyVelocity(m)

	assert.Equal(t, position{3,-1,2}, m.velocity)
	assert.Equal(t, position{4,1,5}, m.position)
}

func TestStep(t *testing.T) {
	assertPosition := func(exp, p position) {
		assert.Equal(t, exp, p, "invalid position")
	}

	assertVelocity := func(exp, p position) {
		assert.Equal(t, exp, p, "invalid velocity")
	}

	initPositions := []position{
		{-1,0,2},
		{2,-10,-7},
		{4,-8,8},
		{3,5,-1},
	}
	s := newSystem(initPositions)
	s.step()

	assertPosition(position{2,-1,1}, s.moons[0].position)
	assertPosition(position{3,-7,-4}, s.moons[1].position)
	assertPosition(position{1,-7,5}, s.moons[2].position)
	assertPosition(position{2,2,0}, s.moons[3].position)

	assertVelocity(position{3,-1,-1}, s.moons[0].velocity)
	assertVelocity(position{1,3,3}, s.moons[1].velocity)
	assertVelocity(position{-3,1,-3}, s.moons[2].velocity)
	assertVelocity(position{-1,-3,1}, s.moons[3].velocity)
}

func TestPart1Example0(t *testing.T) {
	initPositions := []position{
		{-1,0,2},
		{2,-10,-7},
		{4,-8,8},
		{3,5,-1},
	}
	assert.Equal(t, 179, part1(initPositions, 10))
}

func part1(pos []position, steps int) int {
	s := newSystem(pos)
	for i := 0; i < steps; i++ {
		s.step()
	}
	return s.totalEnergy()
}

func TestPart1Example(t *testing.T) {
	initPositions := []position{
		{-8,-10,0},
		{5,5,10},
		{2,-7,3},
		{9,-8,-3},
	}
	assert.Equal(t, 1940, part1(initPositions, 100))
}


func TestPart1(t *testing.T) {
	assert.Equal(t, 9743, part1(parseInput(), 1000))
}

func TestPart2Example(t *testing.T) {
	initPositions := []position{
		{-1,0,2},
		{2,-10,-7},
		{4,-8,8},
		{3,5,-1},
	}
	assert.Equal(t, uint64(2772), part2(initPositions))
}

func TestLeastCommonMultiple(t *testing.T) {
	testCases := []struct {
		a,b, exp uint
	}{
		{4,6,12},
		{12,15,60},
		{6,8,24},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v,%v", tc.a,tc.b), func(t *testing.T) {
			got := lcm(tc.a,tc.b)
			assert.Equal(t, tc.exp, got)
		})
	}
}

func TestPart2Example2(t *testing.T) {
	initPositions := []position{
		{-8,-10,0},
		{5,5,10},
		{2,-7,3},
		{9,-8,-3},
	}
	assert.Equal(t, uint64(4686774924), part2(initPositions))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, uint64(288684633706728), part2(parseInput()))
}

func part2(pos []position) uint64 {
	s := newSystem(pos)

	cntX := counter{coord:"x", set: newSet()}
	cntY := counter{coord:"y", set: newSet()}
	cntZ := counter{coord:"z", set: newSet()}

	for !cntX.found || !cntY.found || !cntZ.found {
		s.step()

		cntX.iterate(s.moons)
		cntY.iterate(s.moons)
		cntZ.iterate(s.moons)
	}

	return uint64(lcm(cntX.cnt, lcm(cntY.cnt, cntZ.cnt)))
}

type counter struct {
	cnt uint
	found bool
	coord string
	set *set
}

func (c *counter) iterate(moons []moon) {
	if c.found {
		return
	}
	key := c.makeKey(moons)
	
	if c.set.contains(key) {
		c.found = true
		return
	}
	c.set.add(key)
	c.cnt++
}

func (c *counter) makeKey(moons []moon) key {
	if c.coord == "x" {
		return key{
			pos0:  moons[0].position.x,
			velo0: moons[0].velocity.x,

			pos1:  moons[1].position.x,
			velo1: moons[1].velocity.x,

			pos2:  moons[2].position.x,
			velo2: moons[2].velocity.x,

			pos3:  moons[3].position.x,
			velo3: moons[3].velocity.x,
		}
	} else if c.coord == "y" {
		return key{
			pos0:  moons[0].position.y,
			velo0: moons[0].velocity.y,

			pos1:  moons[1].position.y,
			velo1: moons[1].velocity.y,

			pos2:  moons[2].position.y,
			velo2: moons[2].velocity.y,

			pos3:  moons[3].position.y,
			velo3: moons[3].velocity.y,
		}
	}

	return key{
		pos0:  moons[0].position.z,
		velo0: moons[0].velocity.z,

		pos1:  moons[1].position.z,
		velo1: moons[1].velocity.z,

		pos2:  moons[2].position.z,
		velo2: moons[2].velocity.z,

		pos3:  moons[3].position.z,
		velo3: moons[3].velocity.z,
	}
}

type position struct {
	x,y,z int
}

func (p *position) addChanges(pos position) {
	p.x += pos.x
	p.y += pos.y
	p.z += pos.z
}

func parseInput() []position {
	out := []position{}
	reg := regexp.MustCompile(`<\w=(\-?\d+), \w=(\-?\d+), \w=(\-?\d+)>`)
	splitted := strings.Split(input, "\n")
	for _, v := range splitted {
		for _, match := range reg.FindAllStringSubmatch(v,-1) {
			x,_ := strconv.Atoi(match[1])
			y,_ := strconv.Atoi(match[2])
			z,_ := strconv.Atoi(match[3])
			out = append(out, position{x,y,z})
		}
	}
	return out
}

type moon struct {
	position position
	velocity position
}

func newMoon(position position) *moon {
	return &moon{position: position}
}

func (m moon) energy() int {
	kinetic := math.Abs(float64(m.velocity.x)) + math.Abs(float64(m.velocity.y)) + math.Abs(float64(m.velocity.z))
	potential := math.Abs(float64(m.position.x)) + math.Abs(float64(m.position.y)) + math.Abs(float64(m.position.z))

	return int(potential) * int(kinetic)
}

func applyGravity(m1 *moon, m2 *moon) (position,position) {
	applyPosition := func(coord1 int, coord2 int) (int,int) {
		if coord1 < coord2 {
			return 1, -1
		} else if coord1 > coord2 {
			return -1, 1
		}
		return 0,0
	}
	changePos1 := position{}
	changePos2 := position{}

	changePos1.x, changePos2.x = applyPosition(m1.position.x, m2.position.x)
	changePos1.y, changePos2.y = applyPosition(m1.position.y, m2.position.y)
	changePos1.z, changePos2.z = applyPosition(m1.position.z, m2.position.z)

	return changePos1, changePos2
}

func applyVelocity(m *moon) {
	(&m.position).addChanges(m.velocity)
}

type system struct {
	moons []moon
}

func newSystem(positions []position) *system {
	moons := []moon{}
	for i := 0; i < len(positions); i++ {
		moons = append(moons, *newMoon(positions[i]))
	}
	return &system{moons}
}

// todo: handle combinations algorithmically and any number of inputs
func (s *system) step() {
	m := s.moons

	changeVel0 := &position{}
	changeVel1 := &position{}
	changeVel2 := &position{}
	changeVel3 := &position{}

	cl0, cl1 := applyGravity(&m[0], &m[1])
	changeVel0.addChanges(cl0)
	changeVel1.addChanges(cl1)

	cl0, cl2 := applyGravity(&m[0], &m[2])
	changeVel0.addChanges(cl0)
	changeVel2.addChanges(cl2)

	cl0, cl3 := applyGravity(&m[0], &m[3])
	changeVel0.addChanges(cl0)
	changeVel3.addChanges(cl3)

	cl1,cl2 = applyGravity(&m[1], &m[2])
	changeVel1.addChanges(cl1)
	changeVel2.addChanges(cl2)

	cl1, cl3 = applyGravity(&m[1], &m[3])
	changeVel1.addChanges(cl1)
	changeVel3.addChanges(cl3)

	cl2,cl3 = applyGravity(&m[2], &m[3])
	changeVel2.addChanges(cl2)
	changeVel3.addChanges(cl3)

	(&m[0].velocity).addChanges(*changeVel0)
	(&m[1].velocity).addChanges(*changeVel1)
	(&m[2].velocity).addChanges(*changeVel2)
	(&m[3].velocity).addChanges(*changeVel3)

	applyVelocity(&m[0])
	applyVelocity(&m[1])
	applyVelocity(&m[2])
	applyVelocity(&m[3])
}

func (s *system) totalEnergy() int {
	sum := 0
	for _, v := range s.moons {
		sum += v.energy()
	}
	return sum
}

type void struct{}
type key struct {
	pos0, velo0 int
	pos1, velo1 int
	pos2, velo2 int
	pos3, velo3 int
}
type set struct {
	d map[key]void
}
func newSet() *set { return &set{map[key]void{}} }

func (s *set) contains(k key) bool {
	_, ok := s.d[k]
	return ok
}

func (s *set) add(k key) {
	var v void 
	s.d[k] = v
}


func lcm(a,b uint) uint {
	return (a*b)/gcd(a,b)
}

func gcd(a,b uint) uint {
	if a == b {
		return a
	}

	if a > b{
		return gcd(a-b,b)
	} 
	return gcd(a, b-a)
}