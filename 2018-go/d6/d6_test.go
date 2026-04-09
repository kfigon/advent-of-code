package d6

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var example = `1, 1
1, 6
8, 3
3, 4
5, 5
8, 9`

func TestP1(t *testing.T) {
	t.Run("ex", func(t *testing.T) {
		d, err := parse(example)
		require.NoError(t, err)

		assert.Equal(t, 17, p1(d))
	})
}

type coord struct {
	x int
	y int
}

func parse(data string) ([]coord, error) {
	out := []coord{}
	for line := range strings.Lines(data) {
		line = strings.TrimSpace(line)
		got := strings.Split(line, ", ")
		if len(got) != 2 {
			return nil, fmt.Errorf("separator not found for line: %q", line)
		}
		a, err := strconv.Atoi(got[0])
		if err != nil {
			return nil, fmt.Errorf("parsing error for %q: %w", line, err)
		}
		b, err := strconv.Atoi(got[1])
		if err != nil {
			return nil, fmt.Errorf("parsing error for %q: %w", line, err)
		}
		out = append(out, coord{a, b})
	}
	return out, nil
}

func p1(c []coord) int {
	return 0
}
