package d3

import (
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestPart1Example(t *testing.T) {
	in := `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`
	got := solveP1(strings.Split(in,"\n"))
	if got != 198 {
		t.Error("Invalid p1", got, "exp", 198)
	}
}

func TestPart2Example(t *testing.T) {
	in := `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`
	got := solveP2(strings.Split(in,"\n"))
	if got != 230 {
		t.Error("Invalid p2", got, "exp", 230)
	}
}
func TestPart1(t *testing.T) {
	got := solveP1(parse(t))
	if got != 2261546 {
		t.Error("Invalid p1", got, "exp", 2261546)
	}
}

func TestPart2(t *testing.T) {
	got := solveP2(parse(t))
	if got != 6775520 {
		t.Error("Invalid p2", got, "exp", 6775520)
	}
}

func TestGrouping(t *testing.T) {
	occ := group(0, []string{
		"00100",
"11110",
"11100",
"10000", })

	if got := len(occ['0']); got != 1 {
		t.Error("Invalid for 0, got", got)
	}

	if got := len(occ['1']); got != 3 {
		t.Error("Invalid for 1, got", got)
	}
}

func parse(t *testing.T) []string {
	d, err := os.ReadFile("data.txt")
	if err != nil {
		t.Fatal("Error in reading file",err)
		return nil
	}
	return strings.Split(string(d), "\r\n")
}

func solveP1(d []string) int {
	out := ""
	for i := 0; i < len(d[0]); i++ {
		occurences := map[byte]int{}
		for _, line := range d {
			c := line[i]
			occurences[c]++
		}
		if occurences['0'] > occurences['1'] {
			out += "0"
		} else {
			out += "1"
		}
	}
	gamma := parseHex(out)
	epsilon := parseHex(negate(out))
	return gamma*epsilon
}

func negate(in string) string {
	out := ""
	for _, c := range in {
		switch c {
		case '0': out += "1"
		case '1': out += "0"
		}
	}
	return out
}

func solveP2(d []string) int {
	og := parseOxygenGenerator(d)
	co2 := parseCo2(d)

	oxygenGen := parseHex(og)
	co2ScrubberRating := parseHex(co2)
	return oxygenGen*co2ScrubberRating
}

func parseOxygenGenerator(d []string) string {
	return parseThing(d, find1)

}

func parseCo2(d []string) string {
	return parseThing(d, find0)
}

func parseThing(d []string, finderFn func(int, []string)[]string) string {
	inputStrings := d
	for i := 0; i < len(d[0]); i++ {
		if len(inputStrings) == 1 {
			return inputStrings[0]
		}
		inputStrings = finderFn(i, inputStrings)
	}
	if len(inputStrings) == 1 {
		return inputStrings[0]
	}
	return ""
}

func find1(idx int, inputStrings []string) []string {
	occurences := group(idx, inputStrings)
	if len(occurences['1']) >= len(occurences['0']) {
		return occurences['1']
	} 		
	return occurences['0']
}

func find0(idx int, inputStrings []string) []string {
	occurences := group(idx, inputStrings)
	if len(occurences['0']) <= len(occurences['1']) {
		return occurences['0']
	} 		
	return occurences['1']
}

func group(idx int, inputStrings []string) map[byte][]string {
	occurences := map[byte][]string{}
	for _, line := range inputStrings {
		c := line[idx]
		x := occurences[c]
		occurences[c] = append(x, line)
	}
	return occurences
}

func parseHex(in string) int {
	v, _ := strconv.ParseInt(in, 2, 64)
	return int(v)
}