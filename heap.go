package dsa

type Heap struct {
	arr []int
}

func NewHeap() *Heap {
	return new(Heap)
}

func (h *Heap) Heapify(i int) {
	size := len(h.arr)
	largest := i
	l := 2*i + 1
	r := 2*i + 2
	if l < size && h.arr[l] > h.arr[largest] {
		largest = l
	}
	if r < size && h.arr[r] > h.arr[largest] {
		largest = r
	}

	if largest != i {
		h.arr[i], h.arr[largest] = h.arr[largest], h.arr[i]
		h.Heapify(largest)
	}
}

func (h *Heap) Insert(num int) {
	h.arr = append(h.arr, num)
	size := len(h.arr)
	if size > 1 {
		for i := size/2 - 1; i >= 0; i-- {
			h.Heapify(i)
		}
	}
}

func (h *Heap) Delete(num int) {
	size := len(h.arr)
	var i int
	for i = 0; i < size; i++ {
		if h.arr[i] == num {
			break
		}
	}

	h.arr[i], h.arr[size-1] = h.arr[size-1], h.arr[i]
	h.arr = h.arr[:size-1]
	size--

	for i = size/2 - 1; i >= 0; i-- {
		h.Heapify(i)
	}
}
