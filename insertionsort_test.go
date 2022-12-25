package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertionSortInts(t *testing.T) {
	lessInt := func(a, b int) bool {
		return a < b
	}

	tests := []struct {
		name  string
		array []int
	}{
		{
			name:  "Test small array",
			array: []int{10, 8, 6, 4, 2, 0},
		},
		{
			name:  "Test bigger array",
			array: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
		{
			name:  "Test empty array",
			array: []int{},
		},
		{
			name:  "Test array of lenght 1",
			array: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InsertionSort(tt.array, lessInt)

			assert.IsNonDecreasing(t, tt.array, "Array should be in ascending order")
		})
	}
}
