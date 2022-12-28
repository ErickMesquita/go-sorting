package sort

type Heap[T any] struct {
	array  []T
	length int
}

func HeapSort[T any](array []T, less func(T, T) bool) {
	if len(array) <= 1 {
		return
	}

	var h Heap[T] = Heap[T]{array: (array), length: len(array)}
	var end int = len(array) - 1

	for i := h.parent(end); i >= 0; i-- {
		h.Heapify(i, less)
	}

	for i := end; i > 0; i-- {
		h.swap(i, 0)
		h.length--
		h.Heapify(0, less)
	}
}

func (h Heap[T]) Heapify(index int, less func(T, T) bool) {
	left, left_exists := h.leftChild(index)
	right, right_exists := h.rightChild(index)

	greater := index

	if left_exists && less(h.get(greater), h.get(left)) {
		greater = left
	}
	if right_exists && less(h.get(greater), h.get(right)) {
		greater = right
	}
	if greater != index {
		h.swap(index, greater)
		h.Heapify(greater, less)
	}
}

func (h Heap[T]) get(index int) T {
	return h.array[index]
}

func (h Heap[T]) swap(index1, index2 int) {
	h.array[index1], h.array[index2] = h.array[index2], h.array[index1]
}

func (h Heap[T]) parent(index int) (parentIndex int) {
	if index <= 0 {
		parentIndex = 0
	} else {
		parentIndex = (index - 1) / 2
	}
	return
}

func (h Heap[T]) leftChild(index int) (leftChildIndex int, leftChildExists bool) {
	if index < 0 {
		leftChildIndex = 0
	} else {
		leftChildIndex = index*2 + 1
	}

	leftChildExists = leftChildIndex < h.length
	return
}

func (h Heap[T]) rightChild(index int) (rightChildIndex int, rightChildExists bool) {
	if index < 0 {
		rightChildIndex = 0
	} else {
		rightChildIndex = index*2 + 2
	}
	rightChildExists = rightChildIndex < h.length
	return
}
