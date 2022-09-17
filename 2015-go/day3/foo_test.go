package main

import (
	"testing"
	"fmt"
)

func TestParsing(t *testing.T) {	
	tt := []struct {
		input string
		expected int
	} {
		{input: "", expected: 1},
		{input: ">", expected: 2},
		{input: "^>v<", expected: 4},
		{input: "^v^v^v^v^v", expected: 2},
	}

	for i,tc := range tt {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			if res := parseThings(tc.input); res != tc.expected {
				t.Error(fmt.Sprintf("got %v, expected %v, input %q", res, tc.expected, tc.input))
			}
		})
	}
}

func TestParsingP2(t *testing.T) {
	tt := []struct {
		input string
		expected int
	}{
		{input: "^v", expected: 3},
		{input: "^>v<", expected: 3},
		{input: "^v^v^v^v^v", expected: 11},
	}
	for i,tc := range tt {
		t.Run(fmt.Sprint(i), func(t *testing.T)  {
			if res := parseThings2(tc.input); res != tc.expected {
				t.Errorf("got %v, expected %v, input %q", res, tc.expected, tc.input)
			}
		})
	}
}

