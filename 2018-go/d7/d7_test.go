package d6

import (
	"fmt"
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
}

type graph map[rune][]rune

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
	}
	return g, nil
}

func p1(g graph) string {
	return ""
}
