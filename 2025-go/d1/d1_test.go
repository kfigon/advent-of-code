package d1

import (
	"aoc2025/util"
	"fmt"
	"io"
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
		sum += num
		base := numberOfHundreds(num)
		// todo: multiple rotations

		if sum > 99 {
			res++
			sum -= (sum / 100) * 100
		} else if sum < 0 {
			sum += ((-1 * sum) / 100) * 100
			res++
		} else if sum == 0 {
			res++
		}
	}

	return res, nil
}

func numberOfHundreds(v int) int {
	if v < 0 {
		return ((-1 * v) / 100)
	}
	return v / 100
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
