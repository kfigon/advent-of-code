package d6

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const example = `Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.`

func TestP1(t *testing.T) {
	t.Run("ex", func(t *testing.T) {
		ex := "CABDFE"
		g, err := parse(example)
		require.NoError(t, err)
		assert.Equal(t, ex, p1(g))
	})

	t.Run("real", func(t *testing.T) {
		ex := "LAPFCRGHVZOTKWENBXIMSUDJQY"
		got, err := os.ReadFile("d7.txt")

		require.NoError(t, err)
		g, err := parse(string(got))

		require.NoError(t, err)
		result := p1(g)

		assert.Equal(t, ex, result)
	})
}

func TestP2(t *testing.T) {
	t.Run("ex", func(t *testing.T) {
		ex := 15
		g, err := parse(example)
		require.NoError(t, err)
		assert.Equal(t, ex, p2(g))
	})

	t.Run("real", func(t *testing.T) {
		ex := 123
		got, err := os.ReadFile("d7.txt")

		require.NoError(t, err)
		g, err := parse(string(got))

		require.NoError(t, err)
		assert.Equal(t, ex, p2(g))
	})
}

// backward graph: key - node. val - required nodes for the key (dependencies)
// forward graph: key - node. val - possible ways from key
// we use forward
type graph map[byte][]byte

func (g graph) String() string {
	var b strings.Builder
	for k, v := range g {
		b.WriteString(fmt.Sprintf("%c -> %s\n", k, v))
	}
	return b.String()
}

func parse(in string) (graph, error) {
	g := graph{}
	for line := range strings.Lines(in) {
		got := strings.Fields(line)
		if len(got) != 10 {
			return nil, fmt.Errorf("invalid field: %v", got)
		}

		start := got[1]
		end := got[7]
		if len(start) != 1 || len(end) != 1 {
			return nil, fmt.Errorf("invalid start or end: %v, %v", start, end)
		}
		startChar := start[0]
		endChar := end[0]

		row := g[startChar]
		row = append(row, endChar)
		g[startChar] = row

		if _, ok := g[endChar]; !ok {
			g[endChar] = nil
		}
	}
	return g, nil
}

func p1(g graph) string {
	// kahn's algorithm

	// how many nodes connect to this node
	inDegree := map[byte]int{}
	for _, directions := range g {
		for _, availableNode := range directions {
			inDegree[availableNode]++
		}
	}

	queue := []byte{}
	for k := range g {
		if inDegree[k] == 0 {
			queue = append(queue, k)
		}
	}

	var out strings.Builder

	for len(queue) > 0 {
		slices.Sort(queue)
		top := queue[0]
		queue = queue[1:]

		out.WriteByte(top)
		for _, v := range g[top] {
			// if degree comes to 0, enqueue next
			if deg := inDegree[v]; deg == 1 {
				queue = append(queue, v)
			}
			inDegree[v]--
		}
		// delete(g, top) // no cycle risk, so we dont need this
	}

	return out.String()
}

func p2(g graph) int {
	return 0
}
