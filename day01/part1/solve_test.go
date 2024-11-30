package main

import "testing"

var INPUT string = `
12
14
1969
100756
`[1:]

var EXPECTED int = 34241

func TestSolve(t *testing.T) {
	actual := solve(INPUT)
	if actual != EXPECTED {
		t.Fatalf("Expected %d got %d\n", EXPECTED, actual)
	}
}