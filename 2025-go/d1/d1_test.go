package d1

import (
	"aoc2025/util"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const example = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

func P1(r io.Reader) (int, error) {
	sum := 50
	res := 0
	for line := range util.Lines(r) {
		num, err := parse(line)
		if err != nil {
			return 0, err
		}
		sum += num

		if sum%100 == 0 {
			res++
		}
	}

	return res, nil
}

func parse(line string) (int, error) {
	num, err := strconv.Atoi(line[1:])
	if err != nil {
		return 0, fmt.Errorf("error reading line: %w", err)
	}

	switch line[0] {
	case 'L':
		return -num, nil
	case 'R':
		return num, nil
	default:
		return 0, fmt.Errorf("invalid sign char on line %q", line)
	}
}

func P2(r io.Reader) (int, error) {
	sum := 50
	res := 0

	for line := range util.Lines(r) {
		num, err := parse(line)
		if err != nil {
			return 0, err
		}
		newSum, movesOverZero := calcP2(sum, num)
		sum = newSum
		res += movesOverZero
	}

	return res, nil
}

func calcP2(position, num int) (newPosition int, movesOverZero int) {
	var numberOfFullRotations int = num / 100
	if num < 0 {
		numberOfFullRotations = -1 * numberOfFullRotations
	}
	// hundreds changed
	// move from positive to negative
	// move from negative to positive
	moved := position/100 != (position+num)/100

	newPosition = position + num
	if moved {
		movesOverZero = 1 + numberOfFullRotations
	} else {
		movesOverZero = numberOfFullRotations
	}

	return
}

func TestP2Logic(t *testing.T) {
	testCases := []struct {
		pos           int
		num           int
		newPos        int
		movesOverZero int
	}{
		{50, 50, 100, 1},
		{50, 150, 200, 2},
		{50, 160, 210, 2},
		{50, 130, 180, 1},
		{50, 3, 53, 0},
		{50, -3, 47, 0},
		{50, -180, -130, 2},
		{50, -130, -80, 1},
		{50, -50, 0, 1},
		{50, 240, 290, 2},
		{50, -240, -190, 2},
		{-50, 80, 30, 1},
		{-50, 130, 80, 1},
		{-50, 3, -47, 0},
		{-50, -3, -53, 0},
		{-50, -180, -230, 2},
		{-50, 240, 190, 2},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprintf("%d+%d=(%d,%d)", tC.pos, tC.num, tC.newPos, tC.movesOverZero), func(t *testing.T) {
			newSum, movesOverZero := calcP2(tC.pos, tC.num)
			assert.Equal(t, tC.newPos, newSum)
			assert.Equal(t, tC.movesOverZero, movesOverZero)
		})
	}
}

func TestP1(t *testing.T) {
	testCases := []struct {
		desc   string
		reader io.Reader
		exp    int
	}{
		{
			desc:   "p1 example",
			reader: strings.NewReader(example),
			exp:    3,
		},
		{
			desc:   "p1 file",
			reader: util.MustReadFile("input.txt"),
			exp:    1076,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := P1(tC.reader)
			assert.NoError(t, err)
			assert.Equal(t, tC.exp, got)
		})
	}
}

func TestP2(t *testing.T) {
	testCases := []struct {
		desc   string
		reader io.Reader
		exp    int
	}{
		{
			desc:   "example",
			reader: strings.NewReader(example),
			exp:    6,
		},
		{
			desc:   "file",
			reader: util.MustReadFile("input.txt"),
			exp:    3723,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := P2(tC.reader)
			assert.NoError(t, err)
			assert.Equal(t, tC.exp, got)
		})
	}
}
