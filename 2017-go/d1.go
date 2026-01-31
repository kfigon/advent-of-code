package main

import "fmt"

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
