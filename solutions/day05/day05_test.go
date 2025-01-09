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
	if res != 143 {
		t.Fatalf(`res = %d`, res)
	}
}

/*
func TestSoln2(t *testing.T) {
	testString := "MMMSXXMASM\nMSAMXMSMSA\nAMXSXMAAMM\nMSAMASMSMX\nXMASAMXAMM\nXXAMMXXAMA\nSMSMSASXSS\nSAXAMASAAA\nMAMMMXMMMM\nMXMXAXMASX"

	res := soln2(testString)
	if res != 9 {
		t.Fatalf(`res = %d`, res)
	}
}
*/
