package d2

import (
	"testing"
	"os"
	"strconv"
	"strings"
)

func TestPart1(t *testing.T) {
	got := solveP1(parse(t))
	if got != 2272262 {
		t.Error("Invalid p1", got, "exp", 2272262)
	}
}

func TestPart2(t *testing.T) {
	got := solveP2(parse(t))
	if got != 2134882034 {
		t.Error("Invalid p1", got, "exp", 2134882034)
	}
}

type data struct {
	cmd string
	step int
}

func parse(t *testing.T) []data {
	var out []data
	d, err := os.ReadFile("data.txt")
	if err != nil {
		t.Fatal("Error in reading file",err)
		return out
	}
	lines := strings.Split(string(d), "\r\n")
	for _, line := range lines {
		splitted := strings.Split(line, " ")
		if len(splitted) < 2 {
			t.Fatal("Error in file, line:", line)
			return nil
		}
		if splitted[0] != "forward" && splitted[0] != "down" && splitted[0] != "up" {
			t.Fatal("Invalid cmd in line:", line)
			return nil
		}
		v, err := strconv.Atoi(splitted[1])
		if err != nil {
			t.Fatal("Error in parsing step, line:", line, "err:",err)
			return nil
		}
		out = append(out, data{splitted[0], v})
	}
	return out
}

func solveP1(d []data) int {
	horizontal := 0
	depth := 0
	for _, step := range d {
		switch step.cmd {
		case "forward": horizontal += step.step
		case "down": depth += step.step
		case "up": depth -= step.step
		}
	}
	return horizontal*depth
}

func solveP2(d []data) int {
	horizontal := 0
	depth := 0
	aim := 0
	for _, step := range d {
		switch step.cmd {
		case "forward": {
			horizontal += step.step
			depth += aim*step.step
		}
		case "down": aim += step.step
		case "up": aim -= step.step
		}
	}
	return horizontal*depth
}