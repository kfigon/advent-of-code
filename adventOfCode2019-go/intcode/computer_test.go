package intcode

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCalc(t *testing.T) {
	testCases := []struct {
		in []int
		exp []int
	}{
		{[]int{1,0,0,3,99}, []int{1,0,0,2,99}},
		{[]int{1,0,0,0,99}, []int{2,0,0,0,99}},
		{[]int{1,9,10,3,2,3,11,0,99,30,40,50}, []int{3500,9,10,70,2,3,11,0,99,30,40,50}},
		{[]int{2,3,0,3,99}, []int{2,3,0,6,99}},
		{[]int{1002,4,3,4,33}, []int{1002,4,3,4,99}},
		{[]int{2,4,4,5,99,0}, []int{2,4,4,5,99,9801}},
		{[]int{1,1,1,4,99,5,6,0,99}, []int{30,1,1,4,2,5,6,0,99}},
		{[]int{1002,5,3,5,99,33}, []int{1002,5,3,5,99,99}},
		{[]int{109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99}, []int{109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99}},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.in), func(t *testing.T) {
			got := NewComputer(tc.in).Calc()
			assert.Equal(t, tc.exp, got)
		})
	}
}

func TestOpcodeExtraction(t *testing.T) {
	testCases := []struct {
		val int
		exp int		
	}{
		{1002, 2},
		{2, 2},
		{1112, 12},
		{1199, 99},
		{99, 99},
	}
	for _, tc := range testCases {
		t.Run(strconv.Itoa(tc.val), func(t *testing.T) {
			got := opcode(tc.val).extractOpcode()
			assert.Equal(t, tc.exp, got)
		})
	}
}

func TestModeExtraction(t *testing.T) {
	op := opcode(12398)
	assert.Equal(t, 98, op.extractOpcode())
	assert.Equal(t, 3, op.mode(0))
	assert.Equal(t, 2, op.mode(1))
	assert.Equal(t, 1, op.mode(2))
	
	assert.Equal(t, 0, op.mode(3))
	assert.Equal(t, 0, op.mode(4))
	assert.Equal(t, 0, op.mode(5))
	assert.Equal(t, 0, op.mode(6))
	assert.Equal(t, 0, op.mode(7))
}

func TestHandleInput(t *testing.T) {
	in := []int{3,0,4,1,99}
	c := NewComputer(in)
	c.SetUserInput(15)
	c.Calc()
	
	require.NotNil(t, c.userInput)
	require.NotNil(t, c.userOutput)

	assert.Equal(t, 0, c.userOutput)
	assert.Equal(t, 15, c.userInput.next())
	assert.Equal(t, []int{15,0,4,1,99}, c.instructions)
}

func TestInputWhenEmpty(t *testing.T) {
	in := newInputHandler()
	assert.Equal(t, 0, in.next())
	assert.Equal(t, 0, in.next())
	assert.Equal(t, 0, in.next())
	assert.Equal(t, 0, in.next())
}

func TestInputWhenSingle(t *testing.T) {
	in := newInputHandler()
	in.add(5)
	assert.Equal(t, 5, in.next())
	assert.Equal(t, 5, in.next())
	assert.Equal(t, 5, in.next())
	assert.Equal(t, 5, in.next())
}

func TestInputWhenMultiple(t *testing.T) {
	in := newInputHandler()
	in.add(5)
	in.add(1)
	assert.Equal(t, 5, in.next())
	assert.Equal(t, 1, in.next())
	assert.Equal(t, 5, in.next())
	assert.Equal(t, 1, in.next())
	assert.Equal(t, 5, in.next())
}

func TestOutputBigNums(t *testing.T) {
	testCases := []struct {
		in []int
		exp int		
	}{
		{[]int{1102,34915192,34915192,7,4,7,99,0}, 1219070632396864},
		{[]int{104,1125899906842624,99}, 1125899906842624},
	}
	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {		
			comp := NewComputer(tc.in)
			comp.Calc()
			assert.Equal(t, tc.exp, comp.GetOutput())
		})
	}
}
func TestMoreOutputCases(t *testing.T) {
	testCases := []struct {
		in []int
		exp int		
	}{
		{[]int{109, -1, 4, 1, 99}, -1},
		{[]int{109, -1, 104, 1, 99}, 1},
		{[]int{109, -1, 204, 1, 99}, 109},
		{[]int{109, 1, 9, 2, 204, -6, 99}, 204},
		{[]int{109, 1, 109, 9, 204, -6, 99}, 204},
		{[]int{109, 1, 209, -1, 204, -106, 99}, 204},
	}
	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {		
			comp := NewComputer(tc.in)
			comp.Calc()
			assert.Equal(t, tc.exp, comp.GetOutput())
		})
	}
}

func TestMoreOutputCasesWithInputs(t *testing.T) {
	testCases := []struct {
		in []int
		exp int		
	}{
		{[]int{109, 1, 3, 3, 204, 2, 99}, 5},
		{[]int{109, 1, 203, 2, 204, 2, 99}, 5},
	}
	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {		
			comp := NewComputer(tc.in)
			comp.SetUserInput(tc.exp)
			comp.Calc()
			assert.Equal(t, tc.exp, comp.GetOutput())
		})
	}
}

func TestRelativeCase(t *testing.T) {
	in := []int{109, 1, 203, 2, 204, 2, 99}
	comp := NewComputer(in)
	comp.SetUserInput(15)
	outputCode := comp.Calc()
	assert.Equal(t, 15, comp.GetOutput())
	assert.Equal(t, []int{109,1,203,15,204,2,99}, outputCode)
}