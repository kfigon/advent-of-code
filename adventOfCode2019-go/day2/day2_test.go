package day2

import (
	"strings"
	"io"
	"strconv"
	"os"
	"testing"
	"github.com/stretchr/testify/assert"
	"aoc2019/intcode"
)

func readFile() []int {
	file, err := os.Open("data.txt")
	if err != nil {
		return []int{}
	}
	defer file.Close()

	out := make([]int,0)
	bytes, err := io.ReadAll(file)
	if err != nil {
		return []int{}
	}
	str := string(bytes)
	for _, v := range strings.Split(str, ",") {
		i, _ := strconv.Atoi(v)
		out = append(out, i)
	}
	return out
}
func TestPart1(t *testing.T) {
	data := readFile()
	data[1] = 12
	data[2] = 2
	out := intcode.NewComputer(data).Calc()

	assert.Equal(t, 5866714, out[0])
}

func TestPart2(t *testing.T) {
	exp := 5208
	assert.Equal(t, exp, findPart2Result())
}

func findPart2Result() int {
	const valueToFind int = 19690720
	fileData := readFile()
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			
			data := make([]int, len(fileData))
			copy(data, fileData)

			data[1] = noun
			data[2] = verb
			out := intcode.NewComputer(data).Calc()
			if out[0] == valueToFind {
				return 100*out[1]+out[2]
			}
		}
	}
	return -1
}
