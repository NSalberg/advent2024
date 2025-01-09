package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Scanner struct {
	file       *strings.Reader
	ch         rune // current character
	offset     int  // character offset
	rdOffset   int  // reading offset (position after current character)
	lineOffset int  // current line offset
}

func abs(x int) int {
	if x < 0 {
		return 0 - x
	} else {
		return x
	}
}
func compare(x int, y int) int {
	if x < y {
		return 1 // Increasing
	} else if x > y {
		return -1 // Decreasing
	}
	return 0 // Equal
}
func find_mul(str []byte) [][]byte {
	re := regexp.MustCompile(`mul\((\d+),(\d+\))`)
	return re.FindAll(str, -1)
}

func find_digits(str []byte) (int, int) {
	re := regexp.MustCompile(`\d+`)
	nums := re.FindAll(str, -1)
	num1, _ := strconv.Atoi(string(nums[0]))
	num2, _ := strconv.Atoi(string(nums[1]))
	return num1, num2
}

func soln1(input []byte) int {
	re := regexp.MustCompile(`mul\((\d+),(\d+\))`)
	matches := re.FindAll(input, -1)

	total_mul := 0
	for _, match := range matches {
		num1, num2 := find_digits(match)
		total_mul += num1 * num2
	}

	return total_mul
}

func soln2(input []byte) int {
	re := regexp.MustCompile(`mul\((\d+),(\d+\))|do\(\)|don't\(\)`)
	matches := re.FindAll(input, -1)

	total_mul := 0
	count := true
	for _, match := range matches {

		if string(match) == "do()" {
			count = true
			continue
		} else if string(match) == "don't()" {
			count = false
			continue
		}

		if count == true {
			num1, num2 := find_digits(match)
			total_mul += num1 * num2
		}
	}

	return total_mul
}

func main() {

	file, err := os.ReadFile("../../inputs/day3.txt")
	//file = []byte(`seafood fool mul(1,2) mul(1.2,1)`)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(soln1(file))

	fmt.Println(soln2(file))
}
