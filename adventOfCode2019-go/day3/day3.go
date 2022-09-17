package day3

import (
	"math"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}
func (p point) String() string {
	return "("+strconv.Itoa(p.x)+","+strconv.Itoa(p.y)+")"
}

func (p point) calcDistance() int {
	return p.calcDistanceFromPoint(newPoint(0,0))
}

func (p point) calcDistanceFromPoint(other point) int {
	return int(math.Abs(float64(p.x-other.x)) + math.Abs(float64(p.y-other.y)))
}

func newPoint(x,y int) point { return point{x,y} }

func parsePointsToVectors(in string) []point {
	splitted := strings.Split(in, ",")
	if len(splitted) < 2 {
		return []point{}
	}

	out := make([]point, len(splitted)+1)
	out[0] = newPoint(0,0)
	for i := 0; i < len(splitted); i++ {
		toParse := splitted[i]
		dir := rune(toParse[0])
		val,_ := strconv.Atoi(toParse[1:])
		
		x,y := out[i].x,out[i].y
		switch dir {
		case 'U': y += val
		case 'D': y -= val
		case 'L': x-=val
		case 'R': x+=val
		}
		out[i+1] = newPoint(x,y)
	}
	return out
}

type segment struct {
	start point
	end point
}

func newSegment(start, end point) segment { return segment{start,end} }

// https://www.c-program.com/c-g-a-faq1.html#q8
func (seg segment) intersection(other segment) *point {
	ya := float64(seg.start.y)
	yb := float64(seg.end.y)
	yc := float64(other.start.y)
	yd := float64(other.end.y)

	xa := float64(seg.start.x)
	xb := float64(seg.end.x)
	xc := float64(other.start.x)
	xd := float64(other.end.x)

	denominator := (xb-xa)*(yd-yc)-(yb-ya)*(xd-xc)
	nominator1 := (ya-yc)*(xd-xc)-(xa-xc)*(yd-yc)
	nominator2 := (ya-yc)*(xb-xa)-(xa-xc)*(yb-ya)

	if nominator1 == 0 || denominator == 0 {
		return nil
	}

	r := nominator1/denominator
	s := nominator2/denominator
	
	if !(r <= 1 && r >= 0 && s <= 1 && s >= 0) {
		return nil
	}
	ptr := newPoint(int(xa + r*(xb-xa)), int(ya+r*(yb-ya)))
	return &ptr
}

func (s segment) isPointWithinSegment(p point) bool {
	horizontalLine := s.end.y == s.start.y	
	verticalLine := s.end.x == s.start.x
	
	pointOnTheHorizontalLine := p.y == s.end.y
	pointOnTheVerticalLine := p.x == s.end.x

	withinHorizontalRange := (p.x >= s.start.x && p.x <= s.end.x) || (p.x >= s.end.x && p.x <= s.start.x)
	withinVerticalRange := (p.y >= s.start.y && p.y <= s.end.y) || (p.y >= s.end.y && p.y <= s.start.y)

	return (horizontalLine && pointOnTheHorizontalLine && withinHorizontalRange) || 
			(verticalLine && pointOnTheVerticalLine && withinVerticalRange)	
}

func findIntersections(wire1, wire2 []point) []point {
	isZeroPoint := func(p point) bool {
		return p.x == 0 && p.y == 0
	}

	intersections := make([]point,0)
	for i := 0; i < len(wire1)-1; i++ {
		for j := 0; j < len(wire2)-1; j++ {
			seg1 := newSegment(wire1[i], wire1[i+1])
			seg2 := newSegment(wire2[j], wire2[j+1])
			if intersection := seg1.intersection(seg2); intersection != nil && !isZeroPoint(*intersection) {
				intersections = append(intersections, *intersection)
			}
		}
	}
	return intersections
}

func findMinDistance(in1, in2 string) int {
	wire1 := parsePointsToVectors(in1)
	wire2 := parsePointsToVectors(in2)
	if len(wire1) < 2 || len(wire2) < 2 {
		return -1
	}

	intersections := findIntersections(wire1,wire2)

	var minIntersection *int
	for i := 0; i < len(intersections); i++ {
		v := intersections[i]
		candidate := v.calcDistance()
		if minIntersection == nil || candidate < *minIntersection {
			minIntersection = &candidate
		}
	}
	if minIntersection == nil {
		return -1
	}
	return *minIntersection
}

func findMinSteps(in1, in2 string) int {
	wire1 := parsePointsToVectors(in1)
	wire2 := parsePointsToVectors(in2)

	if len(wire1) < 2 || len(wire2) < 2 {
		return -1
	}

	calcLengthToIntersection := func(wire []point, intersectionPoint point) int {
		distanceToIntersection := 0
		for i := 0; i < len(wire)-1; i++ {
			start,end := wire[i],wire[i+1]

			if newSegment(start,end).isPointWithinSegment(intersectionPoint) {
				return distanceToIntersection + intersectionPoint.calcDistanceFromPoint(start)
			}
			distanceToIntersection += start.calcDistanceFromPoint(end)
		}
		return distanceToIntersection
	}

	intersections := findIntersections(wire1,wire2)
	routeToIntersection := make([]int, len(intersections))
	
	for j := 0; j < len(intersections); j++ {
		intersectionPoint := intersections[j]		
		routeToIntersection[j] = calcLengthToIntersection(wire1, intersectionPoint) +  calcLengthToIntersection(wire2, intersectionPoint)
	}

	var minIntersectionRoute *int
	for i := 0; i < len(routeToIntersection); i++ {
		v := routeToIntersection[i]
		if minIntersectionRoute == nil || v < *minIntersectionRoute {
			minIntersectionRoute = &v
		}
	}
	if minIntersectionRoute == nil {
		return -1
	}
	return *minIntersectionRoute
}