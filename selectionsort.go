package sort

func SelectionSort[T any](array []T, less func(T, T) bool) {
	for i := 0; i < len(array); i++ {
		_, minIndex := min(array, less)
		array[i], array[minIndex] = array[minIndex], array[i]
	}
}
