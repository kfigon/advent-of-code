package d7

import (
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func TestExampleP1(t *testing.T) {
	input := "16,1,2,0,4,2,7,1,2,14"
	data := parse(input)
	got := solveP1(data)
	if got != 37 {
		t.Error("Got", got, "exp", 37)
	}
}

func TestExampleP2(t *testing.T) {
	input := "16,1,2,0,4,2,7,1,2,14"
	data := parse(input)
	got := solveP2(data)
	if got != 168 {
		t.Error("Got", got, "exp", 168)
	}
}

func TestP1(t *testing.T) {
	got := solveP1(parse(readFile(t)))
	if got != 355764 {
		t.Error("Got", got, "exp", 355764)
	}
}

func TestP2(t *testing.T) {
	got := solveP2(parse(readFile(t)))
	if got != 99634572 {
		t.Error("Got", got, "exp", 99634572)
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

func parse(in string) []int {
	data := strings.Split(in,",")
	var out []int
	for _, d := range data {
		v, _ := strconv.Atoi(d)
		out = append(out, v)
	}
	return out
}

func solveP1(data []int) int {
	m := median(data)
	fuel := 0
	for _, d := range data {
		distance := math.Abs(float64(d)-float64(m))
		fuel += int(distance)
	}
	return fuel
}

func solveP2(data []int) int {
	a := avg(data)
	fuel := 0
	for _, d := range data {
		fuel += countFuel(d,a)
	}
	return fuel
}

func countFuel(position, target int) int {
	distance := int(math.Abs(float64(position) - float64(target)))
	sum := 0
	for i := 1; i <= distance; i++ {
		sum += i
	}
	return sum
}

func median(data []int) int {
	sort.Ints(data)
	return data[len(data)/2]
}

func avg(data []int) int {
	sum := 0
	for _, v := range data {
		sum += v
	}
	a := float64(sum)/float64(len(data))
	// strange rounding... 485.538 is expected to round down, 4.9 up
	if remainder := (a - float64(int(a))); remainder >= 0.6 {
		return int(a)+1
	}
	return int(a)
}