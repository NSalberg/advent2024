package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type count struct {
	num       int
	loopsLeft int
}

var cache = make(map[count]int)

func countFor(num int, loopFor int) int {

	if i, exists := cache[count{num, loopFor}]; exists {
		return i
	}
	total := 0
	if loopFor == 0 {
		return 1
	}

	if i := strconv.Itoa(num); i == "0" {
		c := countFor(1, loopFor-1)
		cache[count{num, loopFor}] = c
		total += c
	} else if i := strconv.Itoa(num); len(i)%2 == 0 {
		lI, _ := strconv.Atoi(i[:len(i)/2])
		rI, _ := strconv.Atoi(i[len(i)/2:])

		cL := countFor(lI, loopFor-1)
		cR := countFor(rI, loopFor-1)
		total += cL + cR
		cache[count{num, loopFor}] = cL + cR
	} else {
		c := countFor(num*2024, loopFor-1)
		total += c
		cache[count{num, loopFor}] = c
	}
	return total
}

func soln1(input string) int {
	q := strings.Fields(input)
	var nums []int
	for _, nS := range q {
		n, _ := strconv.Atoi(nS)
		nums = append(nums, n)
	}

	total := 0
	for _, num := range nums {
		total += countFor(num, 25)
	}

	return total
}

func soln2(input string) int {
	q := strings.Fields(input)
	var nums []int
	for _, nS := range q {
		n, _ := strconv.Atoi(nS)
		nums = append(nums, n)
	}

	total := 0
	for _, num := range nums {
		total += countFor(num, 75)
	}

	return total
}

func main() {

	file, err := os.ReadFile("../../inputs/day11.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := string(file)

	fmt.Println(soln1(input))
	fmt.Println(soln2(input))

}
