package dsa

import (
	"reflect"
	"testing"
)

func TestHeap(t *testing.T) {
	heap := NewHeap()
	heap.Insert(3)
	heap.Insert(4)
	heap.Insert(9)
	heap.Insert(5)
	heap.Insert(2)

	want := []int{9, 5, 4, 3, 2}
	if !reflect.DeepEqual(heap.arr, want) {
		t.Errorf("heap = %v, want %v", heap.arr, want)
	}

	heap.Delete(4)
	want = []int{9, 5, 2, 3}
	if !reflect.DeepEqual(heap.arr, want) {
		t.Errorf("heap after deleting 4 = %v, want %v", heap.arr, want)
	}
}
