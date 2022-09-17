package day7

import (
	"aoc2019/intcode"
	"strconv"

	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// https://adventofcode.com/2019/day/7
func TestPart1Examples(t *testing.T) {
	testCases := []struct {
		exp int
		code []int
	}{
		{43210, []int{3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0}},
		{54321, []int{3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0}},
		{65210, []int{3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33, 1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0}},
	}
	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := doPart1(tc.code)	
			assert.Equal(t, tc.exp, got)
		})
	}
}

func TestPart1(t *testing.T) {
	file := readFile(t)
	assert.Equal(t, 844468, doPart1(file))
}

func TestPart2Examples(t *testing.T) {
	testCases := []struct {
		exp int
		code []int
	}{
		{139629729, []int{3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5}},
		{18216, []int{3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54,-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4,53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10}},
	}
	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := doPart2(tc.code)	
			assert.Equal(t, tc.exp, got)
		})
	}
}

func TestPart2(t *testing.T) {
	file := readFile(t)
	assert.Equal(t, 4215746, doPart2(file))
}

func readFile(t *testing.T) []int{
	file,err := os.Open("data.txt")
	require.NoError(t, err)
	defer file.Close()

	content, err := io.ReadAll(file)
	require.NoError(t, err)
	data := string(content)
	splitted := strings.Split(data, ",")
	out := make([]int,len(splitted))
	
	for i := 0; i < len(splitted); i++ {
		v,err := strconv.Atoi(splitted[i])	
		require.NoError(t,err)
		out[i] = v
	}
	return out
}

func newCode(in []int) []int {
	out := make([]int,len(in))
	for i := 0; i < len(out); i++ {
		out[i]=in[i]
	}
	return out
}
func doLoop(min, max int, singleIteration func(a,b,c,d,e int) int) int {
	var maxSignal int
	for a := min; a < max; a++ {
		for b := min; b < max; b++ {
			for c := min; c < max; c++ {
				for d := min; d < max; d++ {
					for e := min; e < max; e++ {
	
						if a==b || a==c || a==d || a==e || 
							b==c || b==d || b==e ||
							c==d || c==e || 
							d == e {
							continue
						}
						outSignal := singleIteration(a,b,c,d,e)
						if outSignal > maxSignal {
							maxSignal = outSignal
						}
					}
				}
			}
		}
	}
	return maxSignal	
}

func doPart1(code []int) int {

	processSingleComputer := func(c *intcode.Computer, in1, in2 int) {
		c.SetUserInput(in1)
		c.SetUserInput(in2)
		c.Calc()
	}

	singleIteration := func(a,b,c,d,e int) int {
		comp1:=intcode.NewComputer(newCode(code))
		comp2:=intcode.NewComputer(newCode(code))
		comp3:=intcode.NewComputer(newCode(code))
		comp4:=intcode.NewComputer(newCode(code))
		comp5:=intcode.NewComputer(newCode(code))

		processSingleComputer(comp1, a,0)
		processSingleComputer(comp2, b,comp1.GetOutput())
		processSingleComputer(comp3, c,comp2.GetOutput())
		processSingleComputer(comp4, d,comp3.GetOutput())
		processSingleComputer(comp5, e,comp4.GetOutput())

		return comp5.GetOutput()
	}

	return doLoop(0,5, singleIteration)
}

func doPart2(code []int) int {

	singleIteration := func(a,b,c,d,e int) int {
		comp1:=intcode.NewComputer(newCode(code))
		comp2:=intcode.NewComputer(newCode(code))
		comp3:=intcode.NewComputer(newCode(code))
		comp4:=intcode.NewComputer(newCode(code))
		comp5:=intcode.NewComputer(newCode(code))

		comp1.SetUserInput(a)
		comp2.SetUserInput(b)
		comp3.SetUserInput(c)
		comp4.SetUserInput(d)
		comp5.SetUserInput(e)
		
		done := false
		for !done {
			comp1.SetUserInput(comp5.GetOutput())
			comp1.CalcTilOutput()

			comp2.SetUserInput(comp1.GetOutput())
			comp2.CalcTilOutput()

			comp3.SetUserInput(comp2.GetOutput())
			comp3.CalcTilOutput()
			
			comp4.SetUserInput(comp3.GetOutput())
			comp4.CalcTilOutput()
			
			comp5.SetUserInput(comp4.GetOutput())
			done = comp5.CalcTilOutput()

			// current implementation of inputs is
			// a cyclic list - clean that!
			comp1.ClearUserInput()
			comp2.ClearUserInput()
			comp3.ClearUserInput()
			comp4.ClearUserInput()
			comp5.ClearUserInput()
		}
		return comp5.GetOutput()
	}

	return doLoop(5,10, singleIteration)
}