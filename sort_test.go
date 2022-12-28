package sort

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type sorter[T any] func([]T, func(T, T) bool)

var sorters = []struct {
	sorter sorter[int]
	name   string
}{
	{SelectionSort[int], "SelectionSort"},
	{InsertionSort[int], "InsertionSort"},
	{MergeSort[int], "MergeSort"},
	{QuickSort[int], "QuickSort"},
	{QuickSort3way[int], "QuickSort 3 way"},
	{HeapSort[int], "HeapSort"},
}

func TestSortingInts(t *testing.T) {
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
			name:  "Test empty array",
			array: []int{},
		},
		{
			name:  "Test array of lenght 1",
			array: []int{1},
		},
		{
			name: "Test longer array",
			array: []int{20, 19, 18, 17, 17, 17, 16, 16, 15, 14,
				13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
	}

	for _, ss := range sorters {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ss.sorter(tt.array, lessInt)
				assert.IsNonDecreasingf(t, tt.array, "Array should be in ascending order. Sorter: %s", ss.name)
			})
		}
	}
}

func FuzzSortInts(f *testing.F) {
	const max_length uint = 50_000

	f.Add(max_length) // Use f.Add to provide a seed corpus

	f.Fuzz(func(t *testing.T, ArraySize uint) {
		ArraySize = (ArraySize - 1) % max_length
		var array = make([]int, ArraySize)

		for i := uint(0); i < ArraySize; i++ {
			array[i] = int(rand.Int63() - math.MaxInt64/2)
		}
		var originalArray []int
		copy(originalArray, array)

		for _, ss := range sorters {
			ss.sorter(array, lessInt)
			assert.IsNonDecreasingf(t, array, "Array should be in ascending order. Sorter: %s", ss.name)
			copy(array, originalArray)
		}
	},
	)
}

func BenchmarkSortFromFile(b *testing.B) {
	// There are 2212202 lines in this file (and a blank line in the end)
	const filename = "testdata/integersArray.txt"
	const max_nrows = 100_000

	var testSizes = []struct {
		nrows int
	}{
		{nrows: 100},
		{nrows: 1_000},
		{nrows: 10_000},
		{nrows: 100_000},
	}

	file, err := os.Open(filename)
	if err != nil {
		b.Log(err)
	}
	defer file.Close()

	var originalArray []int = make([]int, max_nrows)
	var array []int = make([]int, max_nrows)
	var i int = 0
	for i < max_nrows {
		n, err := fmt.Fscan(file, &array[i])
		if n == 0 || err != nil {
			break
		}
		i++
	}
	copy(originalArray, array)
	b.ResetTimer()
	for _, ss := range sorters {
		ss.sorter(array, lessInt)
		copy(array, originalArray)
	}
	b.Cleanup(func() { copy(array, originalArray) })
	for _, ss := range sorters {
		for _, v := range testSizes {
			b.Run(fmt.Sprintf("sorter: %s, input size: %d", ss.name, v.nrows), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					ss.sorter(array[:v.nrows], lessInt)
				}
			})
		}
	}
}
