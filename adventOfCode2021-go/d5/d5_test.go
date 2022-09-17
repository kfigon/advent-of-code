package d52021

import (
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestExampleP1(t *testing.T) {
	input:=`0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

	got := solveP1(parse(input, "\n"))
	if got != 5 {
		t.Error("Invalid p1 example, got", got)
	}
}

func TestExampleP2(t *testing.T) {
	input:=`0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

	got := solveP2(parse(input, "\n"))
	if got != 12 {
		t.Error("Invalid p2 example, got", got)
	}
}

func TestP1(t *testing.T) {
	data, err := os.ReadFile("data.txt")
	if err != nil {
		t.Fatal("error in opening file:",err)
		return
	}

	got := solveP1(parse(string(data), "\r\n"))
	if got != 6856 {
		t.Error("Invalid p1, got", got)
	}
}

func TestP2(t *testing.T) {
	data, err := os.ReadFile("data.txt")
	if err != nil {
		t.Fatal("error in opening file:",err)
		return
	}

	got := solveP2(parse(string(data), "\r\n"))
	if got != 20666 {
		t.Error("Invalid p1, got", got)
	}
}

func TestInterpolate(t *testing.T) {
	testCases := []struct {
		desc	string
		lin line
		exp []point
	}{
		{
			desc: "1,1->3,3",
			lin: line{point{1,1}, point{3,3}},
			exp: []point{{1,1},{2,2},{3,3}},
		},
		{
			desc: "9,7->7,9",
			lin: line{point{9,7}, point{7,9}},
			exp: []point{{7,9}, {8,8},{9,7}},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := tC.lin.intermediatePointsWithDiagonals()
			if len(got) != len(tC.exp) {
				t.Fatalf("Invalid length, exp %v, got %v", len(tC.exp), len(got))
			}
			for i := 0; i < len(got); i++ {
				e := tC.exp[i]
				g := got[i]
				if e != g {
					t.Errorf("Error on idx %v, got %v, exp %v", i, g, e)
				}
			}
		})
	}
}

type point struct {
	x int
	y int
}

func parse(in string, sep string) []line {
	lines := strings.Split(in, sep)
	out := []line{}
	for _, lineStr := range lines {
		pts := strings.Split(lineStr, " -> ")
		start := parsePoint(pts[0])
		end := parsePoint(pts[1])

		out = append(out, line{start, end})
	}
	return out
}

func parsePoint(in string) point {
	data := strings.Split(in, ",")
	x, _ := strconv.Atoi(data[0])
	y, _ := strconv.Atoi(data[1])
	return point{x,y}
}

type line struct {
	start point
	end point
}

func (li line) isHorizontal() bool {
	return li.start.y == li.end.y
}

func (li line) isVertical() bool {
	return li.start.x == li.end.x
}

func (li line) intermediatePoints() []point {
	out := []point{}
	if li.isHorizontal() {
		startX,endX := order(li.start.x, li.end.x)
		for x := startX; x <= endX; x++ {
			out = append(out, point{x, li.start.y})
		}
	} else if li.isVertical() {
		startY,endY := order(li.start.y, li.end.y)		
		for y := startY; y <= endY; y++ {
			out = append(out, point{li.start.x, y})
		}
	}
	return out
}

func (li line) intermediatePointsWithDiagonals() []point {
	out := li.intermediatePoints()
	// diagonal
	if !li.isHorizontal() && !li.isVertical() {
		a,b := calcStuff(li)
		startX,endX := order(li.start.x, li.end.x)
		for x := startX; x <= endX; x++ {
			out = append(out, calcPoint(a,b,x))
		}
	}
	return out
}

func calcPoint(a,b float64, x int) point {
	return point{x:x, y:int(a*float64(x)+b)}
}

func calcStuff(li line) (float64, float64) {
	ya := float64(li.start.y)
	yb := float64(li.end.y)
	xa := float64(li.start.x)
	xb := float64(li.end.x)
	nom := float64(ya-yb)
	denom := float64(xa-xb)

	return nom/denom, ya - ((xa*nom)/denom)
}

func order(a,b int) (int,int) {
	if a < b {
		return a,b
	}
	return b,a
}

func solveP1(lines []line) int {
	return solve(lines, func(l line) []point {return l.intermediatePoints()})
}

func solve(lines []line, extrapolator func(line)[]point) int {
	pts := map[point]int{}
	for _, line := range lines {
		intermediate := extrapolator(line)
		for _, i := range intermediate {
			pts[i]++
		}
	}
	out := 0
	for _, v := range pts {
		if v >= 2 {
			out++
		}
	}
	return out
}

func solveP2(lines []line) int {
	return solve(lines, func(l line) []point {return l.intermediatePointsWithDiagonals()})
}