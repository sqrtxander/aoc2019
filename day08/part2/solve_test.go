package main

import "testing"

var INPUT string = `
0222112222120000
`[1:]

var EXPECTED string = ` #
# `

func TestSolve(t *testing.T) {
	actual := solve(INPUT, 2, 2)
	if actual != EXPECTED {
		t.Fatalf("Expected %q got %q\n", EXPECTED, actual)
	}
}
