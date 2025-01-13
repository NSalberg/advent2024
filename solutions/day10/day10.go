package main

import (
	"fmt"
	"image"
	"log"
	"os"

	"github.com/NSalberg/advent2024/utils"
)

type pathPair struct {
	start, end image.Point
}

func searchForI(g utils.Grid, s image.Point, i rune) map[image.Point]bool {
	ret := make(map[image.Point]bool)

	for _, dir := range utils.DirectionsSlice {
		step := s.Add(dir)
		if _, exists := g[step]; !exists {
			continue
		}

		oldI := int(g[s] - '0')
		newI := int(g[step] - '0')
		if newI-oldI != 1 {
			continue
		}
		if g[step] == i {
			ret[step] = true
		} else {
			for k, v := range searchForI(g, step, i) {
				ret[k] = v
			}
		}
	}

	return ret

}

func soln1(input string) int {
	total := 0

	var startPoints []image.Point
	handleRune := func(p image.Point, r rune) {
		if r == '0' {
			startPoints = append(startPoints, p)
		}
	}

	g := utils.NewGrid(input, handleRune)

	for _, p := range startPoints {
		total += len(searchForI(g, p, '9'))
	}

	return total
}

func soln2(input string) int {
	total := 0

	return total
}

func main() {

	file, err := os.ReadFile("../../inputs/day10.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := string(file)

	fmt.Println(soln1(input))
	fmt.Println(soln2(input))

}
