package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func nextIndex(ix []int, lengthsFunc func(i int) int) {
	for j := len(ix) - 1; j >= 0; j-- {
		ix[j]++
		if j == 0 || ix[j] < lengthsFunc(j) {
			return
		}
		ix[j] = 0
	}
}

func product[T any](sets ...[]T) [][]T {
	lengths := func(i int) int { return len(sets[i]) }
	var product [][]T
	for ix := make([]int, len(sets)); ix[0] < lengths(0); nextIndex(ix, lengths) {
		var r []T

		for j, k := range ix {
			r = append(r, sets[j][k])
		}
		product = append(product, r)
	}
	return product
}

var directions = product([]int{1, 0, -1}, []int{1, 0, -1})

type loc struct {
	x int
	y int
}

func (coord loc) move(dir []int) loc {
	return loc{x: coord.x + dir[0], y: coord.y + dir[1]}
}

type grid map[loc]rune

func newGrid(input string) grid {
	grid := make(map[loc]rune)
	linescanner := bufio.NewScanner(strings.NewReader(input))

	y := 0
	for linescanner.Scan() {
		for x, letter := range linescanner.Text() {
			grid[loc{x: x, y: y}] = letter
		}
		y += 1
	}
	return grid
}

func countXmas(coord loc, xmasGrid grid) int {
	num_xmas := 0
	xmas := "MAS"

	for _, d := range directions {

		new_coord := coord
		for i, letter := range xmas {
			new_coord = new_coord.move(d)
			grid_letter, exists := xmasGrid[new_coord]

			if exists != true {
				break
			}

			if byte(grid_letter) != byte(letter) {
				break
			}
			if i == 2 {
				num_xmas += 1
			}
		}

	}

	return num_xmas
}

func containsMS(msMap map[rune]bool) bool {
	letters := []rune{'M', 'S'}
	for _, letter := range letters {
		if _, exists := msMap[letter]; !exists {
			return false
		}
	}
	return true
}

func isX_MAS(coord loc, xmasGrid grid) int {
	dirs := [][][]int{{{1, 1}, {-1, -1}}, {{-1, 1}, {1, -1}}}

	for _, dir := range dirs {
		masMap := make(map[rune]bool)
		for _, diag := range dir {
			letter, lexists := xmasGrid[coord.move(diag)]
			masMap[letter] = true
			if lexists != true {
				return 0
			}
		}
		if !containsMS(masMap) {
			return 0
		}
	}

	return 1
}

func soln1(input string) int {
	grid := newGrid(input)
	num_xs := 0
	for k, v := range grid {
		if string(v) == "X" {
			num_xs += countXmas(k, grid)
		}
	}

	return num_xs
}

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

func main() {

	file, err := os.ReadFile("../../inputs/day4.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := string(file)

	fmt.Println(soln1(input))
	fmt.Println(soln2(input))

}
