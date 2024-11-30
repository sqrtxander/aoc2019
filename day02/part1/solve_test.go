package main

import "testing"

var INPUT string = `
1,9,10,3,2,3,11,0,99,30,40,50
`[1:]

var EXPECTED int = 3500

func TestSolve(t *testing.T) {
	actual := solve(INPUT, 9, 10)
	if actual != EXPECTED {
		t.Fatalf("Expected %d got %d\n", EXPECTED, actual)
	}
}
