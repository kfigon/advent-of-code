package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	data, err := os.ReadFile("d1.txt")
	require.NoError(t, err)

	input := string(data)

	t.Run("d1p1", func(t *testing.T) {
		assert.Equal(t, 1390, solveD1(input, d1Provider))
	})

	t.Run("d1p2", func(t *testing.T) {
		assert.Equal(t, 1232, solveD1(input, d2Provider))
	})
}
