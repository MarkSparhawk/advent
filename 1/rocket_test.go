package main

import (
	"testing"
)

func TestFuelNeeded(t *testing.T) {
	type test struct {
		inputArray []int
		want  int
	}

	tests := []test{
		{inputArray: []int{12},  want: 2},
		{inputArray: []int{14},  want: 2},
		{inputArray: []int{1969},  want: 654},
		{inputArray: []int{100756},  want: 33583},
		{inputArray: []int{12,14,1969,100756},  want: 33583 + 654 + 2 + 2},
	}

	for _, tc := range tests {
		t.Run("Test calculating fuel requirements", func(t *testing.T) {
			got := fuelNeeded(tc.inputArray)
			if got != tc.want {
				t.Errorf("Got %d, wanted %d", got, tc.want)
			}
		})
	}
}

func TestGetFuelNeeded(t *testing.T) {
	type test struct {
		input int
		want  int
	}

	tests := []test{
		{input: 14,  want: 2},
		{input: 1969,  want: 966},
		{input: 100756,  want: 50346},
	}

	for _, tc := range tests {
		t.Run("Test calculating fuel requirements", func(t *testing.T) {
			got := getFuel(tc.input)
			if got != tc.want {
				t.Errorf("Got %d, wanted %d", got, tc.want)
			}
		})
	}
}
