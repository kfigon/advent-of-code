package d9

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"testing"
)

const exampleData string = `2199943210
3987894921
9856789892
8767896789
9899965678`

func TestExampleP1(t *testing.T) {
	if got := solveP1(parse(exampleData, "\n")); got != 15 {
		t.Error("Exp", 15, "got", got)
	}
}

func TestP1(t *testing.T) {
	d,err:=os.ReadFile("data.txt")
	if err != nil {
		t.Fatal(err)
		return
	}
	if got := solveP1(parse(string(d), "\r\n")); got != 585 {
		t.Error("Exp", 585, "got", got)
	}
}
func TestExampleP2Basins(t *testing.T) {
	data := parse(exampleData, "\n")
	tdt := []struct {
		p pair
		exp int
	}{
		{pair{r:0, c: 1},3},
		{pair{r:0, c: 9},9},
		{pair{r:2, c: 2},14},
		{pair{r:4, c: 6},9},
	}
	for i, tc := range tdt {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			v, ok := basin(data, tc.p)
			if !ok {
				t.Error("Expected basin, not found")
			}
			if v != tc.exp {
				t.Error("Exp",tc.exp,"got",v)
			}
		})
	}
}

func TestExampleP2(t *testing.T) {
	data := parse(exampleData, "\n")
	got := solveP2(data)
	if got != 1134 {
		t.Error("Exp", 1134, "got", got)
	}
}

func TestP2(t *testing.T) {
	d,err := os.ReadFile("data.txt")
	if err != nil {
		t.Fatal(err)
		return
	}
	got := solveP2(parse(string(d), "\r\n"))
	exp := 827904
	if got != exp {
		t.Error("Exp", exp, "got", got)
	}
}

func parse(data string, sep string) [][]int {
	lines := strings.Split(data,sep)
	out := [][]int{}
	for _, line := range lines {
		row := []int{}
		for _, c := range line {
			v, _ := strconv.Atoi(string(c))
			row = append(row, v)
		}
		out = append(out, row)
	}
	return out
}

type pair struct {
	r int
	c int
}

func inBounds(data [][]int, c pair) bool {
	return c.r >= 0 && c.r < len(data) && c.c >= 0 && c.c < len(data[0])
}

func neighbours(data [][]int, p pair) []pair {
	out := []pair{}

	row := p.r
	col := p.c
	candidates := []pair {
			{row-1, col},
		{row, col-1}, {row, col+1},
			{row+1, col},
	}

	for _, c := range candidates {
		if inBounds(data, c) {
			out = append(out, c)
		}
	}
	return out
}

func isLowPoint(data [][]int, p pair) bool {
	current := data[p.r][p.c]
	nei := neighbours(data, p)

	for _, v := range nei {
		value := data[v.r][v.c]
		if current >= value {
			return false
		}
	}
	return true
}

func findLowPoints(data [][]int) []pair {
	points := []pair{}
	for r := 0; r < len(data); r++ {
		for c := 0; c < len(data[0]); c++ {
			v := pair{r:r,c:c}
			if isLowPoint(data, v) {
				points = append(points, v)
			}
		}
	}
	return points
}

func solveP1(data [][]int) int {
	lowPoints := findLowPoints(data)

	sum := 0
	for _, v := range lowPoints {
		current := data[v.r][v.c]
		sum += current+1
	}
	return sum
}

func solveP2(data [][]int) int {
	basinSizes := []int{}
	lowPoints := findLowPoints(data)
	for _, v := range lowPoints {
		if v, ok := basin(data, v); ok {
			basinSizes = append(basinSizes, v)
		} 
	}

	sort.Ints(basinSizes)
	threeBiggest := basinSizes[len(basinSizes)-3:]
	result := 1
	for _, v := range threeBiggest {
		result *= v
	}
	return result
}

func basin(data[][]int, current pair) (int, bool) {
	basinCandidates := newQueue()

	addNeighbours := func(p pair) {
		for _, v := range neighbours(data, p) {
			candidateValue := data[v.r][v.c]
			currentValue := data[p.r][p.c]

			if candidateValue != 9 && candidateValue > currentValue {
				basinCandidates.enqueue(v)
			}
		}
	}
	
	basin := map[pair]struct{}{}
	basin[current] = struct{}{}

	// recursive also can work
	addNeighbours(current)
	for !basinCandidates.empty() {
		el := basinCandidates.dequeue()
		basin[el]=struct{}{}
		addNeighbours(el)
	}
	return len(basin),true
}

type queue struct {
	d []pair
}

func newQueue() *queue {
	return &queue{
		d: []pair{},
	}
}

func (s *queue) enqueue(p pair) {
	s.d = append(s.d, p)
}

func (s *queue) empty() bool {
	return len(s.d) == 0
}

func (s *queue) dequeue() pair {
	fist := s.d[0]
	s.d = s.d[1:]
	return fist
}