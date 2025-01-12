package main

import (
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
	if res != 1928 {
		t.Fatalf(`res = %d`, res)
	}
}

func TestSoln2(t *testing.T) {
	file, err := os.ReadFile("./test.txt")
	if err != nil {
		log.Fatal(err)
	}
	testInput := string(file)

	res := soln2(testInput)
	if res != 2858 {
		t.Fatalf(`res = %d`, res)
	}
}

func TestSoln3(t *testing.T) {
	file, err := os.ReadFile("./test2.txt")
	if err != nil {
		log.Fatal(err)
	}
	testInput := string(file)

	res := soln2(testInput)
	if res != 169 {
		t.Fatalf(`res = %d`, res)
	}
}
