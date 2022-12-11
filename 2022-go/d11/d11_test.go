package d11

import (
	"os"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExample(t *testing.T) {
	t.Run("p1", func(t *testing.T) {
		rules := parse(strings.Split(example,"\n"))
		assert.Equal(t, 10605, p1(rules))
	})

	t.Run("p2", func(t *testing.T) {
		rules := parse(strings.Split(example,"\n"))
		assert.Equal(t, 2713310158, p2(rules))
	})
}

func TestFile(t *testing.T) {
	raw, err := os.ReadFile("data.txt")
	require.NoError(t, err)
	
	t.Run("p1", func(t *testing.T) {
		rules := parse(strings.Split(string(raw),"\r\n"))
		assert.Equal(t, 120384, p1(rules))
	})

	t.Run("p2", func(t *testing.T) {
		rules := parse(strings.Split(string(raw),"\r\n"))
		assert.Equal(t, 32059801242, p2(rules))
	})
}

type operation struct {
	op string
	a string
	b string
}

func (o operation) eval(oldV int) int {
	parseArg := func(arg string) int {
		v, err := strconv.Atoi(arg)
		if err != nil {
			return oldV
		}
		return v
	}

	switch o.op {
	case "+": return parseArg(o.a) + parseArg(o.b)
	case "*": return parseArg(o.a) * parseArg(o.b)
	case "/": return parseArg(o.a) / parseArg(o.b)
	case "-": return parseArg(o.a) - parseArg(o.b)
	}

	return 0
}

type monkey struct {
	items []int
	op operation
	testDivisibleBy int
	ruleTrue int
	ruleFalse int
	commonDivisor int
}

func parse(lines []string) []*monkey {
	out := []*monkey{}
	current := &monkey{}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			out = append(out, current)
			current = &monkey{}
			continue
		} else if strings.HasPrefix(line, "Monkey") {
			continue
		} else if strings.HasPrefix(line, "Starting") {
			intermediate := strings.TrimPrefix(line, "Starting items: ")
			nums := strings.Split(intermediate, ", ")

			for _, n := range nums {
				v, _ := strconv.Atoi(n)
				current.items = append((*current).items, v)
			}
		} else if strings.HasPrefix(line, "Operation") {
			intermediate := strings.TrimPrefix(line, "Operation: new = ")
			operator := ""
			if strings.Contains(intermediate, "*") {
				operator = "*"
			} else if strings.Contains(intermediate, "+") {
				operator = "+"
			} else if strings.Contains(intermediate, "-") {
				operator = "-"
			} else if strings.Contains(intermediate, "/") {
				operator = "/"
			}

			operands := strings.Split(intermediate, " " +operator+" ")
			current.op = operation{
				op: operator,
				a: operands[0],
				b: operands[1],
			}
		} else if strings.HasPrefix(line, "Test") {
			intermediate := strings.TrimPrefix(line, "Test: divisible by ")
			v, _ := strconv.Atoi(intermediate)
			current.testDivisibleBy = v
		} else if strings.HasPrefix(line, "If true") {
			intermediate := strings.TrimPrefix(line, "If true: throw to monkey ")
			v, _ := strconv.Atoi(intermediate)
			current.ruleTrue = v
		} else if strings.HasPrefix(line, "If false") {
			intermediate := strings.TrimPrefix(line, "If false: throw to monkey ")
			v, _ := strconv.Atoi(intermediate)
			current.ruleFalse = v
		}
	}

	out = append(out, current)
	return out
}

func round(monkeys []*monkey, reduceWory bool) map[int]int {
	interactions := map[int]int{}
	for i := 0; i < len(monkeys); i++ {
		m := monkeys[i]
		for _, item := range m.items {
			interactions[i]++
			v := m.op.eval(item)
			if reduceWory {
				v = v/3
			} else {
				v = v % m.commonDivisor
			}

			if v % m.testDivisibleBy == 0 {
				monkeys[m.ruleTrue].items = append(monkeys[m.ruleTrue].items, v)
			} else {
				monkeys[m.ruleFalse].items = append(monkeys[m.ruleFalse].items, v)
			}
		}
		m.items = []int{}
	}
	return interactions
}

func solve(monkeys []*monkey, reduceWoryLevel bool, rounds int) int {
	interactions := map[int]int{}
	for i := 1; i <= rounds; i++ {
		results := round(monkeys, reduceWoryLevel)
		for monkey,inters := range results {
			interactions[monkey]+=inters
		}
	}
	interactionsTab := []int{}
	for _, v := range interactions {
		interactionsTab = append(interactionsTab, v)
	}

	sort.Ints(interactionsTab)
	ln := len(interactionsTab)
	return interactionsTab[ln-1] * interactionsTab[ln-2]
}

func p1(monkeys []*monkey) int {
	return solve(monkeys, true, 20)
}

func p2(monkeys []*monkey) int {
	// https://en.wikipedia.org/wiki/Modular_arithmetic
	commonDiv := 1
	for _, m := range monkeys {
		commonDiv *= m.testDivisibleBy
	}
	for _, m := range monkeys {
		m.commonDivisor = commonDiv
	}
	return solve(monkeys, false, 10000)
}

const example = `Monkey 0:
Starting items: 79, 98
Operation: new = old * 19
Test: divisible by 23
  If true: throw to monkey 2
  If false: throw to monkey 3

Monkey 1:
Starting items: 54, 65, 75, 74
Operation: new = old + 6
Test: divisible by 19
  If true: throw to monkey 2
  If false: throw to monkey 0

Monkey 2:
Starting items: 79, 60, 97
Operation: new = old * old
Test: divisible by 13
  If true: throw to monkey 1
  If false: throw to monkey 3

Monkey 3:
Starting items: 74
Operation: new = old + 3
Test: divisible by 17
  If true: throw to monkey 0
  If false: throw to monkey 1`