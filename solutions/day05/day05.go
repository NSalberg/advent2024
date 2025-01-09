package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func createOrderingMap(lines string) map[int][]int {
	rules := make(map[int][]int)
	for _, line := range strings.Fields(lines) {
		splitStr := strings.Split(line, "|")

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

func solve(input string, sort bool) int {
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

		if isSorted := slices.IsSortedFunc(pageNumbers, cmp); isSorted && !sort {
			middleNum := pageNumbers[len(pageNumbers)/2]
			total += middleNum
		} else if !isSorted && sort {
			slices.SortFunc(pageNumbers, cmp)
			middleNum := pageNumbers[len(pageNumbers)/2]
			total += middleNum
		}

	}
	return total
}
func soln1(input string) int {
	return solve(input, false)
}

func soln2(input string) int {
	return solve(input, true)
}

func main() {

	file, err := os.ReadFile("../../inputs/day5.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := string(file)

	fmt.Println(soln1(input))
	fmt.Println(soln2(input))

}
