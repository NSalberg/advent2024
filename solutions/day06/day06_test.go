package main

import (
	"image"
	"log"
	"os"
	"testing"
)

func TestSoln1(t *testing.T) {

	file, err := os.ReadFile("./test.txt")
	if err != nil {
		log.Fatal(err)
	}
	testInput := string(file)

	res := soln1(testInput)
	if res != 41 {
		t.Fatalf(`res = %d`, res)
	}
}

func TestSolve(t *testing.T) {

	file, err := os.ReadFile("./testloop.txt")
	if err != nil {
		log.Fatal(err)
	}
	testInput := string(file)

	grid := newGrid(testInput)
	var guardPos image.Point
	guardDir := image.Pt(0, -1)

	for k, v := range grid {
		if string(v) == "^" {
			guardPos = k
			break
		}
	}

	res := solve(grid, guardPos, guardDir)
	if res != false {
		t.Fatalf(`res = %t`, res)
	}
}

func TestSoln2(t *testing.T) {
	file, err := os.ReadFile("./test.txt")
	if err != nil {
		log.Fatal(err)
	}
	testInput := string(file)

	res := soln2(testInput)
	if res != 6 {
		t.Fatalf(`res = %d`, res)
	}
}
func TestSoln3(t *testing.T) {
	file, err := os.ReadFile("./test3.txt")
	if err != nil {
		log.Fatal(err)
	}
	testInput := string(file)

	res := soln2(testInput)
	if res != 4 {
		t.Fatalf(`res = %d`, res)
	}
}
