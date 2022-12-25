package sort

func InsertionSort[T any](array []T, less func(T, T) bool) {
	var lenght int = len(array)
	if lenght <= 1 {
		return
	}

	for i := 1; i < lenght; i++ {
		for j := i; j > 0 && less(array[j], array[j-1]); j-- {
			array[j], array[j-1] = array[j-1], array[j]
		}
	}
}
