package d5

import (
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExample(t *testing.T) {
	input := `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`
	lines := strings.Split(input, "\n")
	t.Run("p1", func(t *testing.T) {
		cranes, rules := parse(lines)
		assert.Equal(t, "CMZ", p1(cranes, rules))
	})

	t.Run("p2", func(t *testing.T) {
		cranes, rules := parse(lines)
		assert.Equal(t, "MCD", p2(cranes, rules))
	})
}

func TestActual(t *testing.T) {
	raw, err := os.ReadFile("data.txt")
	require.NoError(t, err)
	lines := strings.Split(string(raw), "\r\n")

	t.Run("p1", func(t *testing.T) {
		cranes, rules := parse(lines)
		assert.Equal(t, "QNNTGTPFN", p1(cranes, rules))
	})

	t.Run("p2", func(t *testing.T) {
		cranes, rules := parse(lines)
		assert.Equal(t, "GGNPJBTTR", p2(cranes, rules))
	})
}

func TestStack(t *testing.T) {
	s := stack{}
	s.push("1")
	s.push("2")
	s.push("3")

	fn := func(exp string) {
		v, ok := s.pop()
		assert.True(t, ok)
		assert.Equal(t, exp, v)
	}

	fn("3")
	fn("2")
	fn("1")

	_, ok := s.pop()
	assert.False(t, ok)
}

type stack []string

func reverse(s []string) {
	ln := len(s)
	for i := 0; i < ln/2; i++ {
		s[i],s[ln-1-i] = s[ln-1-i],s[i]
	}
}

func (s *stack) push(v string) {
	*s = append(*s, v)
}

func (s *stack) pop() (string, bool) {
	if len(*s) == 0 {
		return "", false
	}
	out := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return out, true
}


type craneSet []stack
type rule struct {
	howMany     int
	source      int
	destination int
}

// all cranes are 4 bytes wide, but the last one. Align it
func alignedLineToProcess(line string) string {
	return line + " "
}

func parse(lines []string) (craneSet, []rule) {
	const craneWidth = 4

	cranes := craneSet{}
	rules := []rule{}
	parsingCranes := true

	howManyCranes := len(alignedLineToProcess(lines[0])) / craneWidth
	for i := 0; i < howManyCranes; i++ {
		cranes = append(cranes, stack{})
	}

	for _, line := range lines {
		if parsingCranes && line == "" {
			parsingCranes = false

			// my parser is bad, reverse the ordering of boxes
			for _, c := range cranes {
				reverse(c)
			}

			continue
		}

		if parsingCranes {
			for i := 0; i < howManyCranes; i++ {
				startIds := i * craneWidth
				endIds := (i + 1) * craneWidth

				lineToProcess := alignedLineToProcess(line)[startIds:endIds]
				if strings.HasPrefix(strings.TrimSpace(lineToProcess), "[") {
					// parsing from the top, need to reverse it later
					cranes[i].push(string(lineToProcess[1])) // skip '[',']'
				}
			}
		} else { // parsing rules
			re, _ := regexp.Compile(`move (\d+) from (\d+) to (\d+)`)
			results := re.FindAllStringSubmatch(line, -1)[0]
			howMany, _ := strconv.Atoi(results[1])
			from, _ := strconv.Atoi(results[2])
			to, _ := strconv.Atoi(results[3])
			rules = append(rules, rule{
				howMany:     howMany,
				source:      from,
				destination: to,
			})
		}
	}
	return cranes, rules
}

func buildOutput(cranes craneSet) string {
	out := ""
	for _, c := range cranes {
		v, ok := c.pop()
		if ok {
			out += v
		}
	}
	return out
}

func p1(cranes craneSet, rules []rule) string {
	for _, r := range rules {
		for i := 0; i < r.howMany; i++ {
			v, _ := cranes[r.source-1].pop()
			cranes[r.destination-1].push(v)
		}
	}

	return buildOutput(cranes)
}

func p2(cranes craneSet, rules []rule) string {
	for _, r := range rules {
		var vs []string
		for i := 0; i < r.howMany; i++ {
			v, _ := cranes[r.source-1].pop()
			vs = append(vs, v)
		}
		reverse(vs)
		for _, v := range vs {
			cranes[r.destination-1].push(v)
		}
	}

	return buildOutput(cranes)
}
