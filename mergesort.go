package sort

import "sync"

func MergeSort[T any](array []T, less func(T, T) bool) {
	if len(array) <= 1 {
		return
	}
	buffer := make([]T, len(array))
	mergeSortRecursive(array, buffer, less)
}

func mergeSortRecursive[T any](array []T, buffer []T, less func(T, T) bool) {
	length := len(array)
	mid := length / 2

	if length <= 1 {
		return
	}

	if length > 100 {

		var wg sync.WaitGroup
		wg.Add(2)

		go func(wgp *sync.WaitGroup) {
			defer (*wgp).Done()
			mergeSortRecursive(array[:mid], buffer[:mid], less)
		}(&wg)
		go func(wgp *sync.WaitGroup) {
			defer (*wgp).Done()
			mergeSortRecursive(array[mid:], buffer[mid:], less)
		}(&wg)

		wg.Wait()

	} else {
		mergeSortRecursive(array[:mid], buffer[:mid], less)
		mergeSortRecursive(array[mid:], buffer[mid:], less)
	}

	mergeSortedArrays(array[:mid], array[mid:], buffer, less)

	copy(array, buffer)
}

func mergeSortedArrays[T any](array1 []T, array2 []T, buffer []T, less func(T, T) bool) {
	var i1, i2, i3 int

	for i1 < len(array1) && i2 < len(array2) {
		if less(array2[i2], array1[i1]) {
			buffer[i3] = array2[i2]
			i2++
		} else {
			buffer[i3] = array1[i1]
			i1++
		}
		i3++
	}

	if i1 < len(array1) {
		copy(buffer[i3:], array1[i1:])
		return
	}

	if i2 < len(array2) {
		copy(buffer[i3:], array2[i2:])
		return
	}
}
