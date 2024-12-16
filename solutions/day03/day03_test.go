package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestSoln1(t *testing.T) {
	testString := "7 6 4 2 1\n 1 2 7 8 9\n 9 7 6 2 1\n 1 3 2 4 5\n 8 6 4 4 1\n 1 3 6 7 9\n"
	res := soln1(bufio.NewScanner(strings.NewReader(testString)))
	if res != 2 {
		t.Fatalf(`res = %d`, res)
	}
}

func TestSoln2(t *testing.T) {
	testString := "7 6 4 2 1\n 1 2 7 8 9\n 9 7 6 2 1\n 1 3 2 4 5\n 8 6 4 4 1\n 1 3 6 7 9\n"
	res := soln2(bufio.NewScanner(strings.NewReader(testString)))
	if res != 4 {
		t.Fatalf(`res = %d`, res)
	}
}

func TestIsLineSafe(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{input: "7 6 4 2 1", expected: true},
		{input: "1 2 7 2 1", expected: false},
		{input: "1 3 2 6 5", expected: false},
		{input: "10 13 14 15 16", expected: true},
		{input: "10 10 11 13 16", expected: false},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			res := isLineSafe(tc.input)
			if res != tc.expected {
				t.Fatalf("For input %q, expected %t but got %t", tc.input, tc.expected, res)
			}
		})
	}
}

func TestIsLineSafePt2(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{input: "7 6 4 2 1", expected: true},
		{input: "1 2 7 2 1", expected: false},
		{input: "9 7 6 2 1", expected: false},
		{input: "1 3 2 4 5", expected: true},
		{input: "8 6 4 4 1", expected: true},
		{input: "1 3 6 7 9", expected: true},
		{input: "13 9 11 9 6", expected: true},
		{input: "48 46 47 49 51 54 56", expected: true},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			res := isLineSafePt2(tc.input)
			if res != tc.expected {
				t.Fatalf("For input %q, expected %t but got %t", tc.input, tc.expected, res)
			}
		})
	}
}
