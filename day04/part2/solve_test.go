package main

import "testing"

var cases = map[int]bool{
	112233: true,
	123444: false,
	111122: true,
	223334: true,
	222333: false,
	999999: false,
	777789: false,
	344569: true,
}

func TestIsValid(t *testing.T) {
	for pass, expected := range cases {
		if actual := isValid(pass); actual != expected {
			t.Fatalf("%d: expected %t, got %t\n", pass, expected, actual)
		}
	}
}
