package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

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

func isLineSafe(line string) bool {
	line = strings.TrimSpace(line)
	splitLn := strings.Split(line, " ")
	num_items := len(line)
	nums := make([]int, num_items)

	var prev_val *int
	var increasing = 0
	for i, str := range splitLn {
		n, _ := strconv.Atoi(str)
		nums[i] = n

		if prev_val == nil {
			prev_val = &nums[i]
		} else {
			if abs(n-*prev_val) > 3 {
				return false
			}

			if compare(*prev_val, n) == 0 {
				return false
			}

			if increasing == 0 {
				increasing = compare(*prev_val, n)
			} else if increasing != compare(*prev_val, n) {
				return false
			}
			prev_val = &nums[i]
		}

	}
	return true
}

func isLineSafePt2(line string) bool {
	splitLn := strings.Split(line, " ")
	if isLineSafe(line) {
		return true
	}
	for j := range len(splitLn) {
		rmStr := make([]string, len(splitLn))
		copy(rmStr, splitLn)
		rmStr = append(rmStr[:j], rmStr[j+1:]...)
		joinLn := strings.Join(rmStr, " ")
		if isLineSafe(joinLn) {
			return true
		}
	}
	return false

}

func soln1(scanner *bufio.Scanner) int {
	var num_safe = 0
	for scanner.Scan() {
		line := scanner.Text()
		if isLineSafe(line) {
			num_safe += 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return num_safe
}

func soln2(scanner *bufio.Scanner) int {
	var num_safe = 0
	for scanner.Scan() {
		line := scanner.Text()
		if isLineSafePt2(line) {
			num_safe += 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return num_safe
}

func main() {

	file, err := os.Open("../../inputs/day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fmt.Println(soln1(scanner))

	if _, err := file.Seek(0, io.SeekStart); err != nil {
		log.Fatal(err)
	}

	scanner = bufio.NewScanner(file)
	fmt.Println(soln2(scanner))
}
