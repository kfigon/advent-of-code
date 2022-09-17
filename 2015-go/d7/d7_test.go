package main

import (
	"fmt"
	"testing"
)

func TestEmpty(t *testing.T) {
	p := newProcessor()
	if v := p.readRegistry("x"); v != 0 {
		t.Errorf("Empty read, expected 0, got %v",v)
	}
}

func TestLoad(t *testing.T) {
	p := newProcessor()
	p.load(123, "x")
	v := p.readRegistry("x")
	if v != 123 {
		t.Errorf("Expected %v, got %v", 123, v)
	}
}

func cmp(t *testing.T, p *processor, reg string, exp uint16) {
	v := p.readRegistry(reg)
	if v != exp {
		t.Errorf("Expected %v, got %v in register %q", exp, v, reg)
	}
}
func TestExample(t *testing.T) {
	p := newProcessor()
	p.load(123, "x")
	p.load(456, "y")
	p.and(p.readRegistry("x"), p.readRegistry("y"), "d")
	p.or(p.readRegistry("x"), p.readRegistry("y"), "e")
	p.lshift(p.readRegistry("x"), 2, "f")
	p.rshift(p.readRegistry("y"), 2, "g")
	p.not(p.readRegistry("x"), "h")
	p.not(p.readRegistry("y"), "i")

	expected := make(map[string]uint16)
	expected["d"] = 72
	expected["e"] = 507
	expected["f"] = 492
	expected["g"] = 114
	expected["h"] = 65412
	expected["i"] = 65079
	expected["x"] = 123
	expected["y"] = 456

	for key := range expected {
		cmp(t, p, key, expected[key])
	}
}

func TestProcessExample(t *testing.T)  {
	lines := []string{
		"123 -> x",
		"456 -> y",
		"x AND y -> d",
		"x OR y -> e",
		"x LSHIFT 2 -> f",
		"y RSHIFT 2 -> g",
		"NOT x -> h",
		"NOT y -> i",
	}
	p := newProcessor()

	for _,line := range lines {
		op, err := parseLine(line)
		if err != nil {
			t.Error("Got error during processing", err)
		}
		p.doOperation(op)
	}

	expected := make(map[string]uint16)
	expected["d"] = 72
	expected["e"] = 507
	expected["f"] = 492
	expected["g"] = 114
	expected["h"] = 65412
	expected["i"] = 65079
	expected["x"] = 123
	expected["y"] = 456

	for key := range expected {
		cmp(t, p, key, expected[key])
	}
}

func TestParser(t *testing.T) {
	testCases := []struct {
		input 	string
		expOp	operation
	}{
		{ input: "123 -> x", 	expOp: operation{ operation: load, arg1: 123,destinationReg: "x" } },
		{ input: "456 -> y", 	expOp: operation{ operation: load, arg1: 456, destinationReg: "y" } },
		{ input: "x AND y -> d",expOp: operation{ operation: and, sourceReg1:"x", sourceReg2:"y", destinationReg: "d" } },
		{ input: "x OR y -> e", expOp: operation{ operation: or, sourceReg1:"x", sourceReg2:"y", destinationReg: "e" } },
		{ input: "x LSHIFT 2 -> f", expOp: operation{ operation: lshift, sourceReg1:"x", arg2:2, destinationReg: "f" } },
		{ input: "y RSHIFT 2 -> g", expOp: operation{ operation: rshift, sourceReg1:"y", arg2:2, destinationReg: "g" } },
		{ input: "y RSHIFT x -> g", expOp: operation{ operation: rshift, sourceReg1:"y", sourceReg2:"x", destinationReg: "g" } },
		{ input: "NOT x -> h", 	expOp: operation{ operation: not, sourceReg1: "x", destinationReg: "h" } },
		{ input: "NOT y -> i", 	expOp: operation{ operation: not, sourceReg1: "y", destinationReg: "i" } },
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprintf("%q",tC.input), func(t *testing.T) {
			res, err := parseLine(tC.input)
			exp := tC.expOp
			if err != nil {
				t.Error("Got error, not expected:", err)
			}
			if res.operation != exp.operation {
				t.Errorf("invalid operation exp %v got %v", exp.operation, res.operation)
			}
			if res.arg1 != exp.arg1 {
				t.Errorf("invalid arg1 exp %v got %v", exp.arg1, res.arg1)
			}
			if res.arg2 != exp.arg2 {
				t.Errorf("invalid arg2 exp %v got %v", exp.arg2, res.arg2)
			}
			if res.sourceReg1 != exp.sourceReg1 {
				t.Errorf("invalid sourceReg1 exp %q got %q", exp.sourceReg1, res.sourceReg1)
			}
			if res.sourceReg2 != exp.sourceReg2 {
				t.Errorf("invalid sourceReg2 exp %q got %q", exp.sourceReg2, res.sourceReg2)
			}
			if res.destinationReg != exp.destinationReg {
				t.Errorf("invalid operation exp %q got %q", exp.destinationReg, res.destinationReg)
			}
		})
	}
}

// func TestReg(t *testing.T) {
// 	// input := "x LSHIFT 2 -> f"
// 	input := "x OR y -> e"
// 	reg := regexp.MustCompile(`(\w+)?(\d+)? (\w+) (\w+)?(\d+)? -> (\w+)`)
// 	loadPattern := reg.FindAllStringSubmatch(input, -1)

// 	fmt.Println(loadPattern[0][1:])
// }