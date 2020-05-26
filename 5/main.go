package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/MarkSparhawk/advent/intcode"
)

func main() {
	lines := readFile("input")

	// Part 1
	code := getCode(lines[0])
	p := intcode.Program{Code: code, Cur: 0, Input: 1, Output: 0}
	p.Run()
	fmt.Printf("Day 5.1 Answer: %d\n", p.Output)

	// Part 2
	code = getCode(lines[0])
	p = intcode.Program{Code: code, Cur: 0, Input: 5, Output: 0}
	p.Run()
	fmt.Printf("Day 5.2 Answer: %d\n", p.Output)
}

func getCode(s string) []int {
	tmp := strings.Split(s, ",")
	values := make([]int, 0, len(tmp))
	for _, raw := range tmp {
		v, err := strconv.Atoi(raw)
		if err != nil {
			log.Print(err)
			continue
		}
		values = append(values, v)
	}
	return values
}

func readFile(filename string) []string {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	return txtlines
}
