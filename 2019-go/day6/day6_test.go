package day6

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)


func TestOrbits(t *testing.T) {
	testCases := []struct {
		node string
		exp int		
	}{
		{"D",3},
		{"L",7},
		{"COM",0},
		{"I",4},
		{"F",5},
		{"E",4},
	}
	testInput := testData()

	for _, tc := range testCases {
		t.Run(tc.node, func(t *testing.T) {
			got := buildOrbits(testInput,"\n").calcOrbits(tc.node)
			assert.Equal(t, tc.exp, got)
		})
	}
}

func testData() string {
	return `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L`
}

func testData2() string {
	return testData() + `
K)YOU
I)SAN`
}

func TestAllOrbits(t *testing.T) {
	orbits := buildOrbits(testData(), "\n")
	got := orbits.calcAllOrbits()
	assert.Equal(t, 42, got)
}

func TestAllOrbitsPart1(t *testing.T) {
	orbits := buildOrbits(readFile(t), "\r\n")
	got := orbits.calcAllOrbits()
	assert.Equal(t, 261306, got)
}

func readFile(t *testing.T) string {
	file, err := os.Open("data.txt")
	require.NoError(t,err)
	defer file.Close()

	content, err := io.ReadAll(file)
	require.NoError(t,err)
	return string(content)
}

func TestPart2Example(t *testing.T) {
	orbits := buildOrbits(testData2(), "\n")
	got := orbits.findPath("YOU", "SAN")
	assert.Equal(t, 4, got)
}

func TestPart2File(t *testing.T) {
	orbits := buildOrbits(readFile(t), "\r\n")
	got := orbits.findPath("YOU", "SAN")
	assert.Equal(t, 382, got)
}
