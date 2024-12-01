package main

import "testing"

var INPUT string = `
123456789013
`[1:]

var EXPECTED int = 1

func TestSolve(t *testing.T) {
	actual := solve(INPUT, 3, 2)
	if actual != EXPECTED {
		t.Fatalf("Expected %d got %d\n", EXPECTED, actual)
	}
}
