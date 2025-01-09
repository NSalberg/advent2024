package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func isCorrectlyOrdered(pageNumbers []int, rules map[int][]int) bool {
	seen := make(map[int]bool)

	for _, number := range pageNumbers {
		seen[number] = true
		for _, rule := range rules[number] {
			if _, seenEarly := seen[rule]; seenEarly {
				return false
			}
		}
	}

	return true
}

func createOrderingMap(lines string) map[int][]int {
	rules := make(map[int][]int)
	for _, line := range strings.Fields(lines) {
		splitStr := strings.Split(line, "|")

		// If len is to we are on a line with rules
		key, _ := strconv.Atoi(splitStr[0])
		value, _ := strconv.Atoi(splitStr[1])

		rules[key] = append(rules[key], value)
	}
	return rules
}

func createCmpFunc(orderMap map[int][]int) func(a, b int) int {
	return func(a, b int) int {
		for _, ruleNumber := range orderMap[a] {
			if b == ruleNumber {
				return -1
			}
		}
		return 0
	}
}

func soln1(input string) int {

	lines := strings.Split(input, "\n\n")

	total := 0

	rules := createOrderingMap(lines[0])
	cmp := createCmpFunc(rules)
	for _, line := range strings.Fields(lines[1]) {
		sPageNumbers := strings.Split(line, ",")
		pageNumbers := make([]int, len(sPageNumbers))

		for i, snum := range sPageNumbers {
			pageNumbers[i], _ = strconv.Atoi(snum)
		}
		isSorted := slices.IsSortedFunc(pageNumbers, cmp)
		fmt.Println(pageNumbers, isSorted)
		if isSorted {
			middleNum := pageNumbers[len(pageNumbers)/2]
			total += middleNum
		}
	}

	fmt.Println(total)

	total = 0
	for _, line := range strings.Fields(lines[1]) {
		sPageNumbers := strings.Split(line, ",")

		pageNumbers := make([]int, len(sPageNumbers))

		for i, snum := range sPageNumbers {
			pageNumbers[i], _ = strconv.Atoi(snum)
		}

		if isCorrectlyOrdered(pageNumbers, rules) {
			// Get the middle number
			middleNum := pageNumbers[len(pageNumbers)/2]
			total += middleNum
		}
	}

	return total
}

/*
func soln2(input string) int {

	grid := newGrid(input)

	num_xmas := 0
	for k, v := range grid {
		if v == 'A' {
			num_xmas += isX_MAS(k, grid)
		}
	}
	return num_xmas
}
*/

func main() {

	file, err := os.ReadFile("../../inputs/day5.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := string(file)

	fmt.Println(soln1(input))
	//	fmt.Println(soln2(input))

}
