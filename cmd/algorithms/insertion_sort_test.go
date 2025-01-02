package algorithms_test

import (
	"reflect"
	"testing"

	"github.com/Maduki-tech/algoprofiler/cmd/algorithms"
)

func TestInsertion_sort(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		input []int
		want  []int
	}{
		// TODO: Add test cases.
		{
			"Simple Test",
			[]int{5, 4, 3, 2, 1},
			[]int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := algorithms.Insertion_sort(tt.input)
			// TODO: update the condition below to compare got with tt.want.
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Insertion_sort() = %v, want %v", got, tt.want)
			}
		})
	}
}
