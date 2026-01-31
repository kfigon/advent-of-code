package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

// go run . -d=1 -p=1 <<EOF
// 1122
// 1111
// 1234
// 91212129
// EOF
func main() {
	fmt.Println("hello, provide file using stdin")

	part := flag.Int("p", 1, "part")
	day := flag.Int("d", 1, "day")
	flag.Parse()

	allowedScenarios := map[Task]func([]string){
		{Day: 1, Part: 1}: d1p1,
		{Day: 1, Part: 2}: d1p2,
	}

	if *part != 1 && *part != 2 {
		return
	}
	task := Task{Day: *day, Part: *part}
	runFn, ok := allowedScenarios[task]
	if !ok {
		fmt.Fprintf(os.Stderr, "unknown input: %+v", task)
		os.Exit(1)
		return
	}
	data, err := collectInput(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading input: %v", err)
		os.Exit(1)
		return
	}
	runFn(data)
}

func collectInput(r io.Reader) ([]string, error) {
	out := []string{}
	s := bufio.NewScanner(r)
	for s.Scan() {
		out = append(out, s.Text())
	}

	return out, s.Err()
}

type Task struct {
	Day  int
	Part int
}

func toInts(s string) ([]int, error) {
	out := make([]int, 0, len(s))
	for i, v := range s {
		got, err := strconv.Atoi(string(v))
		if err != nil {
			return nil, fmt.Errorf("error parsing char #%d: %w", i, err)
		}
		out = append(out, got)

	}
	return out, nil
}

func d1p1(input []string) {
	for _, v := range input {
		fmt.Println(solveD1(v, d1Provider))
	}
}

type NextIdProvider func(int, []int) int

func d1Provider(i int, data []int) int {
	nextId := i + 1
	if nextId >= len(data) {
		return 0
	}
	return nextId
}

func d2Provider(i int, data []int) int {
	return (i + len(data)/2) % len(data)
}

func solveD1(input string, provider NextIdProvider) int {
	data, err := toInts(input)
	if err != nil {
		return 0
	}
	out := 0
	for i := 0; i < len(data); i++ {
		nextId := provider(i, data)
		if data[i] == data[nextId] {
			out += data[i]
		}
	}
	return out
}

func d1p2(input []string) {
	for _, v := range input {
		fmt.Println(solveD1(v, d2Provider))
	}
}
