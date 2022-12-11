package d10

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func readFile(name string, t *testing.T) string {
	raw, err := os.ReadFile(name)
	require.NoError(t, err)
	return string(raw)
}

func TestBigExample(t *testing.T) {
	instructions := parse(strings.Split(readFile("example.txt", t), "\r\n"))
	t.Run("p1", func(t *testing.T) {
		assert.Equal(t, 13140, p1(instructions))
	})

	t.Run("p2", func(t *testing.T) {
		assert.Equal(t, expectedP2, p2(instructions))
	})
}

func TestFile(t *testing.T) {
	instructions := parse(strings.Split(readFile("data.txt", t), "\r\n"))
	t.Run("p1", func(t *testing.T) {
		assert.Equal(t, 13180, p1(instructions))
	})

	t.Run("p2", func(t *testing.T) {
		t.Fail()
	})
}

func TestPositionMapping(t *testing.T) {
	testCases := []struct {
		in int
		exp int
	}{
		{1,1},
		{2,2},
		{5,5},
		{10,10},
		{20,20},
		{40,40},
		
		{41,1},
		{51,11},
		{80,40},
		
		{81,1},
		{85,5},
		{120,40},
		
		{201,1},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprintf("%v->%v", tC.in, tC.exp), func(t *testing.T) {
			assert.Equal(t, tC.exp, registerToPosition(tC.in))
		})
	}
}

const example = `noop
addx 3
addx -5`

type instruction struct {
	cmd string
	val int
}

func parse(lines []string) []instruction {
	out := []instruction{}
	for _, line := range lines {
		parts := strings.Split(line, " ")
		
		if len(parts) == 1 && parts[0] == "noop" {
			out = append(out, instruction{cmd: "noop"})
		} else if len(parts) == 2 && parts[0] == "addx" {
			v, _ := strconv.Atoi(parts[1])
			out = append(out, instruction{cmd: "addx", val: v})
		}
	}

	return out
}

type cpu struct {
	register int
	cpuCycle int
	duringCycleCallback func(cpu)
	endOfCycleCallback func(cpu)
}

func newCpu() *cpu {
	return &cpu{
		register: 1,
		cpuCycle: 1,
	}
}

func (c *cpu) singleCycle(fn func()) {
	if c.duringCycleCallback != nil {
		c.duringCycleCallback(*c)
	}

	c.cpuCycle++
	fn()
	if c.endOfCycleCallback != nil {
		c.endOfCycleCallback(*c)
	}
}

func (c *cpu) processInstruction(instr instruction) {
	if instr.cmd == "noop" {
		c.singleCycle(func() {
		})
	} else if instr.cmd == "addx" {
		c.singleCycle(func() {
		})

		c.singleCycle(func() {
			c.register += instr.val	
		})
	}
}

func p1(instructions []instruction) int {
	probePoints := map[int]bool{
		20: true,
		60: true,
		100: true, 
		140: true, 
		180: true,
		220: true,
	}
	c := newCpu()

	power := 0
	c.endOfCycleCallback = func(c cpu) {
		if probePoints[c.cpuCycle] {
			power += c.cpuCycle * c.register
		}
	}

	for _, instr := range instructions {
		c.processInstruction(instr)
	}
	return power
}

func registerToPosition(reg int) int {
	zeroBased := reg-1
	return (zeroBased % 40) + 1
}

func p2(instructions []instruction) string {
	c := newCpu()
	output := ""

	c.duringCycleCallback = func(c cpu) {
		if c.cpuCycle == 200 {
			fmt.Println("puapka na dzika")
		}
		drawn := false
		middlePixelPosition := registerToPosition(c.register)+1
		cycle := c.cpuCycle % 40

		if cycle == middlePixelPosition - 1 {
			fmt.Println(c.cpuCycle," -> #", ", pixel", middlePixelPosition-1)
			output += "#"
			drawn = true
		}

		if cycle == middlePixelPosition {
			fmt.Println(c.cpuCycle," -> #", ", sprite on", middlePixelPosition)
			output += "#"
			drawn = true
		}
		if cycle == middlePixelPosition + 1 {
			fmt.Println(c.cpuCycle," -> #", ", sprite on", middlePixelPosition)
			output += "#"
			drawn = true
		}

		if !drawn {
			fmt.Println(c.cpuCycle," -> .", ", sprite on", middlePixelPosition)
			output += "."
		}
		
		if cycle == 0 {
			fmt.Println("newline on cycle", c.cpuCycle)
			output += "\n"
		}
	}
	
	c.endOfCycleCallback = func(c cpu) {
		fmt.Println("register", c.register)
	}

	for _, instr := range instructions {
		c.processInstruction(instr)
	}

	return output
}

const expectedP2 = `##..##..##..##..##..##..##..##..##..##..
###...###...###...###...###...###...###.
####....####....####....####....####....
#####.....#####.....#####.....#####.....
######......######......######......####
#######.......#######.......#######.....`