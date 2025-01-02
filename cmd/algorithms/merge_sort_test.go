package algorithms

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Already sorted",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Reverse order",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Unsorted array",
			input:    []int{38, 27, 43, 3, 9, 82, 10},
			expected: []int{3, 9, 10, 27, 38, 43, 82},
		},
		{
			name:     "Empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Single element",
			input:    []int{42},
			expected: []int{42},
		},
		{
			name:     "Duplicate elements",
			input:    []int{5, 1, 2, 5, 3, 5},
			expected: []int{1, 2, 3, 5, 5, 5},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Merge_sort(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("mergeSort(%v) = %v; want %v", tc.input, result, tc.expected)
			}
		})
	}
}
