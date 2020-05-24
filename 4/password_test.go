package password

import (
	"testing"
	"reflect"
)

func TestValidPassword(t *testing.T) {
	type test struct {
		input int 
		want bool 
	}

	tests := []test {
		{input: 11111, want: false},
		{input: 111111, want: true},
		{input: 123456, want: false},
		{input: 654321, want: false},
	}

	for _, tc := range tests {
		t.Run("Checking Password Validator", func(t *testing.T){
			got := validPassword(tc.input, checkLength, checkAdjacentPair, checkNotDecrementing)
			if got != tc.want {
				t.Errorf("Got %t, wanted %t", got, tc.want)
			}
		})
	}
}

func TestCheckNotDecrementing(t *testing.T) {
	type test struct {
		input []int
		want bool
	}

	tests := []test {
		{input: []int{1,2,3,4,5,6}, want: true},
		{input: []int{1,1,1,1,1,1}, want: true},
		{input: []int{1,6,5,4,3,2}, want: false},
		{input: []int{6,5,4,3,2,1}, want: false},
	}

	for _, tc := range tests {
		t.Run("Check not decrementing", func(t *testing.T){
			got := checkNotDecrementing(tc.input)
			if got != tc.want {
				t.Errorf("Got %t, wanted %t", got, tc.want)
			}
		})
	}
}

func TestCheckAdjacentPair(t *testing.T) {
	type test struct {
		input []int
		want bool
	}

	tests := []test {
		{input: []int{1,2,3,4,5,6}, want: false},
		{input: []int{1,1,1,1,1,1}, want: true},
		{input: []int{1,6,5,4,3,2}, want: false},
		{input: []int{6,5,4,3,2,1}, want: false},
		{input: []int{1,2,3,4,5,5}, want: true},
	}

	for _, tc := range tests {
		t.Run("Check not decrementing", func(t *testing.T){
			got := checkAdjacentPair(tc.input)
			if got != tc.want {
				t.Errorf("Got %t, wanted %t", got, tc.want)
			}
		})
	}
}

func TestCheckHasDouble(t *testing.T) {
	type test struct {
		input []int
		want bool
	}

	tests := []test {
		{input: []int{1,2,3,4,5,6}, want: false},
		{input: []int{1,1,1,1,1,1}, want: false},
		{input: []int{1,1,1,1,2,2}, want: true},
		{input: []int{6,5,4,3,2,1}, want: false},
		{input: []int{1,2,3,4,5,5}, want: true},
	}

	for _, tc := range tests {
		t.Run("Check has pair of 2 but not more", func(t *testing.T){
			got := checkHasDouble(tc.input)
			if got != tc.want {
				t.Errorf("Got %t, wanted %t", got, tc.want)
			}
		})
	}
}

func TestCheckLength(t *testing.T) {
	type test struct {
		input []int
		want bool
	}

	tests := []test {
		{input: []int{1}, want: false},
		{input: []int{1,2,3,4,5}, want: false},
		{input: []int{1,2,3,4,5,6}, want: true},
		{input: []int{1,2,3,4,5,6,7}, want: false},
		{input: []int{}, want: false},
	}

	for _, tc := range tests {
		t.Run("Check length validator", func(t *testing.T){
			got := checkLength(tc.input)
			if got != tc.want {
				t.Errorf("Got %t, wanted %t", got, tc.want)
			}
		})
	}
}
func TestIntToSlice(t *testing.T) {
	type test struct {
		input int
		want []int
	}

	tests := []test {
		{input: 123456, want: []int{1,2,3,4,5,6}},
		{input: 111, want: []int{1,1,1}},
		{input: 321, want: []int{3,2,1}},
	}

	for _, tc := range tests {
		t.Run("Checking int to slice converter", func(t *testing.T){
			got := intToSlice(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Got %v, wanted %v", got, tc.want)
			}
		})
	}
}

func TestDriverPart1(t *testing.T) {
	type test struct {
		min int
		max int
		want int
	}
	
	tests := []test {
		{min: 111111, max: 111112, want:2},
		{min: 264793, max: 803935, want:966},
	}

	for _, tc := range tests {
		t.Run("Checking how many possible keys", func(t *testing.T){
			got := len(getKeys(tc.min, tc.max))
			if got != tc.want {
				t.Errorf("Got %d, wanted %d", got, tc.want)
			}
		})
	}
}

func TestDriverPart2(t *testing.T) {
	type test struct {
		min int
		max int
		want int
	}
	
	tests := []test {
		{min: 111111, max: 111112, want:0},
		{min: 111122, max: 111123, want:1},
		{min: 264793, max: 803935, want:628},
	}

	for _, tc := range tests {
		t.Run("Checking how many possible keys", func(t *testing.T){
			got := len(getKeys2(tc.min, tc.max))
			if got != tc.want {
				t.Errorf("Got %d, wanted %d", got, tc.want)
			}
		})
	}
}

func BenchmarkDriverPart2(b *testing.B) {
	type test struct {
		min int
		max int
		want int
	}
	
	tests := []test {
		{min: 111111, max: 111112, want:0},
		{min: 111122, max: 111123, want:1},
		{min: 264793, max: 803935, want:628},
	}

	for i:=0; i < b.N; i++ {
		for _, tc := range tests {
			getKeys2(tc.min, tc.max)
		}
	}
}