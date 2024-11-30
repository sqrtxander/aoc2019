package main

import "testing"

var cases = map[int]bool{
	111111: true,
	223450: false,
	123789: false,
}

func TestIsValid(t *testing.T) {
	for pass, expected := range cases {
		if actual := isValid(pass); actual != expected {
			t.Fatalf("%d: expected %t, got %t\n", pass, expected, actual)
		}
	}
}
