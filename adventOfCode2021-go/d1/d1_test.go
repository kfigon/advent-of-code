package d1

import (
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	got := solveP1(parseData(t))
	if got != 1624 {
		t.Error("Invalid p1", got)
	}
}

func TestPart2(t *testing.T) {
	got := solveP2(parseData(t))
	if got != 1653 {
		t.Error("Invalid p2", got)
	}
}

func parseData(t *testing.T) []int {
	data, err := os.ReadFile("data.txt")
	if err != nil {
		t.Fatal("Error in reading file", err)
		return []int{}
	}
	var out []int
	str := strings.Split(string(data), "\r\n")
	for _, line := range str {
		v, err := strconv.Atoi(line)
		if err != nil {
			t.Fatal("Error parsing data", err)
			return out
		}
		out = append(out, v)
	}
	return out
}

func solveP1(data []int) int {
	increses := 0
	for i := 0; i < len(data)-1; i++ {
		this := data[i]
		next := data[i+1]
		if next > this {
			increses++
		}
	}
	return increses
}

func solveP2(data []int) int {
	increses := 0
	for i := 0; i < len(data)-3; i++ {
		thisSum := data[i]+data[i+1]+data[i+2]
		nextSum := data[i+1]+data[i+2]+data[i+3]
		if nextSum > thisSum {
			increses++
		}
	}
	return increses
}