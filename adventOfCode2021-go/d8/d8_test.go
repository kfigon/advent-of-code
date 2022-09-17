package d8

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)
const exampleInput string = `be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce`


func TestExampleP1(t *testing.T) {
	exp := 26
	got := solveP1(parse(exampleInput, "\n"))

	if got != exp {
		t.Error("Exp", exp, "got",got)
	}
}

func TestPart1(t *testing.T) {
	d, err := os.ReadFile("data.txt")
	if err != nil {
		t.Fatal(err)
		return
	}
	exp := 239
	got := solveP1(parse(string(d), "\r\n"))

	if got != exp {
		t.Error("Exp", exp, "got",got)
	}
}

func TestParseP2(t *testing.T) {
	testCases := []struct {
		input	string
		expected string
	}{
		{
			input: "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf",
			expected: "5353",
		},
		{
			input: "be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe",
			expected: "8394",
		},
		{
			input: "edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc",
			expected: "9781",
		},
		{
			input: "fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg",
			expected: "1197",
		},
		{
			input: "fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb",
			expected: "9361",
		},
		{
			input: "aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea",
			expected: "4873",
		},
		{
			input: "fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb",
			expected: "8418",
		},
		{
			input: "dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe",
			expected: "4548",
		},
		{
			input: "bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef",
			expected: "1625",
		},
		{
			input: "egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb",
			expected: "8717",
		},
		{
			input: "gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce",
			expected: "4315",
		},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := parseP2(parse(tC.input, "\n")[0])
			if got != tC.expected {
				t.Error("Exp",tC.expected, "got", got)
			}
		})
	}
}

func TestExampleP2(t *testing.T) {
	exp := 61229
	got := solveP2(parse(exampleInput, "\n"))

	if got != exp {
		t.Error("Exp", exp, "got",got)
	}
}

func TestP2(t *testing.T) {
	d,err := os.ReadFile("data.txt")
	if err != nil {
		t.Fatal(err)
		return
	}

	exp := 946346
	got := solveP2(parse(string(d), "\n"))
	if got != exp {
		t.Error("Exp", exp, "got",got)
	}
}

func solveP2(input []data) int {
	out := 0
	for _, d := range input {
		v := parseP2(d)
		x, _ := strconv.Atoi(v)
		out += x
	}
	return out
}

type data struct {
	notes []string
	outputs []string
}

func parse(in string, sep string) []data {
	lines := strings.Split(in, sep)
	out := []data{}

	for _, line := range lines {
		parts := strings.Split(line, " | ")
		d := data{
			notes: strings.Fields(parts[0]),
			outputs: strings.Fields(parts[1]),
		}
		out = append(out, d)
	}
	return out
}

func solveP1(input []data) int {
	out := 0
	for _, line := range input {
		for _, v := range line.outputs {
			if _,ok := determineCandidate(v); ok {
				out++
			}
		}
	}
	return out
}

func determineCandidate(in string) (int,bool) {
	candidates := map[int][]int {
		2: {1},
		4: {4},
		6: {6,0,9},
		5: {2,3,5},
		3: {7},
		7: {8},
	}
	c := candidates[len(in)]
	if len(c) != 1 {
		return -1, false
	}
	return c[0],true
}

func parseP2(input data) string {

	mappings := map[string]int{}
	fourVal := ""
	oneVal := ""
	sevenVal := ""

	flatData := append(input.outputs, input.notes...)
	
	twoThreeOrFive := map[string]bool{}
	zeroSixOrNine := map[string]bool{}

	for _, d := range flatData {
		switch len(d) {
		case 2: {
			mappings[d] = 1
			oneVal=d
		}
		case 4:{
			mappings[d]=4
			fourVal=d
		}
		case 3:{
			mappings[d]=7
			sevenVal=d
		}
		case 7:{
			mappings[d]=8
		}
		case 5: twoThreeOrFive[d] = true
		case 6: zeroSixOrNine[d] = true
		}
	}

	for k  := range zeroSixOrNine {
		commonWith4 := commonLettersNumber(fourVal, k)
		commonWith1 := commonLettersNumber(oneVal, k)
		commonWith7 := commonLettersNumber(sevenVal, k)

		switch {
		case commonWith4 == 4 && commonWith7 == 3: mappings[k]=9
		case commonWith4 == 4: mappings[k]=9
		case commonWith4 == 3 && (commonWith1 == 2 || commonWith7 == 3): mappings[k]=0
		case commonWith4 == 3: mappings[k]=6
		case commonWith1 == 1: mappings[k]=6
		case commonWith7 == 2: mappings[k]=6
		}
	}
	
	for k := range twoThreeOrFive {
		commonWith4 := commonLettersNumber(fourVal, k)
		commonWith1 := commonLettersNumber(oneVal, k)
		commonWith7 := commonLettersNumber(sevenVal, k)

		switch {
		case commonWith4 == 2: mappings[k]=2
		case commonWith1 == 2: mappings[k]=3
		case commonWith7 == 3: mappings[k]=3
		case commonWith4 == 3 && (commonWith1 == 1 || commonWith7 == 2): mappings[k]=5
		}
	}

	out := ""
	for _, v := range input.outputs {
		d := mappings[v]
		out += strconv.Itoa(d)
	}
	return out
}

func commonLettersNumber(a,b string) int {
	aSet := map[rune]bool{}
	bSet := map[rune]bool{}
	for _, v := range a {
		aSet[v] = true
	}
	for _, v := range b {
		bSet[v]=true
	}

	common := 0
	for k := range aSet {
		if bSet[k] {
			common++
		}
	}
	return common
}