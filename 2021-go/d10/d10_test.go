package d10

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"testing"
)

const exampleData string = `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`

const fileSeparator = "\r\n"

func TestExampleP1(t *testing.T) {
	exp := 26397
	if got := solveP1(parse(exampleData, "\n")); got != exp {
		t.Error("Invalid p1, got", got, "exp",exp)
	}
}
func TestValidLines(t *testing.T) {
	testCases := []string {
		"([])", "{()()()}", "<([{}])>", "[<>({}){}[([])<>]]", "(((((((((())))))))))",
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprintf("%v",i), func(t *testing.T) {
			if !isValid(tC) {
				t.Error(tC, "should be valid")
			}
		})
	}
}

func TestP1(t *testing.T) {
	f, err := os.ReadFile("data.txt")
	if err != nil {
		t.Fatal(err)
		return
	}

	exp := 240123
	if got := solveP1(parse(string(f), fileSeparator)); got != exp {
		t.Error("Invalid p1, got", got, "exp",exp)
	}
}

func TestInValidLines(t *testing.T) {
	testCases := []string {
		"(]", "{()()()>", "(((()))}", "<([]){()}[{}])",
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprintf("%v",i), func(t *testing.T) {
			if isValid(tC) {
				t.Error(tC, "should be invalid")
			}
		})
	}
}

func TestExampleP2Score(t *testing.T) {
	exp := 288957
	got := solveP2(parse(exampleData, "\n"))

	if got != exp {
		t.Error("Got", got, "exp", exp)
	}
}

func TestP2Score(t *testing.T) {
	f, err := os.ReadFile("data.txt")
	if err != nil {
		t.Fatal(err)
		return
	}

	exp := 3260812321
	got := solveP2(parse(string(f), fileSeparator))

	if got != exp {
		t.Error("Got", got, "exp", exp)
	}
}

func TestExampleP1Fixer(t *testing.T) {
	testCases := []struct {
		line	string
		expected rune
		found rune
	}{
		{"{([(<{}[<>[]}>{[]{[(<()>" ,']','}'},
		{"[[<[([]))<([[{}[[()]]]", ']', ')'},
		{"[{[{({}]{}}([{[{{{}}([]", ')', ']'},
		{"[<(<(<(<{}))><([]([]()", '>', ')'},
		{"<{([([[(<>()){}]>(<<{{", ']','>'},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			result := autocomplete(tC.line)
			shouldBe := result.badEnclosing.shouldBe
			found := result.badEnclosing.actual

			if shouldBe != tC.expected {
				t.Error("Got",shouldBe, "expected", tC.expected)
			}
			if found != tC.found {
				t.Error("Got (found) - ",found, "expected", tC.found)
			}
		})
	}
}

func TestExampleP2(t *testing.T) {
	testCases := []struct {
		input	string
		expectedFix string
	}{
			{"[({(<(())[]>[[{[]{<()<>>" , "}}]])})]"},
			{"[(()[<>])]({[<{<<[]>>(" , ")}>]})"},
			{"(((({<>}<{<{<>}{[]{[]{}" , "}}>}>))))"},
			{"{<[[]]>}<{[{[{[]{()[[[]" , "]]}}]}]}>"},
			{"<{([{{}}[<[[[<>{}]]]>[]]" , "])}>"},			
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			result := autocomplete(tC.input)
			if result.badEnclosing != nil {
				t.Error("Bad enclosing not expected")
			}

			if result.missingChars == nil {
				t.Error("Should be invalid")
			}
			if got := result.missingChars.data; got != tC.expectedFix {
				t.Error("Got", got, "exp", tC.expectedFix)
			}
		})
	}
}

func parse(in, sep string) []string {
	return strings.Split(in,sep)
}

type autocompleteResult struct {
	badEnclosing *badEnclosingError
	missingChars *missingCharsError
}

type badEnclosingError struct {
	shouldBe rune
	actual rune
}
type missingCharsError struct {
	data string
}

func autocomplete(line string) autocompleteResult {
	s := newStack()
	for _, char := range line {
		if isOpeningChar(char) {
			s.push(char)
		} else {
			lastOpening, ok := s.pop()
			if !ok {
				break
			} else if matchingSign(lastOpening) != char {
				return autocompleteResult{ badEnclosing: &badEnclosingError{
					shouldBe:matchingSign(lastOpening),
					actual: char,
				}}
			}
		}
	}

	if len(s.d) != 0 {
		out := ""
		for len(s.d) != 0 {
			if v, ok := s.pop(); ok {
				out += string(matchingSign(v))
			}
		}
		return autocompleteResult{ missingChars: &missingCharsError{ out }}
	}

	return autocompleteResult{}
}

func isValid(line string) bool {
	result := autocomplete(line)
	return result.badEnclosing == nil && result.missingChars == nil
}

func isOpeningChar(r rune) bool {
	return matchingSign(r) != '0'
}

func matchingSign(r rune) rune {
	switch r {
	case '(':return ')'
	case '{':return '}'
	case '<':return '>'
	case '[':return ']'
	default: return '0'
	}
}

func solveP1(lines []string) int {
	scores := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>':25137,
	}
	score := 0

	for _, line := range lines {
		result := autocomplete(line)
		if result.badEnclosing != nil {
			score += scores[result.badEnclosing.actual]
		}
	}
	return score
}

type stack struct {
	d []rune
}

func newStack() *stack {
	return &stack{
		d: []rune{},
	}
}

func (s *stack) push(r rune) {
	s.d = append(s.d, r)
}

func (s *stack) pop() (rune, bool) {
	if len(s.d) == 0 {
		return '0',false
	}
	last := s.d[len(s.d)-1]
	s.d = s.d[:len(s.d)-1]
	return last, true
}

func solveP2(lines []string) int {
	scores := map[rune]int{
		')':1,
		']':2,
		'}':3,
		'>':4,
	}

	score := []int{}
	for _, line := range lines {
		result := autocomplete(line)
		if result.missingChars == nil {
			continue
		}

		subScore := 0
		for _, v := range result.missingChars.data {
			subScore = subScore*5 + scores[v]
		}
		score = append(score, subScore)
	}

	sort.Ints(score)

	return score[len(score)/2]
}