package main

import (
	"io/ioutil"
	"fmt"
)

// https://adventofcode.com/2015/day/3

type coordinate struct {
	x int
	y int
}

func parseSingleChar(char string, inputCoordinate coordinate) coordinate {	
	up := "^"
	down := "v"
	left := "<"
	right := ">"
	
	switch char {
	case up: inputCoordinate.y--
	case down: inputCoordinate.y++
	case left: inputCoordinate.x--
	case right: inputCoordinate.x++
	default: fmt.Println("Invalid char", char)
	}

	return inputCoordinate
}
func validateData(data string) bool {
	return data != ""
}

func parseThings(data string) int {
	if !validateData(data) {
		return 1
	}

	visited := make(map[coordinate]bool)
	lastCoordinate := coordinate{0,0}
	visited[lastCoordinate] = true

	for i := 0; i < len(data); i++ {
		lastCoordinate = parseSingleChar(string(data[i]), lastCoordinate)
		visited[lastCoordinate]=true
	}

	return len(visited)
}

func parseThings2(data string) int {
	if !validateData(data) {
		return 1
	}
	
	visitedBySanta := make(map[coordinate]bool)
	visitedByRobot := make(map[coordinate]bool)

	lastCoordinateRobot := coordinate{0,0}
	lastCoordinateSanta := coordinate{0,0}

	visitedByRobot[lastCoordinateRobot] = true
	visitedBySanta[lastCoordinateSanta] = true

	for i := 0; i < len(data); i++ {
		char := string(data[i])

		if i % 2 == 0 {
			lastCoordinateSanta = parseSingleChar(char, lastCoordinateSanta)
			visitedBySanta[lastCoordinateSanta] = true
		} else {
			lastCoordinateRobot = parseSingleChar(char, lastCoordinateRobot)
			visitedByRobot[lastCoordinateRobot] = true
		}
	}

	uniqueNumbers := len(visitedBySanta)
	for key := range visitedByRobot {
		if _,ok := visitedBySanta[key]; !ok {
			uniqueNumbers++
		}
	}
	return uniqueNumbers
}

func main() {
	path := "input.txt"
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("error when reading file")
		return
	}
	part1 := parseThings(string(bytes))
	part2 := parseThings2(string(bytes))
	fmt.Println("part1",part1)
	fmt.Println("part2",part2)
}
