package algorithms

import (
	"reflect"
	"testing"
)

func TestQuick_sort(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		input []int
		low   int
		want  []int
	}{
		{
			"Simple Test", []int{10, 7, 8, 9, 1, 5}, 0, []int{1, 5, 7, 8, 9, 10},
		},
		{
			"Empty Test", []int{}, 0, []int{},
		},
		{
			"Single Element Test", []int{1}, 0, []int{1},
		},
		{
			"Reverse Test", []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, 0, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			"Random Test", []int{10, 7, 8, 9, 1, 5, 2, 3, 4, 6}, 0, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Quick_sort(tt.input, tt.low, len(tt.input)-1)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Quick_sort() = %v, want %v", got, tt.want)
			}
		})
	}
}
