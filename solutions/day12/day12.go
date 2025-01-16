package main

import (
	"fmt"
	"image"
	"log"
	"os"

	"github.com/NSalberg/advent2024/utils"
)

func calcPerimArea(point image.Point, pointSeen map[image.Point]bool, g utils.Grid) (perim int, area int) {
	perim, area = 0, 1
	region := g[point]
	if pointSeen[point] {
		return 0, 0
	}
	pointSeen[point] = true

	for _, dir := range utils.DirectionsSlice {
		newPoint := point.Add(dir)

		if _, exists := g[newPoint]; !exists || g[newPoint] != region {
			perim += 1
		} else {
			p, a := calcPerimArea(newPoint, pointSeen, g)
			perim += p
			area += a
		}
	}
	return perim, area
}

func soln1(input string) int {
	g := utils.NewGrid(input, nil)
	seenPoints := make(map[image.Point]bool)

	total := 0
	//perim, area := calcPerimArea(image.Pt(7, 4), seenPoints, g)
	//fmt.Println(perim, area)
	for point, _ := range g {
		if !seenPoints[point] {
			perim, area := calcPerimArea(point, seenPoints, g)
			//		fmt.Println(string(region), point, perim, area)
			total += perim * area
		}
	}

	return total
}

func soln2(input string) int {

	total := 0

	return total
}

func main() {

	file, err := os.ReadFile("../../inputs/day12.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := string(file)

	fmt.Println(soln1(input))
	fmt.Println(soln2(input))

}
