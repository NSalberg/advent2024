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

func solve(input string, sorted bool) int {
	split := strings.Split(input, "\n\n")

	rules := createOrderingMap(split[0])

	cmp := func(a, b string) int {
		c, _ := strconv.Atoi(a)
		d, _ := strconv.Atoi(b)
		for _, ruleNumber := range rules[c] {
			if d == ruleNumber {
				return -1
			}
		}
		return 0
	}

	total := 0
	for _, line := range strings.Fields(split[1]) {
		if sPageNumbers := strings.Split(line, ","); slices.IsSortedFunc(sPageNumbers, cmp) == sorted {
			slices.SortFunc(sPageNumbers, cmp)
			middleNum, _ := strconv.Atoi(sPageNumbers[len(sPageNumbers)/2])
			total += middleNum
		}

	}
	return total
}
func soln1(input string) int {
	return solve(input, true)
}

func soln2(input string) int {
	return solve(input, false)
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
