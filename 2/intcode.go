package main

import (
	"fmt"
)

func add(i []int, offset int) []int {
	i[i[offset+3]] = i[i[offset+1]] + i[i[offset+2]]
	return i
}

func multiply(i []int, offset int) []int {
	i[i[offset+3]] = i[i[offset+1]] * i[i[offset+2]]
	return i
}

func intcode(i []int, offset int) []int {
	op := i[offset+0]
	if op == 1 {
		fmt.Printf("Addition Instruction at %d: %d\n", offset, i[offset:offset+3])
		return add(i, offset)
	} else if op == 2 {
		fmt.Printf("Mutiplication Instruction at %d: %d\n", offset, i[offset:offset+3])
		return multiply(i, offset)
	} else if op == 99 {
		fmt.Printf("Quit Instruction at %d: %d\n", offset, i[offset])
		return i
	} else {
		fmt.Printf("Unexpected code at %d: %d", offset, i[offset])
		return i
	}
}

func runProgram(i []int) []int {
	for offset := 0; offset < len(i); offset += 4 {
		if i[offset] != 99 {
			i = intcode(i, offset)
		} else {
			return i
		}
	}
	return i
}