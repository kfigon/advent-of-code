package d2

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_example(t *testing.T) {
	lines := []string{
		"A Y",
		"B X",
		"C Z",
	}
	t.Run("p1", func(t *testing.T) {
		assert.Equal(t, 15, solve(lines, p1))
	})

	t.Run("p2", func(t *testing.T) {
		assert.Equal(t, 12, solve(lines, p2))
	})
}

func Test_files(t *testing.T) {
	d, err := os.ReadFile("data.txt")
	require.NoError(t, err)
	lines := strings.Split(string(d), "\r\n")

	t.Run("p1", func(t *testing.T) {
		assert.Equal(t, 14827, solve(lines, p1))
	})

	t.Run("p2", func(t *testing.T) {
		assert.Equal(t, 13889, solve(lines, p2))
	})
}

type result int

const (
	draw result = iota
	lose
	win
)

func (r result) score() int {
	switch r {
	case draw:
		return 3
	case lose:
		return 0
	case win:
		return 6
	}
	return 0
}

type figure int

const (
	rock figure = iota
	paper
	scissors
)

func (f figure) score() int {
	switch f {
	case paper:
		return 2
	case scissors:
		return 3
	case rock:
		return 1
	}
	return 0
}

func (f figure) game(opponent figure) result {
	mapping := map[figure]map[figure]result{
		rock: {
			rock:     draw,
			paper:    lose,
			scissors: win,
		},
		paper: {
			rock:     win,
			paper:    draw,
			scissors: lose,
		},
		scissors: {
			rock:     lose,
			paper:    win,
			scissors: draw,
		},
	}
	return mapping[f][opponent]
}

func solve(lines []string, mapper figureMapper) int {
	out := 0
	for _, line := range lines {
		vs := strings.Split(line, " ")
		opponent, myFigure := mapper(vs[0], vs[1])
		out += myFigure.score() + myFigure.game(opponent).score()
	}
	return out
}

type figureMapper func(string, string) (figure, figure)

func opponentMapping(code string) figure {
	switch code {
	case "A":
		return rock
	case "B":
		return paper
	case "C":
		return scissors
	}
	return rock
}

func p1(opponentsFigureCode, myFigureCode string) (figure, figure) {
	my := func() figure {
		switch myFigureCode {
		case "X":
			return rock
		case "Y":
			return paper
		case "Z":
			return scissors
		}
		return rock
	}
	return opponentMapping(opponentsFigureCode), my()
}

func p2(opponentsFigureCode, myFigureCode string) (figure, figure) {
	opponentsFigure := opponentMapping(opponentsFigureCode)
	my := func() figure {
		switch myFigureCode {
		case "X": // i need to lose
			switch opponentsFigure {
			case rock:
				return scissors
			case paper:
				return rock
			case scissors:
				return paper
			default:
				return rock
			}
		case "Y":
			return opponentsFigure
		case "Z": // i need to win
			switch opponentsFigure {
			case rock:
				return paper
			case paper:
				return scissors
			case scissors:
				return rock
			default:
				return rock
			}
		default:
			return rock
		}
	}

	return opponentsFigure, my()
}
