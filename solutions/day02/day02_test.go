package main

import (
	"testing"
)

func TestSoln1(t *testing.T) {
	testString := "3   4 4   3 2   5 1   3 3   9 3   3"
	res := soln1(testString)
	if res != 11 {
		t.Fatalf(`res = %d`, res)
	}

}
func TestSoln2(t *testing.T) {
	testString := "3   4 4   3 2   5 1   3 3   9 3   3"
	res := soln2(testString)
	if res != 31 {
		t.Fatalf(`res = %d`, res)
	}

}
