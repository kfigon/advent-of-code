package day9

import (
	"aoc2019/intcode"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// https://adventofcode.com/2019/day/9

func readFile(t *testing.T) []int {
	file, err := os.Open("data.txt")
	require.NoError(t, err)
	defer file.Close()

	content, err := io.ReadAll(file)
	require.NoError(t, err)

	splitted := strings.Split(string(content),",")
	out := make([]int, 0)
	for i := 0; i < len(splitted); i++ {
		v, err := strconv.Atoi(string(splitted[i]))
		require.NoError(t, err)
		
		out = append(out, v)
	}
	return out
	
}
func TestPart1(t *testing.T) {
	file := readFile(t)
	c := intcode.NewComputer(file)
	c.SetUserInput(1)
	c.Calc()
	assert.Equal(t, 3601950151, c.GetOutput())
}

func TestPart2(t *testing.T) {
	file := readFile(t)
	c := intcode.NewComputer(file)
	c.SetUserInput(2)
	c.Calc()
	assert.Equal(t, 64236, c.GetOutput())
}