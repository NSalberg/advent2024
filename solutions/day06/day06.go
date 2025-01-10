package main

import (
	"fmt"
	"image"
	"log"
	"os"
	"strings"
)

func (g grid) print() {
	// Determine the bounds of the grid
	var minX, minY, maxX, maxY int
	for p := range g {
		if p.X < minX {
			minX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		}
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}

	// Print the grid row by row
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if val, exists := g[image.Pt(x, y)]; exists {
				fmt.Print(string(val)) // Print the rune
			} else {
				fmt.Print(".") // Print a placeholder for empty cells
			}
		}
		fmt.Println() // New line after each row
	}
}

type grid map[image.Point]rune

func newGrid(input string) grid {
	grid := make(map[image.Point]rune)
	for y, line := range strings.Fields(input) {
		for x, letter := range line {
			grid[image.Point{X: x, Y: y}] = letter
		}
	}
	return grid
}

func (grid grid) countNumRune(a rune) int {
	total := 0
	for _, v := range grid {
		if v == a {
			total += 1
		}
	}
	return total
}

var rotations = map[image.Point]image.Point{
	image.Pt(0, 1):  image.Pt(-1, 0),
	image.Pt(1, 0):  image.Pt(0, 1),
	image.Pt(0, -1): image.Pt(1, 0),
	image.Pt(-1, 0): image.Pt(0, -1),
}

type guard struct {
	pos image.Point
	dir image.Point
}

func advanceGuard(g grid, guardPos image.Point, guardDir image.Point) (newPos image.Point, newDir image.Point) {

	newGP := guardPos.Add(guardDir)
	newGR, inGrid := g[newGP]
	for inGrid && newGR == '#' {
		guardDir = rotations[guardDir]
		newGP = guardPos.Add(guardDir)
		newGR, inGrid = g[newGP]
	}

	return guardPos.Add(guardDir), guardDir

}

func solve(g grid, guardPos image.Point, guardDir image.Point) bool {
	seen := make(map[guard]bool)
	for {
		_, inGrid := g[guardPos]
		if !inGrid {
			//Guard has left so exit the loop
			break
		}

		if _, exists := seen[guard{pos: guardPos, dir: guardDir}]; exists {
			return false
		} else {
			seen[guard{pos: guardPos, dir: guardDir}] = true
		}
		g[guardPos] = 'X'
		guardPos, guardDir = advanceGuard(g, guardPos, guardDir)
	}
	return true
}

func soln1(input string) int {
	grid := newGrid(input)
	var guardPos image.Point
	guardDir := image.Pt(0, -1)

	for k, v := range grid {
		if string(v) == "^" {
			guardPos = k
			break
		}
	}

	solve(grid, guardPos, guardDir)

	return grid.countNumRune('X')
}

func soln2(input string) int {
	g := newGrid(input)
	var guardPos image.Point
	guardDir := image.Pt(0, -1)

	for k, v := range g {
		if string(v) == "^" {
			guardPos = k
			break
		}
	}

	totalLoops := 0

	seen := make(map[image.Point]bool)
	for {

		newGP, newGD := advanceGuard(g, guardPos, guardDir)

		if _, inGrid := g[newGP]; inGrid {
			if seen[newGP] == false {
				g[newGP] = '#'
				seen[newGP] = true
			}
		} else {
			break
		}
		g[guardPos] = '^'
		//g.print()

		if !solve(g, guardPos, guardDir) {
			totalLoops += 1
		}

		g[guardPos] = '.'
		guardPos = newGP
		guardDir = newGD
		g[newGP] = '.'
		//fmt.Println(totalLoops)

	}
	//g.print()

	return totalLoops
}

func main() {

	file, err := os.ReadFile("../../inputs/day6.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := string(file)

	fmt.Println(soln1(input))
	fmt.Println(soln2(input))

}
