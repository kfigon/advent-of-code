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
		t.Fail()
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

func p1(instructions []instruction) int {
	probePoints := map[int]bool{
		20: true,
		60: true,
		100: true, 
		140: true, 
		180: true,
		220: true,
	}
	register := 1
	cpuCycle := 1
	power := 0

	callback := func(c int) {
		if probePoints[c] {
			power += c * register
		}
	}

	for _, instr := range instructions {
		if instr.cmd == "noop" {
			increaseCpu(&cpuCycle, callback)
		} else if instr.cmd == "addx" {
			increaseCpu(&cpuCycle, callback)
			register += instr.val
			increaseCpu(&cpuCycle, callback)
		}
	}

	return power
}

func increaseCpu(cpu *int, callback func(int)) {
	*cpu++
	callback(*cpu)
}