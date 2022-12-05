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
	cranes, rules := parse(strings.Split(input, "\n"))
	t.Run("p1", func(t *testing.T) {
		assert.Equal(t, "CMZ", p1(cranes, rules))
	})

	t.Run("p2", func(t *testing.T) {
		assert.Equal(t, "1234", p2(cranes, rules))
	})
}

func TestActual(t *testing.T) {
	raw, err := os.ReadFile("data.txt")
	require.NoError(t, err)

	cranes, rules := parse(strings.Split(string(raw), "\r\n"))
	t.Run("p1", func(t *testing.T) {
		assert.Equal(t, "sdaf", p1(cranes, rules))
	})

	t.Run("p2", func(t *testing.T) {
		assert.Equal(t, "adsf", p2(cranes, rules))
	})
}

func TestStack(t *testing.T) {
	s := stack{}
	s.push("1")
	s.push("2")
	s.push("3")

	v, ok := s.pop()
	assert.True(t, ok)
	assert.Equal(t, "3", v)
}

type stack []string

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

func parse(lines []string) (craneSet, []rule) {
	cranes := craneSet{}
	rules := []rule{}
	parsingCranes := true

	howManyCranes := (len(lines[0]) + 1) / 4
	for i := 0; i < howManyCranes; i++ {
		cranes = append(cranes, stack{})
	}

	for _, line := range lines {
		if line == "" {
			parsingCranes = false
			continue
		}

		if parsingCranes {
			for i := 0; i < howManyCranes; i++ {
				startIds := i * 4
				endIds := (i + 1) * 4

				lineToProcess := (line + " ")[startIds:endIds]
				if strings.HasPrefix(strings.TrimSpace(lineToProcess), "[") {
					cranes[i].push(string(lineToProcess[1])) // skip '[',']'
				}
			}
		} else {
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

func p1(cranes craneSet, rules []rule) string {
	for _, r := range rules {
		for i := 0; i < r.howMany; i++ {
			v, _ := cranes[r.source-1].pop()
			cranes[r.destination-1].push(v)
		}
	}

	out := ""
	for _, c := range cranes {
		v, ok := c.pop()
		if ok {
			out += v
		}
	}
	return out
}

func p2(cranes craneSet, rules []rule) string {
	return ""
}
