package sort

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Test struct {
	name  string
	array []int
}

var tests []Test = []Test{
	{
		name:  "Test small sorted array",
		array: []int{0, 0, 0, 0, 1, 1, 1, 2, 2, 2},
	},
	{
		name:  "Test small reversed array",
		array: []int{2, 2, 2, 1, 1, 1, 0, 0, 0},
	},
	{
		name:  "Test pivot between",
		array: []int{0, 0, 0, 0, 1, 1, 1, 2, 2, 2, 2, 1},
	},
	{
		name:  "Test small array",
		array: []int{10, 8, 6, 4, 2, 0},
	},
	{
		name: "Test longer array",
		array: []int{20, 19, 18, 17, 17, 17, 16, 16, 15, 14,
			13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
	},
	{
		name: "Test enormous array",
		array: []int{
			10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 456, 789, 132, 4, 321, 12, 0, 5, 12,
			10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 456, 789, 132, 4, 321, 12, 0, 5, 12,
			10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 456, 789, 132, 4, 321, 12, 0, 5, 12,
			10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 456, 789, 132, 4, 321, 12, 0, 5, 12,
			10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 456, 789, 132, 4, 321, 12, 0, 5, 12,
			10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 456, 789, 132, 4, 321, 12, 0, 5, 12,
			10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 456, 789, 132, 4, 321, 12, 0, 5, 12,
		},
	},
}

func TestQuickSort(t *testing.T) {

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QuickSort(tt.array, lessInt)

			assert.IsNonDecreasing(t, tt.array, "Array should be in ascending order")
		})
	}
}

func TestQuickSort3way(t *testing.T) {

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QuickSort3way(tt.array, lessInt)
			fmt.Println(tt.array)

			assert.IsNonDecreasing(t, tt.array, "Array should be in ascending order")
		})
	}
}

func FuzzQuickSort(f *testing.F) {
	f.Add(uint(22122022)) // Use f.Add to provide a seed corpus

	f.Fuzz(func(t *testing.T, ArraySize uint) {
		var array = make([]int, ArraySize)
		//log.Println("ArraySize:", ArraySize)

		for i := uint(0); i < ArraySize; i++ {
			array[i] = int(rand.Int63() - math.MaxInt64/2)
		}
		/*
			fmt.Println("Array:")
			for i := uint(0); i < ArraySize && i < 30; i++ {
				fmt.Printf("%d, ", array[i])
			}
		*/
		QuickSort(array, lessInt)

		assert.IsNonDecreasing(t, array, "Array should be in ascending order")
	},
	)
}

func BenchmarkQuickSortFromFile(b *testing.B) {
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
		log.Fatal(err)
	}
	defer file.Close()

	var array []int = make([]int, max_nrows)
	var i int = 0
	for i < max_nrows {
		n, err := fmt.Fscan(file, &array[i])
		if n == 0 || err != nil {
			break
		}
		i++
	}
	b.ResetTimer()
	for _, v := range testSizes {
		b.Run(fmt.Sprintf("input size: %d", v.nrows), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				QuickSort(array[:v.nrows], lessInt)
			}
		})
	}
}

func lessInt(a, b int) bool {
	return a < b
}
