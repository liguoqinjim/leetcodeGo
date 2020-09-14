package lab0066

import (
	"reflect"
	"testing"
)

var cases = []struct {
	input  []int
	output []int
}{
	{input: []int{1, 2, 3}, output: []int{1, 2, 4}},
	{input: []int{4, 3, 2, 1}, output: []int{4, 3, 2, 2}},
	{input: []int{0}, output: []int{1}},
	{input: []int{9}, output: []int{1, 0}},
}

func TestPlusOne(t *testing.T) {
	for _, c := range cases {
		input := append([]int{}, c.input...)
		result := PlusOne(input)
		if !reflect.DeepEqual(result, c.output) {
			t.Errorf("PlusOne(%v)=%v,want %v", c.input, result, c.output)
		}
	}
}
