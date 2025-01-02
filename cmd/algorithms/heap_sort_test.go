package algorithms_test

import (
	"reflect"
	"testing"

	"github.com/Maduki-tech/algoprofiler/cmd/algorithms"
)

func TestHeap_sort(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		input []int
		want  []int
	}{
		{
			"Simple Test",
			[]int{9, 4, 3, 8, 10, 2, 5},
			[]int{2, 3, 4, 5, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := algorithms.Heap_sort(tt.input)
			// TODO: update the condition below to compare got with tt.want.
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Heap_sort() = %v, want %v", got, tt.want)
			}
		})
	}
}
