package d10

import (
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
		expectedResult := strings.ReplaceAll(readFile("expectedResult.txt", t),"\r","")
		assert.Equal(t, expectedResult, p2(instructions))
	})
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

func p2(instructions []instruction) string {
	c := newCpu()
	output := ""
	const width = 40
	const numberOfLines = 6

	c.duringCycleCallback = func(c cpu) {
		zeroBasedCpu := c.cpuCycle-1
		zeroReg := c.register-1

		drawn := false
		middlePixelPosition := (zeroReg % width) + 1
		cycle := zeroBasedCpu % width

		if cycle == middlePixelPosition - 1 {
			output += "#"
			drawn = true
		} else if cycle == middlePixelPosition {
			output += "#"
			drawn = true
		} else if cycle == middlePixelPosition + 1 {
			output += "#"
			drawn = true
		}

		if !drawn {
			output += "."
		}

		// last row and not last line
		if cycle == (width-1) && c.cpuCycle != width*numberOfLines {
			output += "\n"
		}
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