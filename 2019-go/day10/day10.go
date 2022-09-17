package day10

import (
	"math"
	"sort"
)


func isAsteroid(c rune) bool { return c =='#' }
type spaceMap []string

func (s spaceMap) charAt(x,y int) rune {
	return rune(s[y][x])
}

func (s spaceMap) rows() int { return len(s) }
func (s spaceMap) cols() int { return len(s[0]) }

type result struct {
	x,y,visibility int
}

func (s spaceMap) findBestPlace() result {
	asteroids := s.filterAsteroids()
	candidate := result{}
	for _,coord := range asteroids {
		v := s.analyzePosition(coord, asteroids)
		if v > candidate.visibility {
			candidate.visibility = v
			candidate.x = coord.x
			candidate.y = coord.y
		}
	}
	return candidate
}

func (s spaceMap) filterAsteroids() []point {
	asteroids := []point{}
	for x := 0; x < s.cols(); x++ {
		for y := 0; y < s.rows(); y++ {
			if !isAsteroid(s.charAt(x,y)) {
				continue
			}

			asteroids = append(asteroids, point{x,y})
		}
	}
	return asteroids
}

// returns number of points visible form the starting point
func (s spaceMap) analyzePosition(startingPoint point, asteroids []point) int {
	visiblePoints := map[float64]float64{}

	for _,v := range asteroids {
		if startingPoint.eq(v) {
			continue
		}

		trig := startingPoint.trigonometryVersion(v)
		angle, ok := visiblePoints[trig.degree]
		if !ok {
			visiblePoints[trig.degree] = trig.length
		} else if trig.degree < angle{
			visiblePoints[trig.degree] = trig.length
		}
	}
	return len(visiblePoints)
}

func (s spaceMap) orderByVaporization(startingPoint point) []point {
	asteroids := s.filterAsteroids()
	sortedByAngle := sortByAngle(startingPoint, asteroids)
	groupedByAngle := groupByAngleAndSort(sortedByAngle)

	uniqueSortedAngles := []float64{}
	for ang := range groupedByAngle {
		uniqueSortedAngles = append(uniqueSortedAngles, ang)
	}
	sort.Float64s(uniqueSortedAngles)

	out := []point{}
	for len(groupedByAngle) > 0 {
		for _, angle := range uniqueSortedAngles {
			pointsOnAngle := groupedByAngle[angle]
			if len(pointsOnAngle) == 0 {
				delete(groupedByAngle, angle)
				continue
			} 
			out = append(out, pointsOnAngle[0].pt)
			groupedByAngle[angle] = pointsOnAngle[1:]
		}
	}

	return out
}

func sortByAngle(startingPoint point, pts []point) []orderingData {
	sortedByAngle := []orderingData{}
	for i := 0; i < len(pts); i++ {
		if startingPoint.eq(pts[i]) {
			continue
		} 
		sortedByAngle = append(sortedByAngle, newOrderingObj(startingPoint, pts[i]))
	}
	sort.Slice(sortedByAngle, func(i, j int) bool {
		return sortedByAngle[i].angle < sortedByAngle[j].angle
	})
	return sortedByAngle
}

func groupByAngleAndSort(sorted []orderingData) map[float64][]orderingData {
	groupedByAngle := map[float64][]orderingData{}
	for i := 0; i < len(sorted); i++ {
		v := sorted[i]
		if _, ok := groupedByAngle[v.angle]; !ok {
			groupedByAngle[v.angle] = []orderingData{}	
		}
		groupedByAngle[v.angle] = append(groupedByAngle[v.angle], v)
	}
	for key := range groupedByAngle {
		sort.Slice(groupedByAngle[key], func(i, j int) bool {
			return groupedByAngle[key][i].length < groupedByAngle[key][j].length
		})
	}
	return groupedByAngle
}

type orderingData struct {
	pt	point
	angle, length float64
}

func newOrderingObj(starting point, pt point) orderingData {
	trig := starting.trigonometryVersion(pt)
	return orderingData{
		pt: pt,
		angle: trig.degree,
		length: trig.length,
	}
}


type point struct { x,y int }
func (p point) eq(other point) bool {
	return p.x == other.x && p.y == other.y
}

type trigonometryPoint struct {
	degree, length float64
}
func (p point) trigonometryVersion(end point) trigonometryPoint {
	deltaX := math.Pow((float64(end.x) - float64(p.x)), 2); 
	deltaY := math.Pow((float64(end.y) - float64(p.y)), 2);

	distance := math.Sqrt(deltaY + deltaX);

	const RADIAN_TO_DEGREE = 180 / math.Pi
	radians := math.Atan2(float64(end.y) - float64(p.y), float64(end.x) - float64(p.x))
	
	angle := (RADIAN_TO_DEGREE * radians) + 90
	if angle < 0 {
		angle += 360
	}
	return trigonometryPoint{
		degree: angle,
		length: distance,
	}
}