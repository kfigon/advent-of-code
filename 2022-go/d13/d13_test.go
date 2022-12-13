package d13

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	data := parse(strings.Split(example, "\n"))
	t.Run("p1", func(t *testing.T) {
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
		{"[[8,7,6]]", []any{[]any{8,7,6}}},
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
	for _, p := range pairs {
		if ordered(p.a, p.b) {
			out++
		}
	}
	return out
}

func ordered(a []any, b []any) bool {
	return false
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