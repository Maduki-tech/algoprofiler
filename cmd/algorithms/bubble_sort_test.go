package algorithms_test

import (
	"reflect"
	"testing"

	"github.com/Maduki-tech/algoprofiler/cmd/algorithms"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		input []int
		want  []int
	}{
		// TODO: Add test cases.
		{"Simple Sort", []int{1, 3, 5, 2, 4, 7, 8}, []int{1, 2, 3, 4, 5, 7, 8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := algorithms.BubbleSort(tt.input)
			// TODO: update the condition below to compare got with tt.want.
			if !reflect.DeepEqual(tt.input, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
