package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var heapMock Heap[int] = Heap[int]{
	array:  []int{0, 2, 4, 6, 8, 10, 12},
	length: 7,
}

func TestHeap_parent(t *testing.T) {
	tests := []struct {
		name            string
		args            int
		wantParentIndex int
	}{
		{
			name:            "Negative index",
			args:            -1,
			wantParentIndex: 0,
		},
		{
			name:            "Index zero",
			args:            0,
			wantParentIndex: 0,
		},
		{
			name:            "Index 1",
			args:            1,
			wantParentIndex: 0,
		},
		{
			name:            "Index 5",
			args:            5,
			wantParentIndex: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotParentIndex := heapMock.parent(tt.args); gotParentIndex != tt.wantParentIndex {
				t.Errorf("Heap.parent() = %v, want %v", gotParentIndex, tt.wantParentIndex)
			}
		})
	}
}

func TestHeap_leftChild(t *testing.T) {
	tests := []struct {
		name                string
		args                int
		wantLeftChildIndex  int
		wantLeftChildExists bool
	}{
		{
			name:                "Negative index",
			args:                -1,
			wantLeftChildIndex:  0,
			wantLeftChildExists: true,
		},
		{
			name:                "Index zero",
			args:                0,
			wantLeftChildIndex:  1,
			wantLeftChildExists: true,
		},
		{
			name:                "Index 1",
			args:                1,
			wantLeftChildIndex:  3,
			wantLeftChildExists: true,
		},
		{
			name:                "Index 5",
			args:                5,
			wantLeftChildIndex:  11,
			wantLeftChildExists: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLeftChildIndex, gotLeftChildExists := heapMock.leftChild(tt.args)
			if gotLeftChildIndex != tt.wantLeftChildIndex {
				t.Errorf("Heap.leftChild() gotLeftChildIndex = %v, want %v", gotLeftChildIndex, tt.wantLeftChildIndex)
			}
			if gotLeftChildExists != tt.wantLeftChildExists {
				t.Errorf("Heap.leftChild() gotLeftChildExists = %v, want %v", gotLeftChildExists, tt.wantLeftChildExists)
			}
		})
	}
}

func TestHeap_rightChild(t *testing.T) {
	tests := []struct {
		name                 string
		args                 int
		wantRightChildIndex  int
		wantRightChildExists bool
	}{
		{
			name:                 "Negative index",
			args:                 -1,
			wantRightChildIndex:  0,
			wantRightChildExists: true,
		},
		{
			name:                 "Index zero",
			args:                 0,
			wantRightChildIndex:  2,
			wantRightChildExists: true,
		},
		{
			name:                 "Index 1",
			args:                 1,
			wantRightChildIndex:  4,
			wantRightChildExists: true,
		},
		{
			name:                 "Index 5",
			args:                 5,
			wantRightChildIndex:  12,
			wantRightChildExists: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRightChildIndex, gotRightChildExists := heapMock.rightChild(tt.args)
			if gotRightChildIndex != tt.wantRightChildIndex {
				t.Errorf("Heap.rightChild() gotRightChildIndex = %v, want %v", gotRightChildIndex, tt.wantRightChildIndex)
			}
			if gotRightChildExists != tt.wantRightChildExists {
				t.Errorf("Heap.rightChild() gotRightChildExists = %v, want %v", gotRightChildExists, tt.wantRightChildExists)
			}
		})
	}
}

func TestHeap_Heapify(t *testing.T) {
	tests := []struct {
		name     string
		h        Heap[int]
		args     int
		wantHeap Heap[int]
	}{
		{
			name:     "Simple 123",
			h:        Heap[int]{array: []int{1, 2, 3}, length: 3},
			args:     0,
			wantHeap: Heap[int]{array: []int{3, 2, 1}, length: 3},
		},
		{
			name:     "Empty Heap",
			h:        Heap[int]{array: []int{}},
			args:     0,
			wantHeap: Heap[int]{array: []int{}},
		},
		{
			name:     "Three levels",
			h:        Heap[int]{array: []int{8, 18, 14, 17, 12, 13, 11, 15, 16}, length: 9},
			args:     0,
			wantHeap: Heap[int]{array: []int{18, 17, 14, 16, 12, 13, 11, 15, 8}, length: 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Heapify(tt.args, func(a, b int) bool { return a < b })
			assert.Equalf(t, tt.wantHeap, tt.h, "Heap.Heapify() got Heap = %v, want %v", tt.h.array, tt.wantHeap.array)
		})
	}
}
