package main

import (
	"fmt"
	"regexp"
	"strings"
	"strconv"
	"os"
	"bufio"
	"time"
)

const (
	TURN_ON = iota
	TURN_DOWN 
	TOGGLE
)

type rectangle struct {
	aX, aY, bX, bY int
}

func (r *rectangle) rows() (int, int) {
	return r.aY, r.bY
}

func (r *rectangle) cols() (int, int) {
	return r.aX, r.bX
}

func (r *rectangle) size() int {
	difX := r.bX - r.aX  + 1
	difY := r.bY - r.aY  + 1
	return difX * difY
}

func parseCoordinate(inputStr string) (int, int, error) {
	parts := strings.Split(inputStr, ",")
	if parts == nil || len(parts) != 2 {
		return 0,0, fmt.Errorf("Error in parsing coordinate: %q", inputStr)
	}
	x,err := strconv.Atoi(parts[0])
	if err != nil {
		return 0,0, fmt.Errorf("Error in parsing coordinate: %q", inputStr)
	}
	y,err := strconv.Atoi(parts[1])
	if err != nil {
		return 0,0, fmt.Errorf("Error in parsing coordinate: %q", inputStr)
	}
	return x,y,nil
}

func parseCmd(inputStr string) (int, rectangle, error) {
	reg, err := regexp.Compile(`(\w+\s?\w+?) (\d+,\d+) through (\d+,\d+)`)
	if err != nil {
		return 0, rectangle{}, fmt.Errorf("Error in parsing: %q, error: %v", inputStr, err)
	}

	parts := reg.FindAllStringSubmatch(inputStr, -1)
	if parts == nil || len(parts) != 1 || len(parts[0]) != 4 {
		return 0, rectangle{}, fmt.Errorf("Error in regex: %q", inputStr)
	}
	parsedParts := parts[0]
	operation := -1
	if strings.Contains(parsedParts[1], "turn on") {
		operation = TURN_ON
	} else if strings.Contains(parsedParts[1], "toggle") {
		operation = TOGGLE
	} else if strings.Contains(parsedParts[1], "turn off") {
		operation = TURN_DOWN
	}

	aX, aY, err := parseCoordinate(parsedParts[2])
	if err != nil {
		return 0, rectangle{}, fmt.Errorf("Error in coordinate: %q", inputStr)
	}

	bX, bY, err := parseCoordinate(parsedParts[3])
	if err != nil {
		return 0, rectangle{}, fmt.Errorf("Error in coordinate: %q", inputStr)
	}

	return operation, rectangle{aX,aY,bX,bY}, nil
}

type processor struct {
	numOfLit int
	table [][]bool
}

func newProcessor() *processor {
	tab := make([][]bool, 1000)
	for i := 0; i < len(tab); i++ {
		tab[i] = make([]bool, 1000)
	}

	return &processor {
		numOfLit: 0,
		table: tab,
	}
}

func (p *processor) processCmd(operation int, rect rectangle) {
	rowStart, rowEnd := rect.rows()
	colStart, colEnd := rect.cols()

	for row := rowStart; row < rowEnd+1; row++ {
		for col := colStart; col < colEnd+1; col++ {
			currentLight := p.table[row][col]
			if operation == TURN_ON && !currentLight {
            	p.table[row][col] = true
				p.numOfLit += 1
			} else if operation == TURN_DOWN && currentLight {
            	p.table[row][col] = false
				p.numOfLit -= 1
			} else if operation == TOGGLE && !currentLight {
            	p.table[row][col] = true
				p.numOfLit += 1
			} else if operation == TOGGLE && currentLight {
	            p.table[row][col] = false
            	p.numOfLit -= 1
			}
		}
	}
}

func readLine(path string) []string{
	lines := make([]string,0)
	inFile, err := os.Open(path)
	if err != nil {
	   fmt.Println(err.Error() + `: ` + path)
	   return lines
	}
	defer inFile.Close()
  
	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
	  lines = append(lines, scanner.Text()) // the line
	}

	return lines
  }

func main() {

	lines := readLine("input.txt")

	start := time.Now()
	p := newProcessor()
	for _,line := range lines {
		op, rec, err := parseCmd(line)
		if err != nil {
			fmt.Println("error during proicessing", err)
			return
		}
		p.processCmd(op, rec)
	}

	fmt.Println("Done", p.numOfLit)
	timeDiff := time.Since(start)
	fmt.Println("took", timeDiff)
}