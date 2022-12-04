package d5

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExample(t *testing.T) {
	input := ``
	data := parse(strings.Split(input, "\n"))
	t.Run("p1", func(t *testing.T) {
		assert.Equal(t, 2, p1(data))
	})

	t.Run("p2", func(t *testing.T) {
		assert.Equal(t, 4, p2(data))
	})
}

func TestActual(t *testing.T) {
	raw, err := os.ReadFile("data.txt")
	require.NoError(t, err)

	data := parse(strings.Split(string(raw), "\r\n"))
	t.Run("p1", func(t *testing.T) {
		assert.Equal(t, 562, p1(data))
	})

	t.Run("p2", func(t *testing.T) {
		assert.Equal(t, 924, p2(data))
	})
}

func parse(lines []string) {

}

func p1() int{
	return 0
}

func p2() int{
	return 0
}