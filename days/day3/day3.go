package day3

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/jeffa5/aoc-19/utils"
)

func Day3(p int, base string) string {
	lines := utils.ReadFileLines(base + "3")

	if len(lines) != 2 {
		return "Expected only 2 lines"
	}

	parts := strings.Split(lines[0], ",")
	points1 := genPoints(parts)

	parts = strings.Split(lines[1], ",")
	points2 := genPoints(parts)

	var d int
	switch p {
	case 1:
		d = findClosestIntersection(points1, points2, true)
	case 2:
		d = findClosestIntersection(points1, points2, false)
	}

	return fmt.Sprintf("%d", d)
}

type Point struct {
	x, y, steps int
}

type Intersect struct {
	x, y, s1, s2 int
}

func findClosestIntersection(points1, points2 []Point, byDist bool) int {
	both := []Intersect{}
	for _, p1 := range points1 {
		for _, p2 := range points2 {
			if p1.x == p2.x && p1.y == p2.y {
				both = append(both, Intersect{p1.x, p1.y, p1.steps, p2.steps})
			}
		}
	}

	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	dist := math.MaxUint32
	for _, i := range both {
		if byDist {
			x := abs(i.x)
			y := abs(i.y)
			if x+y < dist {
				dist = x + y
			}
		} else {
			if i.s1+i.s2 < dist {
				dist = i.s1 + i.s2
			}
		}
	}

	return dist
}

func genPoints(parts []string) []Point {
	points := []Point{}
	point := Point{0, 0, 0}

	for _, part := range parts {
		dir := part[0]
		s, _ := strconv.Atoi(part[1:])

		for i := 0; i < s; i++ {
			point = Point{point.x, point.y, point.steps + 1}
			switch dir {
			case 'R':
				point.x += 1
			case 'U':
				point.y += 1
			case 'D':
				point.y -= 1
			case 'L':
				point.x -= 1
			}
			points = append(points, point)
		}
	}

	return points
}
