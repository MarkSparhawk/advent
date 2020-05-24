package main

import (
	"reflect"
	"testing"
)

func TestDay2Program(t *testing.T) {
	type test struct {
		input []int
		want  []int
	}

	tests := []testProgram{
		{
			input: program{
				code:   []int{1, 1, 1, 4, 99, 5, 6, 0, 99},
				cur:    0,
				input:  0,
				output: 0,
			},
			want: program{
				code:   []int{30, 1, 1, 4, 2, 5, 6, 0, 99},
				cur:    8,
				input:  0,
				output: 0,
			},
		},
	}

	for _, tc := range tests {
		t.Run("Test full program", func(t *testing.T) {
			tc.input.runProgram()
			if reflect.DeepEqual(tc.input, tc.want) != true {
				t.Errorf("Got %+v, wanted %+v", tc.input, tc.want)
			}
		})
	}
}
