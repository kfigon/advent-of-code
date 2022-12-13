package d13

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExample(t *testing.T) {
	data := parse(strings.Split(example, "\n"))
	t.Run("p1", func(t *testing.T) {
		assert.Equal(t, 13, p1(data))
	})
}

func TestFile(t *testing.T) {
	raw,err := os.ReadFile("data.txt")
	require.NoError(t, err)

	data := parse(strings.Split(string(raw), "\r\n"))
	t.Run("p1", func(t *testing.T) {
		assert.NotEqual(t, 711, p1(data)) // too low
		assert.Equal(t, 13, p1(data))
	})
}

func TestParser(t *testing.T) {
	testCases := []struct {
		input	string
		exp []any
	}{
		{"[1,1,3,1,1]", []any{1,1,3,1,1}},
		
		{"[[1],[2,3,4]]", []any{ []any{1}, []any{2,3,4} }},
		{"[[1],4]", []any{[]any{1},4}},

		{"[9]", []any{9}},
		{"[[8,7,6]]", []any{[]any{8,7,6}}},
		
		{"[[[[8],7],6]]", []any{ []any{ []any{ []any{8},7},6}}},
		{"[[[7,2],6,[8],[],[]]]", []any{ []any{ []any{7,2},6,[]any{8},[]any{},[]any{}}}},

		{"[[[]]]", []any{[]any{[]any{}}}},
		{"[1,[2,[3,[4,[5,6,7]]]],8,9]", []any{1,[]any{2,[]any{3,[]any{4, []any{5,6,7}}}}, 8,9}},

	}
	for _, tC := range testCases {
		t.Run(tC.input, func(t *testing.T) {
			got := parseSingle(tC.input)
			assert.Equal(t, tC.exp, got)
		})
	}
}

func TestCompare(t *testing.T) {
	testCases := []struct {
		input1 string
		input2 string
		exp bool
	}{
		{
			input1: "[1,1,3,1,1]",
			input2: "[1,1,5,1,1]",
			exp: true,
		},
		{
			input1: "[[1],[2,3,4]]",
			input2: "[[1],4]",
			exp: true,
		},
		{
			input1: "[9]",
			input2: "[[8,7,6]]",
			exp: false,
		},
		{
			input1: "[[4,4],4,4]",
			input2: "[[4,4],4,4,4]",
			exp: true,
		},
		{
			input1: "[7,7,7,7]",
			input2: "[7,7,7]",
			exp: false,
		},
		{
			input1: "[]",
			input2: "[3]",
			exp: true,
		},
		{
			input1: "[[[]]]",
			input2: "[[]]",
			exp: false,
		},
		{
			input1: "[1,[2,[3,[4,[5,6,7]]]],8,9]",
			input2: "[1,[2,[3,[4,[5,6,0]]]],8,9]",
			exp: false,
		},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprintf("Pair %v", i+1), func(t *testing.T) {
			a := parseSingle(tC.input1)
			b := parseSingle(tC.input2)
			assert.Equal(t, tC.exp, ordered(a,b))
		})
	}
}

type pair struct {
	a []any
	b []any
}

func parse(lines []string) []pair {
	out := []pair{}
	for i := 0; i < len(lines)-3; i+=3 {
		first := lines[i]
		second := lines[i+1]
		out = append(out, pair{parseSingle(first), parseSingle(second)})
	}
	return out
}

// json parser might be easier, but I like the challenge
func parseSingle(line string) []any {
	out := []any{}
	i := 0

	parseInt := func() int {
		buf := ""
		for i < len(line) {
			c := line[i]
			if c >= '0' && c<='9' {
				buf += string(c)
				i++
			} else {
				break
			}
		}
		v, _ := strconv.Atoi(buf)
		return v
	}

	var parseList func() []any
	parseList = func() []any {
		list := []any{}
		for i < len(line) {
			c := line[i]
			if c >= '0' && c <= '9' {
				list = append(list, parseInt())
			} else if c == ',' {
				i++
			} else if c == ']' {
				i++
				return list
			} else if c =='[' {
				i++
				list = append(list, parseList())
			} else {
				i++ // unknown
			}
		}
		return list
	}

	if line[0] == '[' {
		i++
		return parseList()
	}
	return out
}

func p1(pairs []pair) int {
	out := 0
	for i, p := range pairs {
		if ordered(p.a, p.b) {
			out += (i+1)
		}
	}
	return out
}

func ordered(a []any, b []any) bool {
	aI := 0
	bI := 0
	for aI <len(a) && bI < len(b) {
		aV := a[aI]
		bV := b[bI]
		
		aInt, aOk := aV.(int)
		bInt, bOk := bV.(int)
		if aOk && bOk {
			if aInt < bInt {
				return true
			} else if aInt > bInt {
				return false 
			}
			aI++
			bI++
			continue
		}

		aList, aOk := aV.([]any)
		bList, bOk := bV.([]any)
		if aOk && bOk {
			if !ordered(aList, bList) {
				return false
			}
			aI++
			bI++
		} else if !aOk && bOk {
			if !ordered([]any{aV}, bList) {
				return false
			}
			aI++
			bI++
		} else if aOk && !bOk {
			if !ordered(aList, []any{bV}) {
				return false
			}
			aI++
			bI++
		} else {
			// unreachable?
			aI++
			bI++
		}
	}

	if aI >= len(a) {
		return true
	} else if bI >= len(b) && aI < len(a) {
		return false
	}

	return true
}

const example = `[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]`