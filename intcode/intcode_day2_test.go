package intcode

import (
	"reflect"
	"testing"
)

func TestDay2Program(t *testing.T) {
	type test struct {
		Input []int
		want  []int
	}

	tests := []testProgram{
		{
			input: Program{
				Code:   []int{1, 1, 1, 4, 99, 5, 6, 0, 99},
				Cur:    0,
				Input:  0,
				Output: 0,
			},
			want: Program{
				Code:   []int{30, 1, 1, 4, 2, 5, 6, 0, 99},
				Cur:    8,
				Input:  0,
				Output: 0,
			},
		},
	}

	for _, tc := range tests {
		t.Run("Test full Program", func(t *testing.T) {
			tc.input.runProgram()
			if reflect.DeepEqual(tc.input, tc.want) != true {
				t.Errorf("Got %+v, wanted %+v", tc.input, tc.want)
			}
		})
	}
}
