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
		ex := "foobar"
		got, err := os.ReadFile("d7.txt")

		require.NoError(t, err)
		g, err := parse(string(got))

		require.NoError(t, err)
		assert.Equal(t, ex, p1(g))
	})
}

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
	}
	return g, nil
}

func p1(g graph) string {
	// find starting node - node without any dependencies
	dependingNodes := map[byte]bool{}
	startingNodes := map[byte]bool{}
	for k, v := range g {
		startingNodes[k] = true
		for _, e := range v {
			dependingNodes[e] = true
		}
	}

	var startingNode byte
	for v := range startingNodes {
		if _, ok := dependingNodes[v]; !ok {
			startingNode = v
			break
		}
	}

	var out strings.Builder

	var traverse func(byte)
	traverse = func(n byte) {
		out.WriteByte(n)
		children := g[n]
		slices.Sort(children)

		for _, child := range children {
			traverse(child)
		}
	}
	traverse(startingNode)

	return out.String()
}
