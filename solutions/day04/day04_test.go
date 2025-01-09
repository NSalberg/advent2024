package main

import (
	"testing"
)

func TestSoln1(t *testing.T) {
	testString := "MMMSXXMASM\nMSAMXMSMSA\nAMXSXMAAMM\nMSAMASMSMX\nXMASAMXAMM\nXXAMMXXAMA\nSMSMSASXSS\nSAXAMASAAA\nMAMMMXMMMM\nMXMXAXMASX"

	res := soln1(testString)
	if res != 18 {
		t.Fatalf(`res = %d`, res)
	}
}

func TestSoln2(t *testing.T) {
	testString := "MMMSXXMASM\nMSAMXMSMSA\nAMXSXMAAMM\nMSAMASMSMX\nXMASAMXAMM\nXXAMMXXAMA\nSMSMSASXSS\nSAXAMASAAA\nMAMMMXMMMM\nMXMXAXMASX"

	res := soln2(testString)
	if res != 9 {
		t.Fatalf(`res = %d`, res)
	}
}
