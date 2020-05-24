package cross

import (
	"fmt"
	"sort"
	"strings"
	"strconv"
)

type point struct {
	X int
	Y int
}

func closestDistance(wires [2]string) int {
	trail1 := getTrail(wires[0])
	trail2 := getTrail(wires[1])
	intersections := []point{}

	for _, point1 := range trail1 {
		for _, point2 := range trail2 {
			if point1 == point2 {
				intersections = append(intersections, point1)
			}
		}
	}
	closestIntersection := abs(intersections[0].X) + abs(intersections[0].Y)

	for _, intersection := range intersections {
		distance := abs(intersection.X) + abs(intersection.Y)
		if distance < closestIntersection {
			closestIntersection = distance
		}
	}
	return closestIntersection
}

type stepPoint struct {
	steps int
	point point
}

func shortestPath(wires [2]string) int {
	trail1 := getTrail(wires[0])
	trail2 := getTrail(wires[1])
	intersectionSteps := []int{}
	for step1, point1 := range trail1 {
		for step2, point2 := range trail2 {
			if point1 == point2 {
				intersectionSteps = append(intersectionSteps, step1 + step2 +2)
			}
		}
	}

	sort.Ints(intersectionSteps)
	return intersectionSteps[0]

}



func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func getTrail(wire string) []point {
	instructions := strings.Split(wire,",")
	pointer := point{X: 0, Y: 0}
	trail := []point{}
	for _, instruction := range instructions {
		direction := string(instruction[0])
		distance, err := strconv.Atoi(instruction[1:])
		if err != nil {
			fmt.Println("Bad Instruction")
		}

		switch direction {
		case "R":
			for i:=0; i < distance; i++ {
				pointer.X++
				trail = append(trail, pointer)
			}
		case "L":
			for i:=0; i < distance; i++ {
				pointer.X--
				trail = append(trail, pointer)
			}
		case "U":
			for i:=0; i < distance; i++ {
				pointer.Y--
				trail = append(trail, pointer)
			}
		case "D":
			for i:=0; i < distance; i++ {
				pointer.Y++
				trail = append(trail, pointer)
			}
		}
	}
	return trail
}