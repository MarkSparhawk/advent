package cross

import (
	"fmt"
	"sort"
	"strings"
	"strconv"
)


func shortestPathMap(wires [2]string) int {
	path1 := getPathMap(wires[0])
	path2 := getPathMap(wires[1])

	intersectionSteps := []int{}
	for location, steps := range path1 {
		if steps2, found := path2[location]; found {
			intersectionSteps = append(intersectionSteps, steps + steps2)
		} 
	}

	sort.Ints(intersectionSteps)
	return intersectionSteps[0]
}


func getPathMap(wire string) map[point]int {

	step := 0
	path := make(map[point] int)
	pointer := point{X: 0, Y: 0}

	instructions := strings.Split(wire,",")
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
				step++
				if _, found := path[pointer]; !found {
					path[pointer] = step
				}
			}
		case "L":
			for i:=0; i < distance; i++ {
				pointer.X--
				step++
				if _, found := path[pointer]; !found {
					path[pointer] = step
				}
			}
		case "U":
			for i:=0; i < distance; i++ {
				pointer.Y--
				step++
				if _, found := path[pointer]; !found {
					path[pointer] = step
				}
			}
		case "D":
			for i:=0; i < distance; i++ {
				pointer.Y++
				step++
				if _, found := path[pointer]; !found {
					path[pointer] = step
				}
			}
		}
	}
	return path
}