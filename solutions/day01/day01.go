package main

import (
	"fmt"
	"os"
	"slices"
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

func soln1(str string) int {

	words := strings.Fields(str)

	var leftArr []int
	var rightArr []int
	for i := 0; i < len(words); i++ {
		num, _ := strconv.Atoi(words[i])
		if i%2 == 0 {
			leftArr = append(leftArr, num)
		} else {
			rightArr = append(rightArr, num)
		}
	}

	slices.Sort(leftArr)
	slices.Sort(rightArr)

	dif := 0
	for i := 0; i < len(leftArr); i++ {
		dif += abs(leftArr[i] - rightArr[i])
	}

	return dif

}

func soln2(str string) int {

	words := strings.Fields(str)

	var leftArr []int
	right := make(map[int]int)

	for i := 0; i < len(words); i++ {
		num, _ := strconv.Atoi(words[i])
		if i%2 == 1 {
			cur, _ := right[num]
			right[num] = cur + 1
		} else {
			leftArr = append(leftArr, num)
		}
	}

	sim := 0
	for i := 0; i < len(leftArr); i++ {
		sim += leftArr[i] * right[leftArr[i]]
	}

	return sim

}

func main() {
	f, _ := os.ReadFile("../../inputs/day1-1.txt")
	str := string(f)
	fmt.Println(soln1(str))
	fmt.Println(soln2(str))
}
