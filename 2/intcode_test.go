package main

import (
	"reflect"
	"testing"
)

func TestAddition(t *testing.T) {
	type test struct {
		input []int
		want []int
	}

	tests := []test {
		{input: []int{1,0,0,0,99}, want: []int{2,0,0,0,99}},
	}

	for _, tc := range tests {
		t.Run("Test addition opscode", func(t *testing.T) {
			got := add(tc.input, 0)
			if reflect.DeepEqual(got, tc.want) != true {
				t.Errorf("Got %d, wanted %d", got, tc.want)
			}
		})
	}
}
func TestMultiplication(t *testing.T) {
	type test struct {
		input []int
		want []int
	}

	tests := []test {
		{input: []int{2,3,0,3,99}, want: []int{2,3,0,6,99}},
	}

	for _, tc := range tests {
		t.Run("Test multiplication opscode", func(t *testing.T) {
			got := multiply(tc.input, 0)
			if reflect.DeepEqual(got, tc.want) != true {
				t.Errorf("Got %d, wanted %d", got, tc.want)
			}
		})
	}
}


func TestIntCode(t *testing.T) {
	type test struct {
		input []int
		want []int
	}

	tests := []test {
		{input: []int{1,0,0,0,99}, want: []int{2,0,0,0,99}},
		{input: []int{2,3,0,3,99}, want: []int{2,3,0,6,99}},
	}

	for _, tc := range tests {
		t.Run("Test multiplication opscode", func(t *testing.T) {
			got := intcode(tc.input, 0)
			if reflect.DeepEqual(got, tc.want) != true {
				t.Errorf("Got %d, wanted %d", got, tc.want)
			}
		})
	}
}

func TestProgram(t *testing.T) {
	type test struct {
		input []int
		want []int
	}

	tests := []test {
		{input: []int{99,}, want: []int{99,}},
		{input: []int{1,1,1,4,99,5,6,0,99}, want: []int{30,1,1,4,2,5,6,0,99}},
	}

	for _, tc := range tests {
		t.Run("Test full program", func(t *testing.T) {
			got := runProgram(tc.input)
			if reflect.DeepEqual(got, tc.want) != true {
				t.Errorf("Got %d, wanted %d", got, tc.want)
			}
		})
	}
}
