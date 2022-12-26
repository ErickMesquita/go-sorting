package sort

// min finds the minimum element in a slice and returns its value and index
//
// The `less` function is used to compare the slice's elements
func min[T any](array []T, less func(T, T) bool) (minValue T, minIndex int) {
	minIndex = 0
	for i := 0; i < len(array); i++ {
		if less(array[i], array[minIndex]) {
			minIndex = i
		}
	}
	minValue = array[minIndex]
	return
}
