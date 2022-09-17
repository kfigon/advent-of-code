package d6

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestExample(t *testing.T) {
	input := `3,4,3,1,2`
	tcs := []struct {
		input int
		expected int
	} {
		{18,26},
		{1,5},
		{2,6},
		{3,7},
		{10,12},
		{80,5934},
		{256,26984457539},
	}
	for _, tc := range tcs {
		t.Run(fmt.Sprintf("%v",tc.input), func(t *testing.T) {
			got := solve(parse(input), tc.input)
			if got != tc.expected {
				t.Error("Day", tc.input, "got", got,"exp",tc.expected)
			}
		})
	}
	
}
func readFile(t *testing.T) string {
	d, err := os.ReadFile("data.txt")
	if err != nil {
		t.Fatal(err)
		return ""
	}
	return string(d)
}
func TestP1(t *testing.T) {
	got := solve(parse(readFile(t)), 80)
	if got != 352151 {
		t.Error("Got", got)
	}
}

func TestP2(t *testing.T) {
	got := solve(parse(readFile(t)), 256)
	if got != 1601616884019 {
		t.Error("Got", got)
	}
}

func parse(in string) []int {
	numStr := strings.Split(in, ",")
	out := []int{}
	for _, v := range numStr {
		num, _ := strconv.Atoi(v)
		out = append(out, num)
	}
	return out
}

func solve(fishes []int, days int) int {
	fishMap := map[int]int{}
	for _, f := range fishes {
		fishMap[f]++
	}

	fishPopulation := len(fishes)
	for i := 0; i < days; i++ {
		newMap := map[int]int{}
		for fishDay, count := range fishMap { // modifying what we're iterate is bad idea...
			newMap[fishDay-1] = count
		}

		fishMap = newMap
		newFishes := fishMap[-1]
		fishMap[8] = newFishes
		fishMap[6] += newFishes
		fishPopulation += newFishes
		delete(fishMap, -1)
	}
	return fishPopulation
}