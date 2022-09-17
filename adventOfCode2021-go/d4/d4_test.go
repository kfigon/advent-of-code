package d4

import (
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestP1Example(t *testing.T) {
	input := `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
8  2 23  4 24
21  9 14 16  7
6 10  3 18  5
1 12 20 15 19

3 15  0  2 22
9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
2  0 12  3  7`

	nums, boards := parse(input, "\n")
	if got := solveP1(nums, boards); got != 4512 {
		t.Error("Invalid p1, got", got)
	}
}

func TestP1(t *testing.T) {
	d, err := os.ReadFile("data.txt")
	if err != nil {
		t.Fatal("Error in reading file",err)
		return
	}
	nums, boards := parse(string(d), "\r\n")
	if got := solveP1(nums, boards); got != 28082 {
		t.Error("Invalid p1, got", got)
	}
}

func TestP2Ex(t *testing.T) {
	input := `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
8  2 23  4 24
21  9 14 16  7
6 10  3 18  5
1 12 20 15 19

3 15  0  2 22
9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
2  0 12  3  7`

	nums, boards := parse(input, "\n")
	if got := solveP2(nums, boards); got != 1924 {
		t.Error("Invalid p2, got", got)
	}
}

func TestP2(t *testing.T) {
	d, err := os.ReadFile("data.txt")
	if err != nil {
		t.Fatal("Error in reading file",err)
		return
	}
	nums, boards := parse(string(d), "\r\n")
	if got := solveP2(nums, boards); got != 8224 {
		t.Error("Invalid p1, got", got)
	}
}

type pair struct {
	val int
	mark bool
}

type board [][]pair

func (b board) rows() int {
	return len(b)
}

func (b board) cols() int {
	return len(b[0])
}

func (b board) mark(num int) {
	for r := 0; r < b.rows(); r++ {
		for c := 0; c < b.cols(); c++ {
			if b[r][c].val == num {
				b[r][c].mark = true
				break
			}
		}
	}
}

func (b board) won() bool {
	horizontal := func(colIdx int) bool {
		for r := 0; r < b.rows(); r++ {
			if !b[r][colIdx].mark {
				return false
			}
		}
		return true
	}

	vertical := func(rowIdx int) bool {
		for c := 0; c < b.cols(); c++ {
			if !b[rowIdx][c].mark {
				return false
			}
		}
		return true
	}
	for r := 0; r < b.rows(); r++ {
		if vertical(r) {
			return true
		}
	}

	for c:= 0; c< b.cols(); c++ {
		if horizontal(c) {
			return true
		}
	}
	return false
}

func (b board) notMarked() []int {
	out := []int{}
	for r := 0; r < b.rows(); r++ {
		for c := 0; c < b.cols(); c++ {
			if !b[r][c].mark {
				out = append(out, b[r][c].val)
			}
		}
	}
	return out
}

func parse(in string, sep string) ([]int, []board) {
	lines := strings.Split(in, sep)
	nums := []int{}
	numsStr := strings.Split(lines[0], ",")
	for _, v := range numsStr {
		x, _ := strconv.Atoi(v)
		nums = append(nums, x)
	}

	boards := []board{}

	restLines := lines[1:]
	for i := 0; i < len(restLines); i+=6 {
		dataLines := restLines[i:i+6]
		board := board{}

		for _, line := range dataLines {
			if line == "" {
				continue
			}
			
			row := []pair{}
			splitted := strings.Fields(line)
			for _, v := range splitted {
				if v == "" {
					continue
				}
				x, _ := strconv.Atoi(v)
				row = append(row, pair{val: x})
			}
			board = append(board, row)
		}
		boards = append(boards, board)
	}
	return nums, boards
}

func solveP1(nums []int, boards []board) int {
	for _, n := range nums {
		for _, b := range boards {
			b.mark(n)
			if b.won() {
				return sum(b.notMarked()) * n
			}
		}
	}
	return -1
}

func solveP2(nums []int, boards []board) int {
	notWonBoardIdx := map[int]bool{}
	for b := 0; b < len(boards); b++ {
		notWonBoardIdx[b]=true
	}

	findLastBoard := func() int {
		for _, n := range nums {
			for bI := 0; bI < len(boards); bI++ {
				b := boards[bI]
				b.mark(n)
				if notWonBoardIdx[bI] && b.won() {
					delete(notWonBoardIdx, bI)
				}
				if len(notWonBoardIdx) == 1 {
					for k := range notWonBoardIdx {
						return k
					}
				}
			}
		}
		return -1
	}
	lastBoard := findLastBoard()
	if lastBoard == -1 {
		return -1
	}
	return solveP1(nums, []board{boards[lastBoard]})
}

func sum(nums []int) int {
	out := 0
	for _, v := range nums {
		out += v
	}
	return out
}