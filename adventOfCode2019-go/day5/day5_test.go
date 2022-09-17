package day5

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

// https://adventofcode.com/2019/day/5
func TestPart1(t *testing.T) {
	data := readFile(t)
	c := intcode.NewComputer(data)
	c.SetUserInput(1)
	c.Calc()

	assert.Equal(t, 9431221, c.GetOutput())
}

func TestPart2(t *testing.T) {
	data := readFile(t)
	c := intcode.NewComputer(data)
	c.SetUserInput(5)
	c.Calc()

	assert.Equal(t, 1409363, c.GetOutput())
}

func readFile(t *testing.T) []int {
	file, err := os.Open("data.txt")
	require.NoError(t, err)
	defer file.Close()

	content, err := io.ReadAll(file)
	require.NoError(t, err)
	data := string(content)
	splitted := strings.Split(data, ",")
	out := make([]int, len(splitted))
	for i := 0; i < len(out); i++ {
		v, err := strconv.Atoi(splitted[i])
		require.NoError(t, err)
		out[i] = v
	}
	return out
}
