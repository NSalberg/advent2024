package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func testEquation(target int, nums []int, pt2 bool) int {
	fmt.Println(target, nums)
	if len(nums) == 1 {
		num := nums[0]
		if target == num {
			return 1
		} else {
			return 0
		}
	}

	lastNumIdx := len(nums) - 1
	lastNum := nums[lastNumIdx]

	fmt.Println("Div")
	if quotient, remainder := target/lastNum, target%lastNum; remainder == 0 {
		if n := testEquation(quotient, nums[:lastNumIdx], pt2); n > 0 {
			return target
		}
	}

	fmt.Println("Sub")
	if n := testEquation(target-lastNum, nums[:lastNumIdx], pt2); n > 0 {
		return target
	}

	if pt2 && len(nums) >= 2 {
		fmt.Println("Chop")
		//Chop current num off
		newTarget, _ := strconv.Atoi(strings.TrimSuffix(strconv.Itoa(target), strconv.Itoa(lastNum)))
		if newTarget == target {
			return 0
		}

		if n := testEquation(newTarget, nums[:lastNumIdx], pt2); n > 0 {
			return target
		}
	}

	return 0
}

func soln1(input string) int {
	total := 0
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		split := strings.Split(line, ":")

		target, _ := strconv.Atoi(split[0])
		numbers := make([]int, len(strings.Fields(split[1])))
		for i, snumber := range strings.Fields(split[1]) {
			numbers[i], _ = strconv.Atoi(snumber)
		}
		//fmt.Println(target, numbers, testEquation(target, numbers))
		total += testEquation(target, numbers, false)
	}

	return total
}

func soln2(input string) int {
	total := 0
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		split := strings.Split(line, ":")

		target, _ := strconv.Atoi(split[0])
		numbers := make([]int, len(strings.Fields(split[1])))
		for i, snumber := range strings.Fields(split[1]) {
			numbers[i], _ = strconv.Atoi(snumber)
		}

		fmt.Println("START", target, numbers)
		//fmt.Println(testEquation(target, numbers, true))
		total += testEquation(target, numbers, true)
	}

	return total
}

func main() {

	file, err := os.ReadFile("../../inputs/day7.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := string(file)

	fmt.Println(soln1(input))
	fmt.Println(soln2(input))

}
